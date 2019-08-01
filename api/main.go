package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"

)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	validateUserSession(r)
	m.r.ServeHTTP(w, r)
}

func handler () *httprouter.Router {
	router := httprouter.New()
	//普通用户管理
	router.POST("/user/register", Register)
	router.POST("/user/login",LoginByMail)
	//直播间管理
	router.POST("/com/:cid/liverooms/", CreateLiveRoom)
	router.PUT("/com/:cid/liverooms/", UpdateLiveRoom)
	router.GET("/com/:cid/liverooms/", GetLiveRooms)
	router.DELETE("/com/:cid/liverooms/",DeleteLiveRoom)
	router.GET ("/com/:cid/liveroom/", GetLiveRoomByLid)

	return router
}

func main() {
	r := handler()
	mh := NewMiddleWareHandler(r)
	log.Printf("Server start1\n")
	http.ListenAndServe(":9000",mh)

}

