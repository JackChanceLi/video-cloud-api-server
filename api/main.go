package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-api-server/api/defs"
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
	if ok1 := IsCheckSession(r.URL.Path); ok1 {
		fmt.Println("Session check is needed")
		ok2 := validateUserSession(r)
		if !ok2 {
			sendErrorResponse(w, defs.ErrorNotAuthUser)
			return
		}
	}
	//fmt.Println(r.URL.Path)
	m.r.ServeHTTP(w, r)
}

func main() {
	r := handler()
	mh := NewMiddleWareHandler(r)
	log.Printf("Server start1\n")
	http.ListenAndServe(":9000",mh)
}



