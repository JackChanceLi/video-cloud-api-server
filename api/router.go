package main

import (
	"github.com/julienschmidt/httprouter"
)

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
	router.POST("/com/:cid/liveroom/condition/", UpdateLRCondition)
	//服务设置管理
	router.POST("/com/:cid/liveroom/quality/", InsertLRQuality)
	router.PUT("/com/:cid/liveroom/quality/", UpdateLRQuality)
	router.GET("/com/:cid/liveroom/quality/", GetLRQualityByLid)
	//版本安全设置管理
	router.POST("/com/:cid/liveroom/safe/", InsertLRSafe)
	router.PUT("/com/:cid/liveroom/safe/", UpdateLRSafe)
	router.GET("/com/:cid/liveroom/safe/", GetLRSafeByLid)
	//权限安全设置管理
	router.POST("/com/:cid/liveroom/auth_safe/", InsertLRAuthSafe)
	router.PUT("/com/:cid/liveroom/auth_safe/", UpdateLRAuthSafe)
	router.GET("/com/:cid/liveroom/auth_safe_black/", GetLRAuthSafeBlackListByLid)
	router.GET("/com/:cid/liveroom/auth_safe_white/", GetLRAuthSafeWhiteListByLid)

	//获取全部权限信息
	router.GET("/com/:cid/liveroom/all_config", GetLiveRoomAllConfig)

	return router
}
