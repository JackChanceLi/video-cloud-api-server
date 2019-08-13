package dbop

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-api-server/api/defs"
	"go-api-server/api/utils"
	"go-api-server/api/url"
	"log"
	"sync"
)

func CreateLiveRoomByCom(cid string, name string, kind int, size int, startTime string, endTime string) (*defs.LiveRoomIdentity, error ){ //超级管理员创建用户

	createTime, _ := getCurrentTime()
	permission := 1
	status := 2
	pictureUrl := "http://pic-cloud-bupt.oss-cn-beijing.aliyuncs.com/%E7%BE%8E%E5%9B%BD%E5%BE%80%E4%BA%8B.jpg"

	stmtIns, err := dbConn.Prepare("INSERT INTO live_room (lid, cid, name, kind, size, start_time, end_time, push_url, pull_hls_url, pull_rtmp_url, pull_http_flv_url, display_url, status, permission, create_time, picture_url) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	lid, _ := utils.NewUUID()
	streamId, _ := utils.NewStreamID()
	pushUrl,_ := url.NewRtmpUrl(streamId)
	pullHlsUrl,_ := url.NewHlsUrl(streamId)
	pullRtmpUrl := pushUrl
	pullHttpFlvUrl,_ := url.NewFlvUrl(streamId)
	displayUrl,_ := url.NewDisplayUrl(lid)

	_,err = stmtIns.Exec(lid, cid, name, kind, size, startTime, endTime, pushUrl, pullHlsUrl, pullRtmpUrl, pullHttpFlvUrl, displayUrl, status, permission, createTime, pictureUrl)
	if err != nil {
		return nil, err
	}

	log.Printf(" Create live_room success")

	defer stmtIns.Close()
	room := &defs.LiveRoomIdentity{}
	room.Lid =lid
	room.Cid = cid
	room.Name = name
	room.Kind = kind
	room.Size =size
	room.StartTime = startTime
	room.EndTime = endTime
	room.PushUrl = pushUrl
	room.PullHlsUrl = pullHlsUrl
	room.PullRtmpUrl = pullRtmpUrl
	room.PullHttpFlvUrl = pullHttpFlvUrl
	room.DisplayUrl = displayUrl
	room.Status = status
	room.Permission = permission
	room.CreateTime = createTime
	room.PictureUrl = pictureUrl

	defaultConfig := defs.LiveRoomDefaultConfig
    //为直播间设定默认引导界面信息
	_, err = InsertLRIntroByCom(lid, defaultConfig.Qorder, defaultConfig.Prepic)
	if err != nil {
		log.Printf("Error of liverom_intro default setting:%v", err)
		return nil, err
	}
	//为直播间设定默认页面设置
	_, err = InsertLRConfigByCom(lid, defaultConfig.LivePic,defaultConfig.Danmu, defaultConfig.Chat, defaultConfig.Share,
		                defaultConfig.ShareText, defaultConfig.Advertisement, defaultConfig.AdJumpUrl,
		                defaultConfig.AdPicUrl, defaultConfig.AdText)
	if err != nil {
		log.Printf("Error of liverom_config default setting:%v", err)
		return nil, err
	}
	//为直播间设定观看条件设置
	_, err = InsertLRConditionByCom(lid, defaultConfig.VerificationCode, defaultConfig.Condition, defaultConfig.ConditionType, defaultConfig.Duration,
		defaultConfig.TryToSee, defaultConfig.Price)

	if err != nil {
		log.Printf("Error of liverom_condition default setting:%v", err)
		return nil, err
	}
	//为直播间设定服务设置
	_, err = InsertLRQualityByCom(lid, defaultConfig.Delay, defaultConfig.Transcode, defaultConfig.TranscodeType, defaultConfig.Record,
	                     defaultConfig.RecordType)
	if err != nil {
		log.Printf("Error of liverom_quality default setting:%v", err)
		return nil, err
	}
	//为直播间设定版本安全设置
	_, err = InsertLRSafeByCom(lid,defaultConfig.Logo, defaultConfig.LogoUrl, defaultConfig.LogoPosition, defaultConfig.LogoTransparency,
		              defaultConfig.Lamp, defaultConfig.LampType, defaultConfig.LampText, defaultConfig.LampFontSize,
		              defaultConfig.LogoTransparency)
	if err != nil {
		log.Printf("Error of liverom_safe default setting:%v", err)
		return nil, err
	}
	//为直播间设定权限安全设置，默认黑白名单均为空，无需处理
	_, err = InsertLRAuthSafeList(lid, defaultConfig.WhiteSiteList, defaultConfig.BlackSiteList)
	if err != nil {
		log.Printf("Error of liverom_auth_safe default setting:%v", err)
		return nil, err
	}
	return  room, nil
}

func DeleteLiveRoom(lid string) error {
	stmtOut, err := dbConn.Prepare("DELETE FROM live_room WHERE lid = ?")
	if err != nil {
		log.Printf("%s",err)
		return err
	}

	if _, err := stmtOut.Query(lid); err != nil {
		return err
	}

	defer stmtOut.Close()
	return nil
}

func UpdateLiveRoom(lid string, name string, kind int, size int, startTime string, endTime, pictureUrl string, permission int) (*defs.LiveRoomIdentity, error) {
	stmtUpa, err := dbConn.Prepare("UPDATE live_room SET name = ?, kind = ?, size = ?, start_time = ?, end_time = ?, permission = ?, picture_url = ? WHERE lid = ?")
	if err != nil {
		log.Printf("%s",err)
		return nil, err
	}

	_,err = stmtUpa.Exec(name, kind, size, startTime, endTime, permission, pictureUrl, lid)
	if err != nil {
		return nil, err
	}

	defer stmtUpa.Close()
	Lr, err := RetrieveLiveRoomByLid(lid)
		if err != nil {
		return nil, err
	}

	room := &defs.LiveRoomIdentity{}
	room.Aid = Lr.Aid
	room.Lid = lid
	room.Cid = Lr.Cid
	room.Name = name
	room.Kind = kind
	room.Size =size
	room.StartTime = Lr.StartTime
	room.EndTime = Lr.EndTime
	room.PushUrl = Lr.PushUrl
	room.PullHlsUrl = Lr.PullHlsUrl
	room.PullRtmpUrl = Lr.PullRtmpUrl
	room.PullHttpFlvUrl = Lr.PullHttpFlvUrl
	room.DisplayUrl = Lr.DisplayUrl
	room.Status = Lr.Status
	room.Permission = permission
	room.CreateTime = Lr.CreateTime
	room.PictureUrl = pictureUrl
	return room, nil
}

func SearchLiveRoomByCid(Cid string) (*sync.Map, int, error) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("SELECT * FROM live_room WHERE cid = ?")
	cnt := 0
	if err != nil {
		log.Printf("%s", err)
		return nil, 0, err
	}
	//cid := Cid
	rows, err := stmtOut.Query(Cid)
	if err != nil {
		log.Printf("%s", err)
		return nil, 0, err
	}

	for rows.Next() {
		var aid,lid, cid, name, start_time, end_time, push_url, pull_hls_url, pull_rtmp_url, pull_http_flv_url, display_url, create_time string
		var kind, size, status, permission int
		if er := rows.Scan(&aid, &lid, &cid, &name, &kind, &size, &start_time, &end_time, &push_url, &pull_hls_url, &pull_rtmp_url, &pull_http_flv_url, &display_url, &status, &permission, &create_time); er != nil {
			log.Printf("Retrieve live_room error: %s", er)
			break
		}

		Lri := &defs.LiveRoomIdentity{Aid: aid, Lid: lid, Cid: cid,Name: name, Kind: kind, Size: size, StartTime: start_time, EndTime: end_time,
			PushUrl: push_url, PullHlsUrl: pull_hls_url, PullRtmpUrl: pull_rtmp_url, PullHttpFlvUrl: pull_http_flv_url,
			DisplayUrl: display_url, Status: status, Permission: permission, CreateTime: create_time}
		m.Store(cnt,Lri)
		cnt++
	}
	return m, cnt, nil
}

