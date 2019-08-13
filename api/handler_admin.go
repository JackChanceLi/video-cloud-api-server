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

func GetAuth(res_auth string) []string {
	if res_auth == "SA" {
		return defs.AuthFir
	} else if res_auth == "HA" {
		return defs.AuthSed
	} else if res_auth == "JA" {
		return defs.AuthTrd
	}
	return nil
}

func InsertAdmin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { //创建一个子管理员
	cid := ps.ByName("cid")//获取cid
	au := r.URL.Query()
	aid := au.Get("aid")//获取aid
	log.Printf("Aid value is [%s]\n", aid)

	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Http body read failed")
	}
	ubody := &defs.ResAdminIdentity{}

	//解析包
	if err := json.Unmarshal(res, ubody); err != nil {
		fmt.Println(ubody)
		sendErrorResponse(w, defs.ErrRequestBodyParseFailed)
		return
	}
	fmt.Println(ubody)
	//判断是否有权限
	ret, err := dbop.RetrieveAdminByAid(aid)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	var judge bool
	judge = false
	for i := 0; i < len(ret.Auth); i++ {
		if ret.Auth[i] == "c_admin" {
			judge = true
			break
		}
	}
	if aid != cid && judge == false { //没有超级权限
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}

	//有权限
	Res, err := dbop.InsertAdmin(cid, ubody.Uname, ubody.Password, ubody.Email, GetAuth(ubody.Auth), ubody.AvtarUrl, ubody.Descp)
	if  err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	Admin := &defs.Admin{}
	Admin.Code = 200
	Admin.Data.AdminInfo.Aid = Res.Aid
	Admin.Data.AdminInfo.Cid = Res.Cid
	Admin.Data.AdminInfo.Uname = Res.Uname
	Admin.Data.AdminInfo.Password = Res.Password
	Admin.Data.AdminInfo.Email = Res.Email
	Admin.Data.AdminInfo.Auth = Res.Auth
	Admin.Data.AdminInfo.AvtarUrl = Res.AvtarUrl
	Admin.Data.AdminInfo.Descp = Res.Descp

	if resp, err := json.Marshal(Admin); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp),200)
	}
}

func UpdateAdmin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {//更新子管理员信息
	cid := ps.ByName("cid")//获取cid
	au := r.URL.Query()
	aid := au.Get("aid")//获取aid
	said := au.Get("said")//获取所需更新的子管理员的aid
	log.Printf("Aid value is [%v]\n", aid)
	log.Printf("Said value is [%v]\n", said)

	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Http body read failed")
	}
	ubody := &defs.ResAdminIdentity{}

	//解析包
	if err := json.Unmarshal(res, ubody); err != nil {
		fmt.Println(ubody)
		sendErrorResponse(w, defs.ErrRequestBodyParseFailed)//解析错误
		return
	}
	fmt.Println(ubody)
	//判断是否有权限
	ret, err := dbop.RetrieveAdminByAid(aid)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	var judge bool
	judge = false
	for i := 0; i < len(ret.Auth); i++ {
		if ret.Auth[i] == "e_admin" {
			judge = true
			break
		}
	}
	if aid != cid && judge == false { //没有超级权限
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}

	//有权限
	Res, err := dbop.UpdateAdmin(said, ubody.Uname, ubody.Password, ubody.Email, GetAuth(ubody.Auth), ubody.AvtarUrl, ubody.Descp)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	Admin := &defs.Admin{}
	Admin.Code = 200
	Admin.Data.AdminInfo.Aid = Res.Aid
	Admin.Data.AdminInfo.Cid = Res.Cid
	Admin.Data.AdminInfo.Uname = Res.Uname
	Admin.Data.AdminInfo.Password = Res.Password
	Admin.Data.AdminInfo.Email = Res.Email
	Admin.Data.AdminInfo.Auth = Res.Auth
	Admin.Data.AdminInfo.AvtarUrl = Res.AvtarUrl
	Admin.Data.AdminInfo.Descp = Res.Descp

	if resp, err := json.Marshal(Admin); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp),200)
	}
}

func DeleteAdmin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cid := ps.ByName("cid")//获取cid
	au := r.URL.Query()
	aid := au.Get("aid")//获取aid
	said := au.Get("said")//获取所需删除的子管理员的aid
	log.Printf("Aid value is [%v]\n", aid)
	log.Printf("Said value is [%v]\n", said)

	//判断是否有权限
	ret, err := dbop.RetrieveAdminByAid(aid)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	var judge bool
	judge = false
	for i := 0; i < len(ret.Auth); i++ {
		if ret.Auth[i] == "d_admin" {
			judge = true
			break
		}
	}
	if aid != cid && judge == false { //没有超级权限
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}

	//有权限
	err = dbop.DeleteAdmin(said)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	su := defs.Resourse{Code:200}

	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	} else {
		sendNormalResponse(w, string(resp),200)
	}
}

func GetAdminByAid(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { //获取某子管理员全部信息
	cid := ps.ByName("cid")//获取cid
	au := r.URL.Query()
	aid := au.Get("aid")//获取aid
	log.Printf("Aid value is [%v]\n", aid)

	Res, err := dbop.RetrieveAdminByAid(aid)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	if Res.Cid == "" {  //等价于err == sql.ErrNoRows
		sendErrorResponse(w, defs.ErrNoRowsInDB)
		return
	}
	if Res.Cid != cid {
		sendErrorResponse(w, defs.ErrorNotAuthUserForRoom)
		return
	}

	Admin := &defs.Admin{}
	Admin.Code = 200
	Admin.Data.AdminInfo.Aid = Res.Aid
	Admin.Data.AdminInfo.Cid = Res.Cid
	Admin.Data.AdminInfo.Uname = Res.Uname
	Admin.Data.AdminInfo.Password = Res.Password
	Admin.Data.AdminInfo.Email = Res.Email
	Admin.Data.AdminInfo.Auth = Res.Auth
	Admin.Data.AdminInfo.AvtarUrl = Res.AvtarUrl
	Admin.Data.AdminInfo.Descp = Res.Descp

	if resp, err := json.Marshal(Admin); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp),200)
	}
}

