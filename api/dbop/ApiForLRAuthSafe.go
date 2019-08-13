package dbop

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-api-server/api/defs"
	"log"
)


func UpdateLRAuthSafe(lid string, WhiteSiteList string, BlackSiteList string) (*defs.LiveRoomAuthSafeIdentity, error) {
	stmtUpa, err := dbConn.Prepare("UPDATE live_auth_safe SET white_site_list = ?, black_site_list = ? WHERE lid = ? ")
	if err != nil {
		log.Printf("%s",err)
		return nil, err
	}

	_,err = stmtUpa.Exec(WhiteSiteList, BlackSiteList, lid)
	if err != nil {
		log.Printf("%s",err)
		return nil, err
	}
	defer stmtUpa.Close()

	LRAS := &defs.LiveRoomAuthSafeIdentity{}
	LRAS.Lid = lid
	LRAS.WhiteSiteList = WhiteSiteList
	LRAS.BlackSiteList = BlackSiteList

	return LRAS, nil
}

func RetrieveLRAuthSafeByLid(Lid string) (*defs.LiveRoomAuthSafeIdentity, error) {
	stmtOut, err := dbConn.Prepare("SELECT white_site_list, black_site_list FROM live_auth_safe WHERE lid = ?")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	var white_site_list, black_site_list string
	stmtOut.QueryRow(Lid).Scan(&white_site_list, &black_site_list)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	LRAS := &defs.LiveRoomAuthSafeIdentity{Lid: Lid, WhiteSiteList: white_site_list, BlackSiteList: black_site_list}
	defer stmtOut.Close()
	return LRAS, nil
}

func RetrieveLRAuthSafeList(Lid string) (* defs.LiveRoomAuthSafeIdentity, error) {  //以切片的形式返回查询直播间权限安全白名单的结果
	stmtOut, err := dbConn.Prepare("SELECT white_site_list, black_site_list FROM live_auth_safe WHERE lid = ? ")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	//cid := Cid
	var white_site_list, black_site_list string
	stmtOut.QueryRow(Lid).Scan(&white_site_list, &black_site_list)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	LRAS := &defs.LiveRoomAuthSafeIdentity{Lid: Lid, WhiteSiteList: white_site_list, BlackSiteList: black_site_list}
	return LRAS, nil
}

func InsertLRAuthSafeList(Lid, WhiteSiteList, BlackSiteList string) (*defs.LiveRoomAuthSafeIdentity, error) {
	stmtIns, err := dbConn.Prepare("INSERT INTO live_auth_safe (lid, white_site_list, black_site_list) VALUES (?, ?, ?)")
	if err != nil {
		log.Printf("pareparation:%v", err)
		return nil, err
	}

	_, err = stmtIns.Exec(Lid, WhiteSiteList, BlackSiteList)
	if err != nil {
		return nil, err
	}
	defer stmtIns.Close()

	authSafe := &defs.LiveRoomAuthSafeIdentity{}
	authSafe.Lid = Lid
	authSafe.WhiteSiteList = WhiteSiteList
	authSafe.BlackSiteList = BlackSiteList

	return  authSafe, nil
}
