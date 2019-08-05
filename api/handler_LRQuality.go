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

func InsertLRQuality(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { //插入直播引导界面信息
	//cid := ps.ByName("cid")
	au := r.URL.Query()
	aid := au.Get("aid")//获取aid
	log.Printf("Aid value is [%s]\n", aid)

	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Http body read failed")
	}
	ubody := &defs.LiveRoomQualityIdentity{}

	//解析包
	if err := json.Unmarshal(res, ubody); err != nil {
		fmt.Println(ubody)
		sendErrorResponse(w, defs.ErrRequestBodyParseFailed)
		return
	}
	fmt.Println(ubody)
	Res, err := dbop.InsertLRQualityByCom(ubody.Lid, ubody.Delay, ubody.Transcode, ubody.TranscodeType, ubody.Record, ubody.RecordType)
	if  err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	LRQua := &defs.LiveRoomQuality{}
	LRQua.Code = 200
	LRQua.Data.LiveRoomQualityInfo.Lid = Res.Lid
	LRQua.Data.LiveRoomQualityInfo.Delay = Res.Delay
	LRQua.Data.LiveRoomQualityInfo.Transcode = Res.Transcode
	LRQua.Data.LiveRoomQualityInfo.TranscodeType = Res.TranscodeType
	LRQua.Data.LiveRoomQualityInfo.Record = Res.Record
	LRQua.Data.LiveRoomQualityInfo.RecordType = Res.RecordType

	if resp, err := json.Marshal(LRQua); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp),200)
	}
}

func UpdateLRQuality(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {//更新直播引导界面信息
	//cid := ps.ByName("cid")//获取cid
	au := r.URL.Query()
	aid := au.Get("aid")//获取aid
	log.Printf("Aid value is [%s]\n", aid)

	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Http body read failed")
	}
	ubody := &defs.LiveRoomQualityIdentity{}

	//解析包
	if err := json.Unmarshal(res, ubody); err != nil {
		fmt.Println(ubody)
		sendErrorResponse(w, defs.ErrRequestBodyParseFailed)//解析错误
		return
	}
	fmt.Println(ubody)

	Res, err := dbop.UpdateLRQuality(ubody.Lid, ubody.Delay, ubody.Transcode, ubody.TranscodeType, ubody.Record, ubody.RecordType)
	if  err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	LRQua := &defs.LiveRoomQuality{}
	LRQua.Code = 200
	LRQua.Data.LiveRoomQualityInfo.Lid = Res.Lid
	LRQua.Data.LiveRoomQualityInfo.Delay = Res.Delay
	LRQua.Data.LiveRoomQualityInfo.Transcode = Res.Transcode
	LRQua.Data.LiveRoomQualityInfo.TranscodeType = Res.TranscodeType
	LRQua.Data.LiveRoomQualityInfo.Record = Res.Record
	LRQua.Data.LiveRoomQualityInfo.RecordType = Res.RecordType

	if resp, err := json.Marshal(LRQua); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp),200)
	}
}

func GetLRQualityByLid(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { //获取直播引导界面信息
	//cid := ps.ByName("cid")
	vars := r.URL.Query()
	//aid := vars.Get("aid")
	lid := vars.Get("lid")

	Res, err := dbop.RetrieveLRQualityByLid(lid)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	LRQua := &defs.LiveRoomQuality{}
	LRQua.Code = 200
	LRQua.Data.LiveRoomQualityInfo.Lid = Res.Lid
	LRQua.Data.LiveRoomQualityInfo.Delay = Res.Delay
	LRQua.Data.LiveRoomQualityInfo.Transcode = Res.Transcode
	LRQua.Data.LiveRoomQualityInfo.TranscodeType = Res.TranscodeType
	LRQua.Data.LiveRoomQualityInfo.Record = Res.Record
	LRQua.Data.LiveRoomQualityInfo.RecordType = Res.RecordType

	if resp, err := json.Marshal(LRQua); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp),200)
	}

}

