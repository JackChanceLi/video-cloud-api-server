package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
)

func GetToken(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	url := "http://114.116.180.115:9002/ljczjnjyl"
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Get failed:%v", err)
		return
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Read failed:%v", err)
		return
	}
	fmt.Println(string(content))
	sendNormalResponse(w, string(content),200)

}