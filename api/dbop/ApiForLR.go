package dbop

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-api-server/api/defs"
	"go-api-server/api/utils"
	"log"
	"sync"
)

func CreateLiveRoomByCom(cid string, name string, kind int, size int, startTime string, endTime string) (*defs.LiveRoomIdentity, error ){ //超级管理员创建用户

	createTime, _ := getCurrentTime()
	permission := "logged_user"
	status := 2
	stmtIns, err := dbConn.Prepare("INSERT INTO live_room (lid, cid, name, kind, size, start_time, end_time, push_url, pull_hls_url, pull_rtmp_url, pull_http_flv_url, display_url, status, permission, create_time) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	lid, _ := utils.NewUUID()
	pushUrl := "www.baidu.com"
	pullHlsUrl := "www.11111.com"
	pullRtmpUrl := "www.22222.com"
	pullHttpFlvUrl := "www.33333.com"
	displayUrl := "www.44444.com"

	_,err = stmtIns.Exec(lid, cid, name, kind, size, startTime, endTime, pushUrl, pullHlsUrl, pullRtmpUrl, pullHttpFlvUrl, displayUrl, status, permission, createTime)
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

	log.Printf(" Delete success")

	defer stmtOut.Close()
	return nil
}

func UpdateLiveRoom(aid string, lid string, name string, kind int, size int, start_time string, end_time string, permission string) (*defs.LiveRoomIdentity, error) {
	stmtUpa, err := dbConn.Prepare("UPDATE live_room SET name = ?, kind = ?, size = ?, start_time = ?, end_time = ?, permission = ? WHERE lid = ?")
	if err != nil {
		log.Printf("%s",err)
		return nil, err
	}

	_,err = stmtUpa.Exec(name, kind, size, start_time, end_time, permission, lid)
	if err != nil {
		return nil, err
	}

	log.Printf(" Update success")

	defer stmtUpa.Close()
	Lr, err := RetrieveLiveRoomByLid(lid)
		if err != nil {
		return nil, err
	}

	room := &defs.LiveRoomIdentity{}
	room.Aid = aid
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
		var aid,lid, cid, name, start_time, end_time, push_url, pull_hls_url, pull_rtmp_url, pull_http_flv_url, display_url, permission, create_time string
		var kind, size, status int
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
	stmtOut, err := dbConn.Prepare("SELECT cid, name, kind, size, start_time, end_time, push_url, pull_hls_url, pull_rtmp_url, pull_http_flv_url, display_url, status, permission, create_time FROM live_room WHERE lid = ?")
	if err != nil && err != sql.ErrNoRows{
		log.Printf("%s", err)
		return nil, err
	}
	var  cid, name, start_time, end_time, push_url, pull_hls_url, pull_rtmp_url, pull_http_flv_url, display_url, permission, create_time string
	var kind, size, status int
	stmtOut.QueryRow(Lid).Scan(&cid, &name, &kind, &size, &start_time, &end_time, &push_url, &pull_hls_url, &pull_rtmp_url, &pull_http_flv_url, &display_url, &status, &permission, &create_time)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	Lri := &defs.LiveRoomIdentity{Lid: Lid, Cid: cid,Name: name, Kind: kind, Size: size, StartTime: start_time, EndTime: end_time,
		PushUrl: push_url, PullHlsUrl: pull_hls_url, PullRtmpUrl: pull_rtmp_url, PullHttpFlvUrl: pull_http_flv_url,
		DisplayUrl: display_url, Status: status, Permission: permission, CreateTime: create_time}//绑定全部信息
	defer stmtOut.Close()
	return Lri, nil
}
func CreateLiveRoomByAdmin (cid, aid, name, startTime, endTime string, kind, size int) (*defs.LiveRoomIdentity, error ) {
	createTime, _ := getCurrentTime()
	permission := "logged_user"
	status := 2
	stmtIns, err := dbConn.Prepare("INSERT INTO live_room (lid, cid, name, kind, size, start_time, end_time, push_url, pull_hls_url, pull_rtmp_url, pull_http_flv_url, display_url, status, permission, create_time) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	lid, _ := utils.NewUUID()
	pushUrl := "www.baidu.com"
	pullHlsUrl := "www.11111.com"
	pullRtmpUrl := "www.22222.com"
	pullHttpFlvUrl := "www.33333.com"
	displayUrl := "www.44444.com"

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
		var lid, cid, name, startTime, endTime, pushUrl, pullHlsUrl, pullRtmpUrl, pullHttpFlvUrl, displayUrl, permission, createTime string
		var kind, size, status int
		if er := rows.Scan(&lid, &cid, &name, &kind, &size, &startTime, &endTime, &pushUrl, &pullHlsUrl, &pullRtmpUrl, &pullHttpFlvUrl, &displayUrl, &status, &permission, &createTime); er != nil {
			log.Printf("Retrieve live_room error: %s", er)
			return nil, err
		}

		Lri := defs.LiveRoomIdentity{Lid: lid, Cid: cid,Name: name, Kind: kind, Size: size, StartTime: startTime, EndTime: endTime,
			PushUrl: pushUrl, PullHlsUrl: pullHlsUrl, PullRtmpUrl: pullRtmpUrl, PullHttpFlvUrl: pullHttpFlvUrl,
			DisplayUrl: displayUrl, Status: status, Permission: permission, CreateTime: createTime}
		Lri.Aid = cid
		room = append(room, Lri)
	}
	return room, nil
}


