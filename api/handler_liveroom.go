package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-api-server/api/dbop"
	"go-api-server/api/defs"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateLiveRoom(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cid := ps.ByName("cid")
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Http body read failed")
	}
	ubody := &defs.LiveRoomIdentity{}


	//解析包
	if err := json.Unmarshal(res, ubody); err != nil {
		fmt.Println(ubody)
		sendErrorResponse(w, defs.ErrRequestBodyParseFailed)
		return
	}
	fmt.Println(ubody)
	w.Header().Set("Access-Control-Allow-Origin","*")  //"*"表示接受任意域名的请求，这个值也可以根据自己需要，设置成不同域名
	if cid == ubody.Aid {  //表示创建者为超级管理员
		room, err := dbop.CreateLiveRoomByCom(cid, ubody.Name, ubody.Kind, ubody.Size, ubody.StartTime, ubody.EndTime)
		if  err != nil {
			sendErrorResponse(w, defs.ErrorDBError)
			return
		}
		roomInfo := &defs.LiveRoom{}
		roomInfo.Code = 200
		roomInfo.Data.LiveRoomInfo.Aid = room.Aid
		roomInfo.Data.LiveRoomInfo.Lid = room.Lid
		roomInfo.Data.LiveRoomInfo.Cid = room.Cid
		roomInfo.Data.LiveRoomInfo.Name = room.Name
		roomInfo.Data.LiveRoomInfo.Kind = room.Kind
		roomInfo.Data.LiveRoomInfo.Size = room.Size
		roomInfo.Data.LiveRoomInfo.StartTime = room.StartTime
		roomInfo.Data.LiveRoomInfo.EndTime = room.EndTime
		roomInfo.Data.LiveRoomInfo.PushUrl = room.PushUrl
		roomInfo.Data.LiveRoomInfo.PullHlsUrl = room.PullHlsUrl
		roomInfo.Data.LiveRoomInfo.PullRtmpUrl = room.PullRtmpUrl
		roomInfo.Data.LiveRoomInfo.PullHttpFlvUrl = room.PullHttpFlvUrl
		roomInfo.Data.LiveRoomInfo.DisplayUrl = room.DisplayUrl
		roomInfo.Data.LiveRoomInfo.Status = room.Status
		roomInfo.Data.LiveRoomInfo.Permission = room.Permission
		roomInfo.Data.LiveRoomInfo.CreateTime = room.CreateTime

		if resp, err := json.Marshal(roomInfo); err != nil {
			sendErrorResponse(w, defs.ErrorInternalFaults)
			return
		} else {
			sendNormalResponse(w, string(resp),200)
		}

	} else {  //表示创建者为普通管理员
		room, err := dbop.CreateLiveRoomByAdmin(cid, ubody.Aid, ubody.Name, ubody.StartTime, ubody.EndTime, ubody.Kind, ubody.Size)
		if  err != nil {
			log.Println(err)
			sendErrorResponse(w, defs.ErrorDBError)
			return
		}

		roomInfo := &defs.LiveRoom{}
		roomInfo.Code = 200
		roomInfo.Data.LiveRoomInfo.Aid = room.Aid
		roomInfo.Data.LiveRoomInfo.Lid = room.Lid
		roomInfo.Data.LiveRoomInfo.Cid = room.Cid
		roomInfo.Data.LiveRoomInfo.Name = room.Name
		roomInfo.Data.LiveRoomInfo.Kind = room.Kind
		roomInfo.Data.LiveRoomInfo.Size = room.Size
		roomInfo.Data.LiveRoomInfo.StartTime = room.StartTime
		roomInfo.Data.LiveRoomInfo.EndTime = room.EndTime
		roomInfo.Data.LiveRoomInfo.PushUrl = room.PushUrl
		roomInfo.Data.LiveRoomInfo.PullHlsUrl = room.PullHlsUrl
		roomInfo.Data.LiveRoomInfo.PullRtmpUrl = room.PullRtmpUrl
		roomInfo.Data.LiveRoomInfo.PullHttpFlvUrl = room.PullHttpFlvUrl
		roomInfo.Data.LiveRoomInfo.DisplayUrl = room.DisplayUrl
		roomInfo.Data.LiveRoomInfo.Status = room.Status
		roomInfo.Data.LiveRoomInfo.Permission = room.Permission
		roomInfo.Data.LiveRoomInfo.CreateTime = room.CreateTime

		if resp, err := json.Marshal(roomInfo); err != nil {
			sendErrorResponse(w, defs.ErrorInternalFaults)
			return
		} else {
			sendNormalResponse(w, string(resp),200)
		}
	}
}

