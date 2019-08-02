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

func UploadResourse(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { //上传文件
	cid := ps.ByName("cid")
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Http body read failed")
	}
	ubody := &defs.ResourseIdentity{}

	//解析包
	if err := json.Unmarshal(res, ubody); err != nil {
		fmt.Println(ubody)
		sendErrorResponse(w, defs.ErrRequestBodyParseFailed)
		return
	}
	fmt.Println(ubody)
	Res, err := dbop.UploadResourseByCom(ubody.Aid, cid, ubody.Name, ubody.Rtype, ubody.Size, ubody.Label)
	if  err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	ResInfo := &defs.Resourse{}
	ResInfo.Code = 200
	ResInfo.Data.ResourseInfo.Rid = Res.Rid
	ResInfo.Data.ResourseInfo.Aid = Res.Aid
	ResInfo.Data.ResourseInfo.Cid = Res.Cid
	ResInfo.Data.ResourseInfo.Name = Res.Name
	ResInfo.Data.ResourseInfo.Rtype = Res.Rtype
	ResInfo.Data.ResourseInfo.Size = Res.Size
	ResInfo.Data.ResourseInfo.Label = Res.Label
	ResInfo.Data.ResourseInfo.Time = Res.Time

	if resp, err := json.Marshal(ResInfo); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp),200)
	}
}

func UpdateResourse(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {//更新文件资源信息
	//cid := ps.ByName("cid")//获取cid
	au := r.URL.Query()
	aid := au.Get("aid")//获取aid
	log.Printf("Aid value is [%s]\n", aid)

	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Http body read failed")
	}
	ubody := &defs.ResourseIdentity{}

	//解析包
	if err := json.Unmarshal(res, ubody); err != nil {
		fmt.Println(ubody)
		sendErrorResponse(w, defs.ErrRequestBodyParseFailed)//解析错误
		return
	}
	fmt.Println(ubody)

	Res, err := dbop.UpdateResourse(ubody.Rid, ubody.Name, ubody.Label)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	ResInfo := &defs.Resourse{}
	ResInfo.Code = 200
	ResInfo.Data.ResourseInfo.Rid = Res.Rid
	ResInfo.Data.ResourseInfo.Aid = Res.Aid
	ResInfo.Data.ResourseInfo.Cid = Res.Cid
	ResInfo.Data.ResourseInfo.Name = Res.Name
	ResInfo.Data.ResourseInfo.Rtype = Res.Rtype
	ResInfo.Data.ResourseInfo.Size = Res.Size
	ResInfo.Data.ResourseInfo.Label = Res.Label
	ResInfo.Data.ResourseInfo.Time = Res.Time

	if resp, err := json.Marshal(ResInfo); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp),200)
	}
}

func GetResourses(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { //得到某个公司的所有资源文件信息
	cid := ps.ByName("cid")
	vars := r.URL.Query()
	aid := vars.Get("aid")
	log.Printf("%s", aid)

	res, err := dbop.RetrieveResourseByCid(cid)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
	}
	fmt.Println(res)

	ResList := &defs.ResourseList{}
	ResList.Code = 200
	ResList.Data = res
	if resp, err := json.Marshal(ResList); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp),200)
	}
}

func DeleteResourse(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//cid := ps.ByName("cid")//获取cid
	au := r.URL.Query()
	aid := au.Get("aid")//获取aid
	rid := au.Get("rid")//获取rid
	log.Printf("Aid value is [%s]\n", aid)
	log.Printf("Lid value is [%s]\n", rid)

	err := dbop.DeleteResourse(rid)
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

func GetResourseByRid(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { //获取某资源文件全部信息
	cid := ps.ByName("cid")
	vars := r.URL.Query()
	rid := vars.Get("rid")

	Res, err := dbop.RetrieveResourseByRid(rid)
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
	ResInfo := &defs.Resourse{}
	ResInfo.Code = 200
	ResInfo.Data.ResourseInfo.Rid = Res.Rid
	ResInfo.Data.ResourseInfo.Aid = Res.Aid
	ResInfo.Data.ResourseInfo.Cid = Res.Cid
	ResInfo.Data.ResourseInfo.Name = Res.Name
	ResInfo.Data.ResourseInfo.Rtype = Res.Rtype
	ResInfo.Data.ResourseInfo.Size = Res.Size
	ResInfo.Data.ResourseInfo.Label = Res.Label
	ResInfo.Data.ResourseInfo.Time = Res.Time

	if resp, err := json.Marshal(ResInfo); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp),200)
	}

}

