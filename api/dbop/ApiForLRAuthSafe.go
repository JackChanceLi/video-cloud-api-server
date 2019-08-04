package dbop

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-api-server/api/defs"
	"log"
)

func InsertLRAuthSafeByCom(lid string, website string, wtype int) (*defs.LiveRoomAuthSafeIdentity,error) {
	stmtIns, err := dbConn.Prepare("INSERT INTO live_auth_safe (lid, website, wtype) VALUES (?, ?, ?)")
	if err != nil {
		log.Printf("%s",err)
		return nil, err
	}

	_,err = stmtIns.Exec(lid, website, wtype)
	if err != nil {
		return nil, err
	}

	log.Printf(" Insert success")

	defer stmtIns.Close()

	LRAS := &defs.LiveRoomAuthSafeIdentity{}
	LRAS.Lid = lid
	LRAS.Website = website
	LRAS.Wtype = wtype

	return LRAS, nil
}

func UpdateLRAuthSafe(lid string, website string, wtype int) (*defs.LiveRoomAuthSafeIdentity,error) {
	stmtUpa, err := dbConn.Prepare("UPDATE live_auth_safe SET wtype = ? WHERE lid = ? and website = ?")
	if err != nil {
		log.Printf("%s",err)
		return nil, err
	}

	_,err = stmtUpa.Exec(wtype, lid, website)
	if err != nil {
		log.Printf("%s",err)
		return nil, err
	}

	log.Printf(" Update success")

	defer stmtUpa.Close()

	LRAS := &defs.LiveRoomAuthSafeIdentity{}
	LRAS.Lid = lid
	LRAS.Website = website
	LRAS.Wtype = wtype

	return LRAS, nil
}

func RetrieveLRAuthSafeByLid(Lid string) (*defs.LiveRoomAuthSafeIdentity, error) {
	stmtOut, err := dbConn.Prepare("SELECT website, wtype FROM live_auth_safe WHERE lid = ?")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	var website string
	var wtype int
	stmtOut.QueryRow(Lid).Scan(&website, &wtype)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	LRAS := &defs.LiveRoomAuthSafeIdentity{Lid: Lid, Website: website, Wtype: wtype}
	defer stmtOut.Close()
	return LRAS, nil
}

func RetrieveLRAuthSafeWhiteList(Lid string) ([] defs.LiveRoomAuthSafeIdentity, error) {  //以切片的形式返回查询直播间权限安全白名单的结果
	var LRAS [] defs.LiveRoomAuthSafeIdentity
	stmtOut, err := dbConn.Prepare("SELECT * FROM live_auth_safe WHERE lid = ? and wtype = ?")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	//cid := Cid
	rows, err := stmtOut.Query(Lid, 0)
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	for rows.Next() {
		var lid, website string
		var wtype int
		if er := rows.Scan(&lid, &website, &wtype); er != nil {
			log.Printf("Retrieve live_auth_safe error: %s", er)
			return nil, err
		}

		Lras := defs.LiveRoomAuthSafeIdentity{Lid: Lid, Website: website, Wtype: 0}
		LRAS = append(LRAS, Lras)
	}
	return LRAS, nil
}

func RetrieveLRAuthSafeBlackList(Lid string) ([] defs.LiveRoomAuthSafeIdentity, error) {  //以切片的形式返回查询直播间权限安全黑名单的结果
	var LRAS [] defs.LiveRoomAuthSafeIdentity
	stmtOut, err := dbConn.Prepare("SELECT * FROM live_auth_safe WHERE lid = ? and wtype = ?")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	//cid := Cid
	rows, err := stmtOut.Query(Lid, 1)
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	for rows.Next() {
		var lid, website string
		var wtype int
		if er := rows.Scan(&lid, &website, &wtype); er != nil {
			log.Printf("Retrieve live_auth_safe error: %s", er)
			return nil, err
		}

		Lras := defs.LiveRoomAuthSafeIdentity{Lid: Lid, Website: website, Wtype: 1}
		LRAS = append(LRAS, Lras)
	}
	return LRAS, nil
}