func UpdateLiveRoom(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {//更新直播间信息
	cid := ps.ByName("cid")//获取cid
	au := r.URL.Query()
	aid := au.Get("aid")//获取aid
	log.Printf("Aid value is [%s]\n", aid)

	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Http body read failed")
	}
	ubody := &defs.LiveRoomIdentity{}

	w.Header().Set("Access-Control-Allow-Origin","*")  //"*"表示接受任意域名的请求，这个值也可以根据自己需要，设置成不同域名
	//解析包
	if err := json.Unmarshal(res, ubody); err != nil {
		fmt.Println(ubody)
		sendErrorResponse(w, defs.ErrRequestBodyParseFailed)//解析错误
		return
	}
	fmt.Println(ubody)

	if cid == ubody.Aid { //表示更新者为超级管理员
		room, err := dbop.UpdateLiveRoom(aid, ubody.Lid, ubody.Name, ubody.Kind, ubody.Size, ubody.StartTime, ubody.EndTime, ubody.Permission)
		if err != nil {
			sendErrorResponse(w, defs.ErrorDBError)
			return
		}
		roomInfo := &defs.LiveRoom{}
		roomInfo.Code = 200
		roomInfo.Data.LiveRoomInfo.Aid = room.Aid
		roomInfo.Data.LiveRoomInfo.Lid = room.Lid
		roomInfo.Data.LiveRoomInfo.Cid = room.Cid
		roomInfo.Data.LiveRoomInfo.Name = room.Name
		roomInfo.Data.LiveRoomInfo.Kind = room.Kind
		roomInfo.Data.LiveRoomInfo.Size = room.Size
		roomInfo.Data.LiveRoomInfo.StartTime = room.StartTime
		roomInfo.Data.LiveRoomInfo.EndTime = room.EndTime
		roomInfo.Data.LiveRoomInfo.PushUrl = room.PushUrl
		roomInfo.Data.LiveRoomInfo.PullHlsUrl = room.PullHlsUrl
		roomInfo.Data.LiveRoomInfo.PullRtmpUrl = room.PullRtmpUrl
		roomInfo.Data.LiveRoomInfo.PullHttpFlvUrl = room.PullHttpFlvUrl
		roomInfo.Data.LiveRoomInfo.DisplayUrl = room.DisplayUrl
		roomInfo.Data.LiveRoomInfo.Status = room.Status
		roomInfo.Data.LiveRoomInfo.Permission = room.Permission
		roomInfo.Data.LiveRoomInfo.CreateTime = room.CreateTime

		if resp, err := json.Marshal(roomInfo); err != nil {
			sendErrorResponse(w, defs.ErrorInternalFaults)
			return
		} else {
			sendNormalResponse(w, string(resp), 200)
		}
	} else {  //表示更新者为普通管理员
		room, err := dbop.UpdateLiveRoom(aid, ubody.Lid, ubody.Name, ubody.Kind, ubody.Size, ubody.StartTime, ubody.EndTime, ubody.Permission)
		if err != nil {
			sendErrorResponse(w, defs.ErrorDBError)
			return
		}
		roomInfo := &defs.LiveRoom{}
		roomInfo.Code = 200
		roomInfo.Data.LiveRoomInfo.Aid = room.Aid
		roomInfo.Data.LiveRoomInfo.Lid = room.Lid
		roomInfo.Data.LiveRoomInfo.Cid = room.Cid
		roomInfo.Data.LiveRoomInfo.Name = room.Name
		roomInfo.Data.LiveRoomInfo.Kind = room.Kind
		roomInfo.Data.LiveRoomInfo.Size = room.Size
		roomInfo.Data.LiveRoomInfo.StartTime = room.StartTime
		roomInfo.Data.LiveRoomInfo.EndTime = room.EndTime
		roomInfo.Data.LiveRoomInfo.PushUrl = room.PushUrl
		roomInfo.Data.LiveRoomInfo.PullHlsUrl = room.PullHlsUrl
		roomInfo.Data.LiveRoomInfo.PullRtmpUrl = room.PullRtmpUrl
		roomInfo.Data.LiveRoomInfo.PullHttpFlvUrl = room.PullHttpFlvUrl
		roomInfo.Data.LiveRoomInfo.DisplayUrl = room.DisplayUrl
		roomInfo.Data.LiveRoomInfo.Status = room.Status
		roomInfo.Data.LiveRoomInfo.Permission = room.Permission
		roomInfo.Data.LiveRoomInfo.CreateTime = room.CreateTime

		if resp, err := json.Marshal(roomInfo); err != nil {
			sendErrorResponse(w, defs.ErrorInternalFaults)
			return
		} else {
			sendNormalResponse(w, string(resp), 200)
		}
	}
}

