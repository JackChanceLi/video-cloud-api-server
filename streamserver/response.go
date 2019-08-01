package main

import (
	"io"
	"net/http"
)

func SendErrorResponse(w http.ResponseWriter, sc int, errMsg string) {

	w.WriteHeader(sc)
	io.WriteString(w, errMsg)
}

func SendNormalResponse(w http.ResponseWriter, sc int, norMsg string) {
	w.WriteHeader(sc)
	io.WriteString(w, norMsg)
}