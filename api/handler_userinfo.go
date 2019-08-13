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
	"strconv"
)

func GetUserInfo (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	aid := ps.ByName("aid")
	vars := r.URL.Query()
	Com := vars.Get("is_com")
	isCom, err:= strconv.Atoi(Com)
	//res, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	log.Printf("Http body read failed")
	//}
	//ubody := &defs.ResourseIdentity{}
	//
	////解析包
	//if err := json.Unmarshal(res, ubody); err != nil {
	//	fmt.Println(ubody)
	//	sendErrorResponse(w, defs.ErrRequestBodyParseFailed)
	//	return
	//}
	//fmt.Println(ubody)
	info, err := dbop.GetUserInfomation(aid, isCom)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	userInfo := &defs.UserInfo{}
	userInfo.Code = 200
	userInfo.Data.Desc = info.Desc
	userInfo.Data.Email = info.Email
	userInfo.Data.AvtarUrl = info.AvtarUrl
	userInfo.Data.Auth = info.Auth
	userInfo.Data.Name = info.Name

	if resp, err := json.Marshal(userInfo); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp),200)
	}
	defer session.UpdateSession(aid)
}

func UpdateUserInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	aid := ps.ByName("aid")
	vars := r.URL.Query()
	Com := vars.Get("is_com")
	isCom, err:= strconv.Atoi(Com)
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Http body read failed")
	}
	ubody := &defs.UserData{}

	//解析包
	if err := json.Unmarshal(res, ubody); err != nil {
		fmt.Println(ubody)
		sendErrorResponse(w, defs.ErrRequestBodyParseFailed)
		return
	}
	fmt.Println(ubody)
	ok, er := dbop.IsEmailRegisetredByUpdate(ubody.Email, aid, isCom)
	if ok && er == nil {
		sendErrorResponse(w, defs.ErrorEmailRegistered)
		return
	}
	info, err := dbop.UpdateUserInfo(aid, ubody.Name, ubody.Email, ubody.Desc, ubody.AvtarUrl, isCom)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	userInfo := defs.UserInfo{}
	userInfo.Code = 200
	userInfo.Data.Email = info.Email
	userInfo.Data.Desc = info.Desc
	userInfo.Data.AvtarUrl = info.AvtarUrl
	userInfo.Data.Name = info.Name
	log.Println(userInfo)
	if resp, err := json.Marshal(userInfo); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp),200)
	}
	defer session.UpdateSession(aid)
}
