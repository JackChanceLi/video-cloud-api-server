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

func InsertLRIntro(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { //插入直播引导界面信息
	//cid := ps.ByName("cid")
	au := r.URL.Query()
	aid := au.Get("aid")//获取aid
	log.Printf("Aid value is [%s]\n", aid)

	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Http body read failed")
	}
	ubody := &defs.LiveRoomIntroIdentity{}

	//解析包
	if err := json.Unmarshal(res, ubody); err != nil {
		fmt.Println(ubody)
		sendErrorResponse(w, defs.ErrRequestBodyParseFailed)
		return
	}
	fmt.Println(ubody)
	Res, err := dbop.InsertLRIntroByCom(ubody.Lid, ubody.Qorder, ubody.Prepic)
	if  err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	LRIn := &defs.LiveRoomIntro{}
	LRIn.Code = 200
	LRIn.Data.LiveRoomIntroInfo.Lid = Res.Lid
	LRIn.Data.LiveRoomIntroInfo.Qorder = Res.Qorder
	LRIn.Data.LiveRoomIntroInfo.Prepic = Res.Prepic

	if resp, err := json.Marshal(LRIn); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp),200)
	}
}

func UpdateLRIntro(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {//更新直播引导界面信息
	//cid := ps.ByName("cid")//获取cid
	au := r.URL.Query()
	aid := au.Get("aid")//获取aid
	log.Printf("Aid value is [%s]\n", aid)

	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Http body read failed")
	}
	ubody := &defs.LiveRoomIntroIdentity{}

	//解析包
	if err := json.Unmarshal(res, ubody); err != nil {
		fmt.Println(ubody)
		sendErrorResponse(w, defs.ErrRequestBodyParseFailed)//解析错误
		return
	}
	fmt.Println(ubody)

	Res, err := dbop.UpdateLRIntro(ubody.Lid, ubody.Qorder, ubody.Prepic)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	LRIn := &defs.LiveRoomIntro{}
	LRIn.Code = 200
	LRIn.Data.LiveRoomIntroInfo.Lid = Res.Lid
	LRIn.Data.LiveRoomIntroInfo.Qorder = Res.Qorder
	LRIn.Data.LiveRoomIntroInfo.Prepic = Res.Prepic

	if resp, err := json.Marshal(LRIn); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp),200)
	}
}

func GetLRIntroByLid(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { //获取直播引导界面信息
	//cid := ps.ByName("cid")
	vars := r.URL.Query()
	//aid := vars.Get("aid")
	lid := vars.Get("lid")

	Res, err := dbop.RetrieveLRIntroByLid(lid)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	LRIn := &defs.LiveRoomIntro{}
	LRIn.Code = 200
	LRIn.Data.LiveRoomIntroInfo.Lid = Res.Lid
	LRIn.Data.LiveRoomIntroInfo.Qorder = Res.Qorder
	LRIn.Data.LiveRoomIntroInfo.Prepic = Res.Prepic

	if resp, err := json.Marshal(LRIn); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp),200)
	}

}

