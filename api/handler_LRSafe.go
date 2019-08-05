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

func InsertLRSafe(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { //插入直播间版本安全设置信息
	//cid := ps.ByName("cid")
	au := r.URL.Query()
	aid := au.Get("aid")//获取aid
	log.Printf("Aid value is [%s]\n", aid)

	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Http body read failed")
	}
	ubody := &defs.LiveRoomSafeIdentity{}

	//解析包
	if err := json.Unmarshal(res, ubody); err != nil {
		fmt.Println(ubody)
		sendErrorResponse(w, defs.ErrRequestBodyParseFailed)
		return
	}
	fmt.Println(ubody)
	Res, err := dbop.InsertLRSafeByCom(ubody.Lid, ubody.Logo, ubody.LogoUrl, ubody.LogoPosition, ubody.LogoTransparency, ubody.Lamp, ubody.LampType, ubody.LampText, ubody.LampFontSize, ubody.LampTransparency)
	if  err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	LRSf := &defs.LiveRoomSafe{}
	LRSf.Code = 200
	LRSf.Data.LiveRoomSafeInfo.Lid = Res.Lid
	LRSf.Data.LiveRoomSafeInfo.Logo = Res.Logo
	LRSf.Data.LiveRoomSafeInfo.LogoUrl = Res.LogoUrl
	LRSf.Data.LiveRoomSafeInfo.LogoPosition = Res.LogoPosition
	LRSf.Data.LiveRoomSafeInfo.LogoTransparency = Res.LogoTransparency
	LRSf.Data.LiveRoomSafeInfo.Lamp = Res.Lamp
	LRSf.Data.LiveRoomSafeInfo.LampType = Res.LampType
	LRSf.Data.LiveRoomSafeInfo.LampText = Res.LampText
	LRSf.Data.LiveRoomSafeInfo.LampFontSize = Res.LampFontSize
	LRSf.Data.LiveRoomSafeInfo.LampTransparency = Res.LampTransparency

	if resp, err := json.Marshal(LRSf); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp),200)
	}
}

func UpdateLRSafe(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {//更新直播间版本安全设置信息
	//cid := ps.ByName("cid")//获取cid
	au := r.URL.Query()
	aid := au.Get("aid")//获取aid
	log.Printf("Aid value is [%s]\n", aid)

	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Http body read failed")
	}
	ubody := &defs.LiveRoomSafeIdentity{}

	//解析包
	if err := json.Unmarshal(res, ubody); err != nil {
		fmt.Println(ubody)
		sendErrorResponse(w, defs.ErrRequestBodyParseFailed)//解析错误
		return
	}
	fmt.Println(ubody)

	Res, err := dbop.UpdateLRSafe(ubody.Lid, ubody.Logo, ubody.LogoUrl, ubody.LogoPosition, ubody.LogoTransparency, ubody.Lamp, ubody.LampType, ubody.LampText, ubody.LampFontSize, ubody.LampTransparency)
	if  err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	LRSf := &defs.LiveRoomSafe{}
	LRSf.Code = 200
	LRSf.Data.LiveRoomSafeInfo.Lid = Res.Lid
	LRSf.Data.LiveRoomSafeInfo.Logo = Res.Logo
	LRSf.Data.LiveRoomSafeInfo.LogoUrl = Res.LogoUrl
	LRSf.Data.LiveRoomSafeInfo.LogoPosition = Res.LogoPosition
	LRSf.Data.LiveRoomSafeInfo.LogoTransparency = Res.LogoTransparency
	LRSf.Data.LiveRoomSafeInfo.Lamp = Res.Lamp
	LRSf.Data.LiveRoomSafeInfo.LampType = Res.LampType
	LRSf.Data.LiveRoomSafeInfo.LampText = Res.LampText
	LRSf.Data.LiveRoomSafeInfo.LampFontSize = Res.LampFontSize
	LRSf.Data.LiveRoomSafeInfo.LampTransparency = Res.LampTransparency

	if resp, err := json.Marshal(LRSf); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp),200)
	}
}

func GetLRSafeByLid(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { //获取直播间版本安全设置信息
	//cid := ps.ByName("cid")
	vars := r.URL.Query()
	//aid := vars.Get("aid")
	lid := vars.Get("lid")

	Res, err := dbop.RetrieveLRSafeByLid(lid)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	LRSf := &defs.LiveRoomSafe{}
	LRSf.Code = 200
	LRSf.Data.LiveRoomSafeInfo.Lid = Res.Lid
	LRSf.Data.LiveRoomSafeInfo.Logo = Res.Logo
	LRSf.Data.LiveRoomSafeInfo.LogoUrl = Res.LogoUrl
	LRSf.Data.LiveRoomSafeInfo.LogoPosition = Res.LogoPosition
	LRSf.Data.LiveRoomSafeInfo.LogoTransparency = Res.LogoTransparency
	LRSf.Data.LiveRoomSafeInfo.Lamp = Res.Lamp
	LRSf.Data.LiveRoomSafeInfo.LampType = Res.LampType
	LRSf.Data.LiveRoomSafeInfo.LampText = Res.LampText
	LRSf.Data.LiveRoomSafeInfo.LampFontSize = Res.LampFontSize
	LRSf.Data.LiveRoomSafeInfo.LampTransparency = Res.LampTransparency

	if resp, err := json.Marshal(LRSf); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp),200)
	}
}