func GetLiveRooms(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cid := ps.ByName("cid")
	vars := r.URL.Query()
	aid := vars.Get("aid")

	w.Header().Set("Access-Control-Allow-Origin","*")  //"*"表示接受任意域名的请求，这个值也可以根据自己需要，设置成不同域名

	if aid == cid {
		room, err := dbop.RetrieveLiveRoomByCid(cid)
		if err != nil {
			sendErrorResponse(w, defs.ErrorDBError)
		}
		fmt.Println(room)

		liveroomList := &defs.LiveRoomList{}
		liveroomList.Code = 200
		liveroomList.Data = room
		if resp, err := json.Marshal(liveroomList); err != nil {
			sendErrorResponse(w, defs.ErrorInternalFaults)
			return
		} else {
			sendNormalResponse(w, string(resp),200)
		}
	}
}

func DeleteLiveRoom(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//cid := ps.ByName("cid")//获取cid
	au := r.URL.Query()
	aid := au.Get("aid")//获取aid
	lid := au.Get("lid")//获取lid
	log.Printf("Aid value is [%s]\n", aid)
	log.Printf("Lid value is [%s]\n", lid)

	w.Header().Set("Access-Control-Allow-Origin","*")  //"*"表示接受任意域名的请求，这个值也可以根据自己需要，设置成不同域名
	err := dbop.DeleteLiveRoom(lid)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	su := defs.LiveRoom{Code:200}

	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	} else {
		sendNormalResponse(w, string(resp),200)
	}
}

func GetLiveRoomByLid(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cid := ps.ByName("cid")
	vars := r.URL.Query()
	lid := vars.Get("lid")
	aid := vars.Get("aid")

	w.Header().Set("Access-Control-Allow-Origin","*")  //"*"表示接受任意域名的请求，这个值也可以根据自己需要，设置成不同域名

	room, err := dbop.RetrieveLiveRoomByLid(lid)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	if room.Cid == "" {  //等价于err == sql.ErrNoRows
		sendErrorResponse(w, defs.ErrNoRowsInDB)
		return
	}
	if room.Cid != cid {
		sendErrorResponse(w, defs.ErrorNotAuthUserForRoom)
		return
	}
	roomInfo := &defs.LiveRoom{}
	roomInfo.Code = 200
	roomInfo.Data.LiveRoomInfo.Aid = aid
	roomInfo.Data.LiveRoomInfo.Lid = room.Lid
	roomInfo.Data.LiveRoomInfo.Cid = room.Cid
	roomInfo.Data.LiveRoomInfo.Name = room.Name
	roomInfo.Data.LiveRoomInfo.Kind = room.Kind
	roomInfo.Data.LiveRoomInfo.Size = room.Size
	roomInfo.Data.LiveRoomInfo.StartTime = room.StartTime
	roomInfo.Data.LiveRoomInfo.EndTime = room.EndTime
	roomInfo.Data.LiveRoomInfo.PushUrl = room.PushUrl
	roomInfo.Data.LiveRoomInfo.PullHlsUrl = room.PullHlsUrl
	roomInfo.Data.LiveRoomInfo.PullRtmpUrl = room.PullRtmpUrl
	roomInfo.Data.LiveRoomInfo.PullHttpFlvUrl = room.PullHttpFlvUrl
	roomInfo.Data.LiveRoomInfo.DisplayUrl = room.DisplayUrl
	roomInfo.Data.LiveRoomInfo.Status = room.Status
	roomInfo.Data.LiveRoomInfo.Permission = room.Permission
	roomInfo.Data.LiveRoomInfo.CreateTime = room.CreateTime

	if resp, err := json.Marshal(roomInfo); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp), 200)
	}

}
