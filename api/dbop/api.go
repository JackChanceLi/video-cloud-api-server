package dbop

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-api-server/api/defs"
	"go-api-server/api/utils"
	"log"
	"strings"
	"time"
)

//获取当前的时间
func getCurrentTime() (string, error) {
	tNow := time.Now()
	timeNow := tNow.Format("2006-01-02 15:04:05")
	return timeNow, nil
}

func UserLogin(email string) (*defs.UserInformation, string, error) {
	actOut,err := dbConn.Prepare("SELECT cid, uname, password, auth from company WHERE email = ?")
	if err != nil {
		log.Printf("%s", err)
		return nil, "", err
	}

	var cid, uname, auth, password, aid string
	UI := &defs.UserInformation{}
	err = actOut.QueryRow(email).Scan(&cid, &uname, &password, &auth)
	if err != nil && err != sql.ErrNoRows {
		return nil, "", err
	}
	if err == sql.ErrNoRows { //表示用户为普通管理员用户，在另一个表中重新进行查询
		actOut1,err := dbConn.Prepare("SELECT aid, cid, uname, password, auth from admin WHERE email = ?")
		if err != nil {
			log.Printf("%s", err)
			return nil, "", err
		}
		err = actOut1.QueryRow(email).Scan(&aid, &cid, &uname, &password, &auth)
		if err != nil && err != sql.ErrNoRows {
			return nil, "", err
		}

		defer actOut1.Close()
		UI.Cid = cid
		UI.Aid = aid
		UI.Name = uname
		UI.Auth = auth
		UI.Email = email

		return UI, password, nil

	}
	defer actOut.Close()

	UI.Cid = cid
	UI.Aid = cid
	UI.Name = uname
	UI.Auth = auth
	UI.Email = email

	return UI, password, nil

}
func IsEmailRegister(email string) (bool, string, error) {
	actOut,err := dbConn.Prepare("SELECT uname from company WHERE email = ?")
	if err != nil { //表示查询出错
		log.Printf("%s", err)
		return true, "", err
	}
	var uname string
	err = actOut.QueryRow(email).Scan(&uname)
	if err!= nil && err == sql.ErrNoRows {  //没有此邮箱，表示该邮箱未注册
		return false, "", err
	}
	if err != nil {
		return true, "", err  //表示查询过程出现了错误
	}
	defer actOut.Close()
	return false, uname, nil    //表示查询有结果，邮箱已注册
}

func UserRegister(uname, email, password string) error {
	cid, _ := utils.NewUUID()
	log.Printf("User Register uid:%s",cid)
	avrtalUrl := "http://pic-cloud-bupt.oss-cn-beijing.aliyuncs.com/3QbpEyjbGT.png"
	auth := defs.AuthFir
	var authList string
	for i:= 0; i < len(auth) - 1; i++ {
		authList += auth[i] + ";"
	}
	authList += auth[len(auth) - 1]
	desc := ""
	actIns,err := dbConn.Prepare("INSERT  INTO company (cid, uname, email, password, auth, register_date, " +
		"avtar_url, descp) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Println(err)
		return err
	}
	tNow, _ := getCurrentTime()
	_, err = actIns.Exec(cid, uname, email, password, authList, tNow, avrtalUrl, desc)
	if err != nil {
		return err
	}

	defer actIns.Close()
	return nil
}

