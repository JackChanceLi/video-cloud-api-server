package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-api-server/api/defs"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestMain(m *testing.M) {
	//dbClear()
	m.Run()
	//dbClear()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("register", testUserRegister)
	t.Run("login", testUserLogin)
}

func testUserRegister(t *testing.T) {
	//初始设置
	url := "http://127.0.0.1:9090/user/register"
	contentType := "application/json;charset=utf-8"
	//设置发送包内容
	var userID defs.UserIdentity
	userID.UserName = "zhangfanhao"
	userID.Password = "789789"
	userID.Email = "505608470@qq.com"

	b, err := json.Marshal(userID)
	if err != nil {
		fmt.Println("json format error:", err)
		return
	}

	body := bytes.NewBuffer(b)//包的body部分

	resp, err := http.Post(url, contentType, body)
	if err != nil {
		fmt.Println("Post failed:", err)
		return
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read failed:", err)
		return
	}

	fmt.Println("header:", resp.Header)
	fmt.Println("content:", string(content))
}

func testUserLogin(t *testing.T) {
	fmt.Println("Login success")
}
func testRecieveToken(t *testing.T) {
	url := "http://114.116.180.115:9002/ljczjnjyl"
	//contentType := "application/json;charset=utf-8"
	resp, err := http.Get(url)
	if err != nil {
		t.Errorf("Get failed:%v", err)
		return
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read failed:", err)
		return
	}

	fmt.Println( resp.Body)
	fmt.Println("content:", string(content))
}