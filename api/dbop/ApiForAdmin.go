package dbop

import (
	_ "github.com/go-sql-driver/mysql"
	"go-api-server/api/defs"
	"log"
)

func InsertAdmin(aid string, cid string, uname string, password string, register_date string, email string, auth string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO admin (aid, cid, uname, password, register_date, email, auth) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_,err = stmtIns.Exec(aid, cid, uname, password, register_date, email, auth)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return  nil
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

func UpdateAdmin(aid string, uname string, password string, email string, auth string) error {
	stmtUpa, err := dbConn.Prepare("UPDATE admin SET uname = ?, password = ?, email = ?, auth = ? WHERE aid = ?")
	if err != nil {
		log.Printf("%s",err)
		return err
	}

	_,err = stmtUpa.Exec(uname, password, email, auth, aid)
	if err != nil {
		return err
	}

	defer stmtUpa.Close()
	return nil
}

func RetrieveAdminByAid(Aid string) (*defs.UserInformation, error) {
	stmtOut, err := dbConn.Prepare("SELECT * FROM admin WHERE aid = ?")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	row, err := stmtOut.Query(Aid)
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	var aid, cid, uname, password, registerDate, email, auth string
	if er := row.Scan(&aid, &cid, &uname, &password, &registerDate, &email, &auth); er != nil {
		log.Printf("Retrieve live_room error: %s", er)
	}

	Ad := &defs.UserInformation{Aid: aid, Cid: cid, Name: uname, Email:email, Auth: auth}
	return Ad, nil
}