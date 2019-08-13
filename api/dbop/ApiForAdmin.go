package dbop

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-api-server/api/defs"
	"go-api-server/api/utils"
	"log"
	"strings"
)

func InsertAdmin(cid string, uname string, password string, email string, auth []string, avtar_url string, descp string) (*defs.AdminIdentity, error) {
	register_date, _ := getCurrentTime()
	stmtIns, err := dbConn.Prepare("INSERT INTO admin (aid, cid, uname, password, register_date, email, auth, avtar_url, descp) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	aid, _ := utils.NewUUID()
	var store string
	store = auth[0]
	for i := 1; i < len(auth); i++ {
		store += ";" + auth[i]
	}
	_,err = stmtIns.Exec(aid, cid, uname, password, register_date, email, store, avtar_url, descp)
	if err != nil {
		return nil, err
	}

	defer stmtIns.Close()
	AD := &defs.AdminIdentity{}
	AD.Aid = aid
	AD.Cid = cid
	AD.Uname = uname
	AD.Password = password
	AD.RegisterDate = register_date
	AD.Email = email
	AD.Auth = auth
	AD.AvtarUrl = avtar_url
	AD.Descp = descp

	return AD, nil
}

func DeleteAdmin(aid string) error {
	stmtOut, err := dbConn.Prepare("DELETE FROM admin WHERE aid = ?")
	if err != nil {
		log.Printf("%s",err)
		return err
	}

	if _, err := stmtOut.Query(aid); err != nil {
		return err
	}

	defer stmtOut.Close()
	return nil
}

func UpdateAdmin(aid string, uname string, password string, email string, auth []string, avtar_url string, descp string) (*defs.AdminIdentity, error) {
	stmtUpa, err := dbConn.Prepare("UPDATE admin SET uname = ?, password = ?, email = ?, auth = ?, avtar_url = ?, descp = ? WHERE aid = ?")
	if err != nil {
		log.Printf("%s",err)
		return nil, err
	}
	var store string
	store = auth[0]
	for i := 1; i < len(auth); i++ {
		store += ";" + auth[i]
	}

	_,err = stmtUpa.Exec(uname, password, email, store, avtar_url, descp, aid)
	if err != nil {
		return nil, err
	}

	defer stmtUpa.Close()
	Res, err := RetrieveAdminByAid(aid)
	if err != nil {
		return nil, err
	}
	res := &defs.AdminIdentity{}
	res.Aid = aid
	res.Cid = Res.Cid
	res.Uname = uname
	res.Password = password
	res.RegisterDate = Res.RegisterDate
	res.Email = email
	res.Auth = auth
	res.AvtarUrl = avtar_url
	res.Descp = descp
	return res, nil
}

func RetrieveAdminByAid(Aid string) (*defs.AdminIdentity, error) {
	stmtOut, err := dbConn.Prepare("SELECT cid, uname, password, register_date, email, auth, avtar_url, descp FROM admin WHERE aid = ?")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	var cid, uname, password, register_date, email, aauth, avtar_url, descp string
	var auth []string
	stmtOut.QueryRow(Aid).Scan(&cid, &uname, &password, &register_date, &email, &aauth, &avtar_url, &descp)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	auth = strings.Split(aauth, ";")
	res := &defs.AdminIdentity{Aid: Aid, Cid: cid, Uname: uname, Password: password, RegisterDate: register_date, Email: email, Auth: auth, AvtarUrl: avtar_url, Descp: descp}
	defer stmtOut.Close()
	return res, nil
}