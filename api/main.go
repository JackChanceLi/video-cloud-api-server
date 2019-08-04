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
	if r.Method == "OPTIONS" {
		NormalHandler(w, r)
		return
	}
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
	router.DELETE("/com/:cid/liverooms/", DeleteLiveRoom)
	router.GET ("/com/:cid/liveroom/", GetLiveRoomByLid)
	//文件资源管理
	router.POST("/com/:cid/resourses/", UploadResourse)
	router.PUT("/com/:cid/resourses/", UpdateResourse)
	router.GET("/com/:cid/resourses/", GetResourses)
	router.DELETE("/com/:cid/resourses/", DeleteResourse)
	router.GET ("/com/:cid/resourse/", GetResourseByRid)
	//直播引导界面信息管理
	router.POST("/com/:cid/liveroom/intro/", InsertLRIntro)
	router.PUT("/com/:cid/liveroom/intro/", UpdateLRIntro)
	router.GET("/com/:cid/liveroom/intro/", GetLRIntroByLid)
	//直播信息界面管理
	router.POST("/com/:cid/liveroom/config/", InsertLRConfig)
	router.PUT("/com/:cid/liveroom/config/", UpdateLRConfig)
	router.GET("/com/:cid/liveroom/config/", GetLRConfigByLid)
	//直播观看条件管理



	//服务设置管理



	//版本安全设置管理


	//router.OPTIONS("/", NormalHandler)

	//权限安全设置管理
	router.POST("/com/:cid/liveroom/auth_safe/", InsertLRAuthSafe)
	router.PUT("/com/:cid/liveroom/auth_safe/", UpdateLRAuthSafe)
	router.GET("/com/:cid/liveroom/auth_safe_black/", GetLRAuthSafeBlackListByLid)
	router.GET("/com/:cid/liveroom/auth_safe_white/", GetLRAuthSafeWhiteListByLid)

	return router
}

func main() {
	r := handler()
	mh := NewMiddleWareHandler(r)
	log.Printf("Server start1\n")
	http.ListenAndServe(":9000",mh)
}



