package main

import (
	"encoding/json"
	"go-api-server/api/defs"
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, errResp defs.ErrorResponse) {
	//允许跨域
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Credentials","true")
	//允许请求类型
	w.Header().Set("Access-Control-Allow-Methods","*")
	//允许自定义头部
	w.Header().Set("Access-Control-Allow-Headers","*")
	w.Header().Set("Access-Control-Expose-Headers","*")

	w.WriteHeader(200)

	resStr, _ := json.Marshal(errResp)
	io.WriteString(w, string(resStr))
}

func sendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	//允许跨域
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Credentials","true")
	//允许请求类型
	w.Header().Set("Access-Control-Allow-Methods","*")
	//允许自定义头部
	w.Header().Set("Access-Control-Allow-Headers","*")
	w.Header().Set("Access-Control-Expose-Headers","*")

	w.WriteHeader(sc)
	io.WriteString(w, resp)
}