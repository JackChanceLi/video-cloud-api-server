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

func InsertLRAuthSafe(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { //插入直播间权限安全信息
	//cid := ps.ByName("cid")
	au := r.URL.Query()
	aid := au.Get("aid")//获取aid
	log.Printf("Aid value is [%s]\n", aid)

	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Http body read failed")
	}
	ubody := &defs.LiveRoomAuthSafeIdentity{}

	//解析包
	if err := json.Unmarshal(res, ubody); err != nil {
		fmt.Println(ubody)
		sendErrorResponse(w, defs.ErrRequestBodyParseFailed)
		return
	}
	fmt.Println(ubody)
	Res, err := dbop.UpdateLRAuthSafe(ubody.Lid, ubody.WhiteSiteList, ubody.BlackSiteList)
	if  err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	LRAS := &defs.LiveRoomAuthSafe{}
	LRAS.Code = 200
	LRAS.Data.LiveRoomAuthSafeInfo.Lid = Res.Lid
	LRAS.Data.LiveRoomAuthSafeInfo.WhiteSiteList = Res.WhiteSiteList
	LRAS.Data.LiveRoomAuthSafeInfo.BlackSiteList = Res.BlackSiteList

	if resp, err := json.Marshal(LRAS); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp),200)
	}
}

func UpdateLRAuthSafe(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {//更新直播间权限安全信息
	//cid := ps.ByName("cid")//获取cid
	au := r.URL.Query()
	aid := au.Get("aid")//获取aid
	log.Printf("Aid value is [%s]\n", aid)

	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Http body read failed")
	}
	ubody := &defs.LiveRoomAuthSafeIdentity{}

	//解析包
	if err := json.Unmarshal(res, ubody); err != nil {
		fmt.Println(ubody)
		sendErrorResponse(w, defs.ErrRequestBodyParseFailed)//解析错误
		return
	}
	fmt.Println(ubody)

	Res, err := dbop.UpdateLRAuthSafe(ubody.Lid, ubody.WhiteSiteList, ubody.BlackSiteList)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	LRAS := &defs.LiveRoomAuthSafe{}
	LRAS.Code = 200
	LRAS.Data.LiveRoomAuthSafeInfo.Lid = Res.Lid
	LRAS.Data.LiveRoomAuthSafeInfo.WhiteSiteList = Res.WhiteSiteList
	LRAS.Data.LiveRoomAuthSafeInfo.BlackSiteList = Res.BlackSiteList

	if resp, err := json.Marshal(LRAS); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp),200)
	}
}

func GetLRAuthSafeByLid(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { //获取直播间权限安全黑白名单
	//cid := ps.ByName("cid")
	vars := r.URL.Query()
	//aid := vars.Get("aid")
	lid := vars.Get("lid")

	Res, err := dbop.RetrieveLRAuthSafeList(lid)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	LRASWL := &defs.LiveRoomAuthSafe{}
	LRASWL.Code = 200
	LRASWL.Data.LiveRoomAuthSafeInfo.Lid = Res.Lid
	LRASWL.Data.LiveRoomAuthSafeInfo.WhiteSiteList = Res.WhiteSiteList
	LRASWL.Data.LiveRoomAuthSafeInfo.BlackSiteList = Res.BlackSiteList

	if resp, err := json.Marshal(LRASWL); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp),200)
	}

}
