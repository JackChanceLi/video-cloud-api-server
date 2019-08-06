package main

import (
	"go-api-server/api/defs"
	"go-api-server/api/session"
	"net/http"
)

//http协议中自定义的header字段
var HEADER_FILED_SESSION = "X-Session-Id"
var HEADER_FILED_UNAME = "X-User-Name"
var URL_COLLECTION = [...]string {"/user/login", "/user/register"}

func validateUserSession(r *http.Request) bool {
	//将自定义的header字段加入鉴权过程
	sid := r.Header.Get(HEADER_FILED_SESSION)
	if len(sid) == 0 {
		return false
	}

	uname, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}
	r.Header.Add(HEADER_FILED_UNAME, uname)
	return true
}

func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HEADER_FILED_UNAME)
	if len(uname) == 0 {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}
	return true
}

func IsCheckSession(path string) bool {
	for i:= 0; i < len(URL_COLLECTION); i++ {
		if URL_COLLECTION[i] == path {
			return false
		}
	}
	return true
}
