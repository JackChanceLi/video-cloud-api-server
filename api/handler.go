package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-api-server/api/dbop"
	"go-api-server/api/defs"
	"go-api-server/api/session"
	"io/ioutil"
	"log"
	"net/http"
)
//通过用户名、密码登录验证
func LoginByMail(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Http body read failed")
	}
	ubody := &defs.UserIdentity{}
	w.Header().Set("Access-Control-Allow-Origin","*")  //"*"表示接受任意域名的请求，这个值也可以根据自己需要，设置成不同域名
	//解析包
	if err := json.Unmarshal(res, ubody); err != nil {
		fmt.Println(ubody)
		sendErrorResponse(w, defs.ErrRequestBodyParseFailed)
		return
	}
	fmt.Println(ubody)
	userInfo, password, err := dbop.UserLogin(ubody.Email)

	fmt.Println(password)
	fmt.Println(userInfo)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	cid := userInfo.Cid
	name := userInfo.Name
	auth := userInfo.Auth
	email := userInfo.Email
	aid := userInfo.Aid
	if password == ubody.Password {
		//fmt.Println("Login successfully!")

		id := session.GenerateNewSessionID(userInfo.Cid)
		su := &defs.SignedUp{Code:201,Data:defs.DataForUser{SessionID:id,User:defs.UserInformation{Aid:aid, Cid:cid,Name:name, Email:email, Auth:auth}}, Msg:defs.Message{Error:"",ErrorCode:""}}

		if resp, err := json.Marshal(su); err != nil {
			sendErrorResponse(w, defs.ErrorInternalFaults)
			return
		} else {
			fmt.Println(su)
			sendNormalResponse(w, string(resp),200)
		}
	} else {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}
}

//用户注册的实现，包含用户名、密码、邮箱、权限等信息
func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res, err := ioutil.ReadAll(r.Body)
	fmt.Println(res)
	if err != nil {
		log.Printf("Http body read failed")
	}
	w.Header().Set("Access-Control-Allow-Origin","*")  //"*"表示接受任意域名的请求，这个值也可以根据自己需要，设置成不同域名
	ubody := &defs.UserIdentity{}
	if err := json.Unmarshal(res, ubody); err != nil {
		fmt.Println(ubody)
		sendErrorResponse(w, defs.ErrRequestBodyParseFailed)
		return
	}
	fmt.Println(ubody)

	ok, err := dbop.IsEmailRegister(ubody.Email)  //判断该邮箱是否已经注册
	if !ok && err == nil {  //表示该邮箱已经被注册
		sendErrorResponse(w, defs.ErrorEmailRegistered)
		return
	}
	if err := dbop.UserRegister(ubody.UserName, ubody.Email, ubody.Password); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	su := defs.SignedUp{Code:200}

	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp),200)
	}
	//fmt.Fprintf(w, "Request success!\n")
}
//路由测试函数，目前没有太大用处
func Handle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "111111111111111111111111111111\n")
}


func CrteateLiveRoom(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	liveInfo := &defs.LiveRoom{}
	liveInfo.Code = 200
	liveInfo.Data.LiveRoomInfo.Lid = "000001"
	liveInfo.Data.LiveRoomInfo.Name = "ljc"
	liveInfo.Data.LiveRoomInfo.Kind = 1
	liveInfo.Data.LiveRoomInfo.Size = 100
	liveInfo.Data.LiveRoomInfo.StartTime = "2006-01-02 15:04:05"
	liveInfo.Data.LiveRoomInfo.EndTime = "2006-01-02 15:04:05"
	liveInfo.Data.LiveRoomInfo.PushUrl = "www.baidu.com"
	liveInfo.Data.LiveRoomInfo.PullHlsUrl = "www.google.com"
	liveInfo.Data.LiveRoomInfo.PullRtmpUrl = "rtmp://www.ljc.com"
	liveInfo.Data.LiveRoomInfo.PullHttpFlvUrl = "www.hlv.com"
	liveInfo.Data.LiveRoomInfo.DisplayUrl = "www.display.com"
	liveInfo.Data.LiveRoomInfo.Status = 5
	liveInfo.Data.LiveRoomInfo.Permission = "Auth users"
	liveInfo.Msg.ErrorCode = ""
	liveInfo.Msg.Error = ""

	if resp, err := json.Marshal(liveInfo); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp),200)
	}
}


func test(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {


}