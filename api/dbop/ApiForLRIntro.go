package dbop

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-api-server/api/defs"
	"log"
)

func InsertLRIntroByCom(lid string, qorder int, pre_pic string) (*defs.LiveRoomIntroIdentity,error) {
	stmtIns, err := dbConn.Prepare("INSERT INTO live_intro (lid, qorder, pre_pic) VALUES (?, ?, ?)")
	if err != nil {
		log.Printf("%s",err)
		return nil, err
	}

	_,err = stmtIns.Exec(lid, qorder, pre_pic)
	if err != nil {
		return nil, err
	}

	log.Printf(" Insert success")

	defer stmtIns.Close()
	LRIn := &defs.LiveRoomIntroIdentity{}
	LRIn.Lid = lid
	LRIn.Qorder = qorder
	LRIn.Prepic = pre_pic

	return LRIn, nil
}

func UpdateLRIntro(lid string, qorder int, pre_pic string) (*defs.LiveRoomIntroIdentity,error) {
	stmtUpa, err := dbConn.Prepare("UPDATE live_intro SET qorder = ?, pre_pic = ? WHERE lid = ?")
	if err != nil {
		log.Printf("%s",err)
		return nil, err
	}

	_,err = stmtUpa.Exec(qorder, pre_pic, lid)
	if err != nil {
		return nil, err
	}

	log.Printf(" Update success")

	defer stmtUpa.Close()

	LRIn := &defs.LiveRoomIntroIdentity{}
	LRIn.Lid = lid
	LRIn.Qorder = qorder
	LRIn.Prepic = pre_pic


	return LRIn, nil
}

func RetrieveLRIntroByLid(Lid string) (*defs.LiveRoomIntroIdentity, error) {
	stmtOut, err := dbConn.Prepare("SELECT qorder, pre_pic FROM live_intro WHERE lid = ?")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	var pre_pic string
	var qorder int
	stmtOut.QueryRow(Lid).Scan(&qorder, &pre_pic)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	LRIn := &defs.LiveRoomIntroIdentity{Lid: Lid, Qorder: qorder, Prepic: pre_pic}
	defer stmtOut.Close()
	return LRIn, nil
}