func RetrieveLiveRoomByLid(Lid string) (*defs.LiveRoomIdentity, error) {//通过lid遍历查询
	stmtOut, err := dbConn.Prepare("SELECT cid, name, kind, size, start_time, end_time, push_url, pull_hls_url, pull_rtmp_url, pull_http_flv_url, display_url, status, permission, create_time, picture_url FROM live_room WHERE lid = ?")
	if err != nil && err != sql.ErrNoRows{
		log.Printf("%s", err)
		return nil, err
	}
	var  cid, name, startTime, endTime, pushUrl, pullHlsUrl, pullRtmpUrl, pullHttpFlvUrl, displayUrl, createTime, pictureUrl string
	var kind, size, status, permission int
	stmtOut.QueryRow(Lid).Scan(&cid, &name, &kind, &size, &startTime, &endTime, &pushUrl, &pullHlsUrl, &pullRtmpUrl, &pullHttpFlvUrl, &displayUrl, &status, &permission, &createTime, &pictureUrl)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	Lri := &defs.LiveRoomIdentity{Lid: Lid, Cid: cid,Name: name, Kind: kind, Size: size, StartTime: startTime, EndTime: endTime,
		PushUrl: pushUrl, PullHlsUrl: pullHlsUrl, PullRtmpUrl: pullRtmpUrl, PullHttpFlvUrl: pullHttpFlvUrl,
		DisplayUrl: displayUrl, Status: status, Permission: permission, CreateTime: createTime, PictureUrl: pictureUrl} //绑定全部信息
	defer stmtOut.Close()
	return Lri, nil
}

