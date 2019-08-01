package main

import (
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"

)

type middleWareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}

func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m
}

//劫持http请求，添加流控功能
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetConn() { //没有获取到令牌，对用的状态码为429
		SendErrorResponse(w, http.StatusTooManyRequests, "Too many requests.")
		return
	}
	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()  //结束的时候将令牌还回去
}
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/videos/:vid", StreamHandler)
	router.POST("/upload/:vid", UploadHandler)
	router.GET("/testpage", testPageHandler)
	return router
}

func main() {
	log.Println("Video server start")
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r, 2)
	http.ListenAndServe(":9001", mh)
}
