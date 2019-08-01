package dbop

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-micro-server/api/defs"
	"go-micro-server/api/utils"
	"log"
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
func IsEmailRegister(email string) (bool, error) {
	actOut,err := dbConn.Prepare("SELECT cid from company WHERE email = ?")
	if err != nil { //表示查询出错
		log.Printf("%s", err)
		return true, err
	}
	var cid string
	err = actOut.QueryRow(email).Scan(&cid)
	if err!= nil && err == sql.ErrNoRows {  //没有此邮箱，表示该邮箱未注册
		return false, err
	}
	if err != nil {
		return true, err  //表示查询过程出现了错误
	}
	defer actOut.Close()
	return false, nil    //表示查询有结果，邮箱已注册
}
func userRegister(userName string, email string, password string, role int) error {  //old version:register user
	uid, _ := utils.NewUUID()
	log.Printf("uid:%s",uid)
	actIns,err := dbConn.Prepare("INSERT  INTO user_information (user_id, user_name, register_date, email, user_passwd, role) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
    tNow, _ := getCurrentTime()
	_, err = actIns.Exec(uid, userName, tNow, email, password, role)
	if err != nil {
		return err
	}
	defer actIns.Close()
	return nil
}

func UserRegister(uname, email, password string) error {
	cid, _ := utils.NewUUID()
	log.Printf("User Register uid:%s",cid)
	actIns,err := dbConn.Prepare("INSERT  INTO company (cid, uname, email, password, auth, register_date) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	tNow, _ := getCurrentTime()
	auth := "all"
	_, err = actIns.Exec(cid, uname, email, password, auth, tNow)
	if err != nil {
		return err
	}

	defer actIns.Close()
	return nil

}
