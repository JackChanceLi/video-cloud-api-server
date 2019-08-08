package dbop

import (
	"database/sql"
	"fmt"
	"log"

	"testing"
)

func dbClear(){
	dbConn.Exec("truncate user_information")
	dbConn.Exec(("truncate user_identity_information"))
}

func TestMain(m *testing.M) {
	//dbClear()
	m.Run()
	//dbClear()
}

func TestUserWorkFlow(t *testing.T) {
	//t.Run("register", testUserRegister)
	//t.Run("login", testUserLogin)
	//t.Run("time",testGetCurrentTime)
	t.Run("latest_register",testNewUserRegister)
	t.Run("insert", testCreateLiveRoom)
	t.Run("delete", testDeleteLiveRoom)
	//t.Run("update", testUpdateLiveRoom)
	t.Run("retrieve", testRetrieveLiveRoomByCid)
}

func testUserRegister(t *testing.T) {  // testing old register function
	err := userRegister("zheng","1258@gmail.com","000000",1)
	if err != nil {
		t.Errorf("Error of register: %v", err)
	}
}

func TestUserLogin(t *testing.T) {
	user, password,err := UserLogin("1023546080@qq.com")
	if password != "000000" {
		t.Errorf("Error of user login for wrong password:%s",password)
	}
	if err != nil {
		t.Errorf("Error of login: %v", err)
	}
	log.Println(user)
}

func testGetCurrentTime(t *testing.T) {
	tNow, err:= getCurrentTime()
	fmt.Printf("Time now:%s\n", tNow)
	if err != nil {
		t.Errorf("Error of time format:%s", err)
	}
}

func testNewUserRegister(t *testing.T) {  //testing latest register function
	err := UserRegister("ljc", "jackchance@163.com", "123456789")
	if err != nil {
		t.Errorf("Error of latest user register: %v",err)
	}
}

func testCreateLiveRoom(t *testing.T) {
	cid := "847a0971-db46-49b4-b30b19f9-5cc2c8e1"
	name := "cq"
	kind := 2
	size := 400
	startTime := "2019-07-26 15:11:35"
	endTime := "2019-07-26 16:30:45"
	_, err := CreateLiveRoomByCom(cid, name, kind, size, startTime, endTime)
	if err != nil {
		t.Errorf("Insert live_room failed\nerr:%v", err)
	} else {
		t.Log("Insert live_room success")
	}
}

func testDeleteLiveRoom(t *testing.T) {
	err := DeleteLiveRoom("d874e3f3-fab7-4e2b-a65478f0-f189e32a")
	if err != nil {
		t.Errorf("Delete live_room failed\nerr:%v", err)
	} else {
		t.Log("Delete live_room success")
	}
}

func testUpdateLiveRoom(t *testing.T) {
	nlid := "7a988d51-6712-4898-975cfe6f-6aff85dc"
	nname := "zjn"
	nkind := 1
	nsize := 300
	nstart_time := "2019-07-26 10:10:08"
	nend_time := "2019-07-26 12:18:35"
	npermission := 1
	nurl := ""
	_,err := UpdateLiveRoom(nlid, nname, nkind, nsize, nstart_time, nend_time, nurl, npermission)
	if err != nil {
		t.Errorf("Update live_room failed\nerr:%v", err)
	} else {
		t.Log("Update live_room success")
	}
}

func testRetrieveLiveRoomByCid(t *testing.T) {
	//Lri := &defs.LiveRoomIdentity{}
	m, cnt, err := SearchLiveRoomByCid("0fb1e47c-5e5c-4094-855d09ce-a598d3dc")
	if err != nil {
		t.Errorf("Retrieve live_room by cid failed\nerr:%v", err)
	} else {
		t.Log("Retrieve live_room by cid success")
		for i := 0; i < cnt; i++ {
			//log.Printf("Live_room lid:%s, cid:%s, name:%s, kind:%d, size:%d, start_time:%s, end_time:%s, push_url:%s, pull_hls_url:%s, pull_rtmp_url:%s, pull_http_flv_url:%s, display_url:%s, status:%d, permission:%s, create_time:%s", m.Load(i).lid, cid, name, kind, size, start_time, end_time, push_url, pull_hls_url, pull_rtmp_url, pull_http_flv_url, display_url, status, permission, create_time)
			fmt.Println(m.Load(i))
		}
	}
}

func TestIsEmailRegister(t *testing.T) {
	ok, err := IsEmailRegister("102354600@qq.com")
	if ok && err != nil {
		t.Errorf("Error of DB ops:%v", err)
	}
	if !ok && err == sql.ErrNoRows {
		t.Log("The email hasn't been registered")
	}
	if !ok && err == nil {
		t.Log("The email has been registered")
	}
}