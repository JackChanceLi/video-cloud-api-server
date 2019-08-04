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

func InsertLRConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { //插入直播界面信息
	//cid := ps.ByName("cid")
	au := r.URL.Query()
	aid := au.Get("aid")//获取aid
	log.Printf("Aid value is [%s]\n", aid)

	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Http body read failed")
	}
	ubody := &defs.LiveRoomConfigIdentity{}

	//解析包
	if err := json.Unmarshal(res, ubody); err != nil {
		fmt.Println(ubody)
		sendErrorResponse(w, defs.ErrRequestBodyParseFailed)
		return
	}
	fmt.Println(ubody)
	Res, err := dbop.InsertLRConfigByCom(ubody.Lid, ubody.LivePic, ubody.Danmu, ubody.Chat, ubody.Share, ubody.ShareText, ubody.Advertisement, ubody.AdJumpUrl, ubody.AdPicUrl, ubody.AdText)
	if  err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	LRCon := &defs.LiveRoomConfig{}
	LRCon.Code = 200
	LRCon.Data.LiveRoomConfigInfo.Lid = Res.Lid
	LRCon.Data.LiveRoomConfigInfo.LivePic = Res.LivePic
	LRCon.Data.LiveRoomConfigInfo.Danmu = Res.Danmu
	LRCon.Data.LiveRoomConfigInfo.Chat = Res.Chat
	LRCon.Data.LiveRoomConfigInfo.Share = Res.Share
	LRCon.Data.LiveRoomConfigInfo.ShareText = Res.ShareText
	LRCon.Data.LiveRoomConfigInfo.Advertisement = Res.Advertisement
	LRCon.Data.LiveRoomConfigInfo.AdJumpUrl = Res.AdJumpUrl
	LRCon.Data.LiveRoomConfigInfo.AdPicUrl = Res.AdPicUrl
	LRCon.Data.LiveRoomConfigInfo.AdText = Res.AdText

	if resp, err := json.Marshal(LRCon); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp),200)
	}
}

func UpdateLRConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {//更新直播界面信息
	//cid := ps.ByName("cid")//获取cid
	au := r.URL.Query()
	aid := au.Get("aid")//获取aid
	log.Printf("Aid value is [%s]\n", aid)

	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Http body read failed")
	}
	ubody := &defs.LiveRoomConfigIdentity{}

	//解析包
	if err := json.Unmarshal(res, ubody); err != nil {
		fmt.Println(ubody)
		sendErrorResponse(w, defs.ErrRequestBodyParseFailed)//解析错误
		return
	}
	fmt.Println(ubody)

	Res, err := dbop.UpdateLRConfig(ubody.Lid, ubody.LivePic, ubody.Danmu, ubody.Chat, ubody.Share, ubody.ShareText, ubody.Advertisement, ubody.AdJumpUrl, ubody.AdPicUrl, ubody.AdText)
	if  err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	LRCon := &defs.LiveRoomConfig{}
	LRCon.Code = 200
	LRCon.Data.LiveRoomConfigInfo.Lid = Res.Lid
	LRCon.Data.LiveRoomConfigInfo.LivePic = Res.LivePic
	LRCon.Data.LiveRoomConfigInfo.Danmu = Res.Danmu
	LRCon.Data.LiveRoomConfigInfo.Chat = Res.Chat
	LRCon.Data.LiveRoomConfigInfo.Share = Res.Share
	LRCon.Data.LiveRoomConfigInfo.ShareText = Res.ShareText
	LRCon.Data.LiveRoomConfigInfo.Advertisement = Res.Advertisement
	LRCon.Data.LiveRoomConfigInfo.AdJumpUrl = Res.AdJumpUrl
	LRCon.Data.LiveRoomConfigInfo.AdPicUrl = Res.AdPicUrl
	LRCon.Data.LiveRoomConfigInfo.AdText = Res.AdText

	if resp, err := json.Marshal(LRCon); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp),200)
	}
}

func GetLRConfigByLid(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { //获取直播界面信息
	//cid := ps.ByName("cid")
	vars := r.URL.Query()
	//aid := vars.Get("aid")
	lid := vars.Get("lid")

	Res, err := dbop.RetrieveLRConfigByLid(lid)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	LRCon := &defs.LiveRoomConfig{}
	LRCon.Code = 200
	LRCon.Data.LiveRoomConfigInfo.Lid = Res.Lid
	LRCon.Data.LiveRoomConfigInfo.LivePic = Res.LivePic
	LRCon.Data.LiveRoomConfigInfo.Danmu = Res.Danmu
	LRCon.Data.LiveRoomConfigInfo.Chat = Res.Chat
	LRCon.Data.LiveRoomConfigInfo.Share = Res.Share
	LRCon.Data.LiveRoomConfigInfo.ShareText = Res.ShareText
	LRCon.Data.LiveRoomConfigInfo.Advertisement = Res.Advertisement
	LRCon.Data.LiveRoomConfigInfo.AdJumpUrl = Res.AdJumpUrl
	LRCon.Data.LiveRoomConfigInfo.AdPicUrl = Res.AdPicUrl
	LRCon.Data.LiveRoomConfigInfo.AdText = Res.AdText

	if resp, err := json.Marshal(LRCon); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp),200)
	}
}