func GetUserInfomation(cid string, isCom int)(*defs.UserData, error) {
	var uname, email, auth, avtarUrl, descp string
	if isCom == 1 {
		//在company表中进行查询
		actOut1,err := dbConn.Prepare("SELECT  uname, email, auth, avtar_url, descp from company WHERE cid = ?")
		if err != nil {
			log.Printf("%s", err)
			return nil, err
		}
		err = actOut1.QueryRow(cid).Scan(&uname, &email, &auth, &avtarUrl, &descp)
		if err != nil{
			log.Println(err)
			return nil, err
		}
		Auth := strings.Split(auth, ";")
		cinfo := &defs.UserData{}
		cinfo.Name = uname
		cinfo.Auth = Auth
		cinfo.AvtarUrl = avtarUrl
		cinfo.Email  = email
		cinfo.Desc = descp
		return cinfo, nil

	} else {
		//在admin表中进行查询
		actOut2,err := dbConn.Prepare("SELECT  uname, email, auth, avtar_url, descp from company WHERE cid = ?")
		if err != nil {
			log.Printf("%s", err)
			return nil, err
		}
		err = actOut2.QueryRow(cid).Scan(&uname, &email, &auth, &avtarUrl, &descp)
		if err != nil{
			return nil, err
		}
		Auth := strings.Split(auth, ";")
		cinfo := &defs.UserData{}
		cinfo.Name = uname
		cinfo.Auth = Auth
		cinfo.AvtarUrl = avtarUrl
		cinfo.Email  = email
		cinfo.Desc = descp
		return cinfo, nil
	}
}

func UpdateUserInfo(cid, uname, email, desc, avtarUrl string, isCom int) (*defs.UserData, error) {
	log.Println(desc)
	if isCom == 1{
		stmtUpa, err := dbConn.Prepare("UPDATE company SET uname = ?, email = ?, avtar_url = ?, descp = ? WHERE cid = ? ")
		if err != nil {
			log.Printf("%s",err)
			return nil, err
		}

		_,err = stmtUpa.Exec(uname, email, avtarUrl, desc, cid)
		if err != nil {
			log.Printf("%s",err)
			return nil, err
		}
		defer stmtUpa.Close()
		userInfo := &defs.UserData{}
		userInfo.AvtarUrl = avtarUrl
		userInfo.Name = uname
		userInfo.AvtarUrl = avtarUrl
		userInfo.Desc =desc
		return userInfo, nil
	} else {
		stmtUpa, err := dbConn.Prepare("UPDATE admin SET uname = ?, email = ?, descp = ?, avtar_url = ? WHERE aid = ? ")
		if err != nil {
			log.Printf("%s",err)
			return nil, err
		}

		_,err = stmtUpa.Exec(uname, email, desc, avtarUrl, cid)
		if err != nil {
			log.Printf("%s",err)
			return nil, err
		}
		defer stmtUpa.Close()
		userInfo := &defs.UserData{}
		userInfo.Name = uname
		userInfo.AvtarUrl = avtarUrl
		userInfo.Desc = desc
		userInfo.Email = email
		return userInfo, nil
	}
}
func IsEmailRegisetredByUpdate(email, cid string, isCom int) (bool, error) { //在修改个人信息时检测邮箱是否可用
	if isCom == 1 {
		actOut,err := dbConn.Prepare("SELECT cid from company WHERE email = ?")
		if err != nil { //表示查询出错
			log.Printf("%s", err)
			return false, err
		}
		var uid string
		err = actOut.QueryRow(email).Scan(&uid)
		if err!= nil && err == sql.ErrNoRows {  //没有此邮箱，表示该邮箱未注册
			return false, nil
		}
		if err != nil {
			return false, err  //表示查询过程出现了错误
		}
		defer actOut.Close()
		if uid == cid{ //表明此次修改没有更新邮箱，操作正常进行
			return false, nil
		}
		return true, nil    //表示查询有结果，邮箱已注册,此次修改失败
	} else {
		actOut,err := dbConn.Prepare("SELECT aid from company WHERE email = ?")
		if err != nil { //表示查询出错
			log.Printf("%s", err)
			return false, err
		}
		var uid string
		err = actOut.QueryRow(email).Scan(&uid)
		if err!= nil && err == sql.ErrNoRows {  //没有此邮箱，表示该邮箱未注册
			return false, nil
		}
		if err != nil {
			return false, err  //表示查询过程出现了错误
		}
		defer actOut.Close()
		if uid == cid{ //表明此次修改没有更新邮箱，操作正常进行
			return false, nil
		}
		return true, nil    //表示查询有结果，邮箱已注册,此次修改失败
	}
}
