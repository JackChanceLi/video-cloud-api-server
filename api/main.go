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

func main() {
	r := handler()
	mh := NewMiddleWareHandler(r)
	log.Printf("Server start1\n")
	http.ListenAndServe(":9000",mh)
}