func CreateLiveRoomByAdmin (cid, aid, name, startTime, endTime string, kind, size int) (*defs.LiveRoomIdentity, error ) {
	createTime, _ := getCurrentTime()
	permission := 1
	status := 2
	stmtIns, err := dbConn.Prepare("INSERT INTO live_room (lid, cid, name, kind, size, start_time, end_time, push_url, pull_hls_url, pull_rtmp_url, pull_http_flv_url, display_url, status, permission, create_time) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	lid, _ := utils.NewUUID()
	streamId, _ := utils.NewStreamID()
	pushUrl,_ := url.NewRtmpUrl(streamId)
	pullHlsUrl,_ := url.NewHlsUrl(streamId)
	pullRtmpUrl := pushUrl
	pullHttpFlvUrl,_ := url.NewFlvUrl(streamId)
	displayUrl,_ := url.NewDisplayUrl(lid)

	_,err = stmtIns.Exec(lid, cid, name, kind, size, startTime, endTime, pushUrl, pullHlsUrl, pullRtmpUrl, pullHttpFlvUrl, displayUrl, status, permission, createTime)
	if err != nil {
		return nil, err
	}

	stmtIns1, err := dbConn.Prepare("INSERT INTO UToL(aid, lid) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}
	_, err = stmtIns1.Exec(aid, lid)
	if err != nil {
		return nil, err
	}

	defer stmtIns.Close()
	log.Printf(" Create live_room success")
	room := &defs.LiveRoomIdentity{}
	room.Lid =lid
	room.Aid = aid
	room.Cid = cid
	room.Name = name
	room.Kind = kind
	room.Size =size
	room.StartTime = startTime
	room.EndTime = endTime
	room.PushUrl = pushUrl
	room.PullHlsUrl = pullHlsUrl
	room.PullRtmpUrl = pullRtmpUrl
	room.PullHttpFlvUrl = pullHttpFlvUrl
	room.DisplayUrl = displayUrl
	room.Status = status
	room.Permission = permission
	room.CreateTime = createTime

	return  room, nil
}


func RetrieveLiveRoomByCid(Cid string) ([] defs.LiveRoomIdentity, error) {  //以切片的形式返回查询直播间的结果
	var room [] defs.LiveRoomIdentity
	stmtOut, err := dbConn.Prepare("SELECT * FROM live_room WHERE cid = ?")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	//cid := Cid
	rows, err := stmtOut.Query(Cid)
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	for rows.Next() {
		var lid, cid, name, startTime, endTime, pushUrl, pullHlsUrl, pullRtmpUrl, pullHttpFlvUrl, displayUrl, createTime, pictureUrl string
		var kind, size, status, permission int
		if er := rows.Scan(&lid, &cid, &name, &kind, &size, &startTime, &endTime, &pushUrl, &pullHlsUrl, &pullRtmpUrl, &pullHttpFlvUrl, &displayUrl, &status, &permission, &createTime, &pictureUrl); er != nil {
			log.Printf("Retrieve live_room error: %s", er)
			return nil, err
		}

		Lri := defs.LiveRoomIdentity{Lid: lid, Cid: cid,Name: name, Kind: kind, Size: size, StartTime: startTime, EndTime: endTime,
			PushUrl: pushUrl, PullHlsUrl: pullHlsUrl, PullRtmpUrl: pullRtmpUrl, PullHttpFlvUrl: pullHttpFlvUrl,
			DisplayUrl: displayUrl, Status: status, Permission: permission, CreateTime: createTime, PictureUrl:pictureUrl}
		Lri.Aid = cid
		room = append(room, Lri)
	}
	return room, nil
}

func SearchAuth(aid string, lid string) (bool, error) { //在UToL表中寻找是否有更新权限
	stmtOut, err := dbConn.Prepare("SELECT * FROM UToL WHERE aid = ? and lid = ?")
	if err != nil {
		log.Printf("%s", err)
		return false, err
	}

	var naid, nlid string
	stmtOut.QueryRow(aid, lid).Scan(&naid, &nlid)
	if err != nil {
		log.Printf("%s", err)
		return false, err
	}

	if naid == "" && nlid == "" {
		log.Printf("No rows")
		return false,nil
	}
	 return true, nil
}
