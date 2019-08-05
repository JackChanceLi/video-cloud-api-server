package dbop

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-api-server/api/defs"
	"log"
)

func InsertLRSafeByCom(lid string, logo int, logo_url string, logo_position int, logo_transparency int, lamp int, lamp_type int, lamp_text string, lamp_font_size int, lamp_transparency int) (*defs.LiveRoomSafeIdentity,error) {
	stmtIns, err := dbConn.Prepare("INSERT INTO live_safe (lid, logo, logo_url, logo_position, logo_transparency, lamp, lamp_type, lamp_text, lamp_font_size, lamp_transparency) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Printf("%s",err)
		return nil, err
	}

	_,err = stmtIns.Exec(lid, logo, logo_url, logo_position, logo_transparency, lamp, lamp_type, lamp_text, lamp_font_size, lamp_transparency)
	if err != nil {
		return nil, err
	}

	log.Printf(" Insert success")

	defer stmtIns.Close()
	LRSf := &defs.LiveRoomSafeIdentity{}
	LRSf.Lid = lid
	LRSf.Logo = logo
	LRSf.LogoUrl = logo_url
	LRSf.LogoPosition = logo_position
	LRSf.LogoTransparency = logo_transparency
	LRSf.Lamp = lamp
	LRSf.LampType = lamp_type
	LRSf.LampText = lamp_text
	LRSf.LampFontSize = lamp_font_size
	LRSf.LampTransparency = lamp_transparency

	return LRSf, nil
}

func UpdateLRSafe(lid string, logo int, logo_url string, logo_position int, logo_transparency int, lamp int, lamp_type int, lamp_text string, lamp_font_size int, lamp_transparency int) (*defs.LiveRoomSafeIdentity,error) {
	stmtUpa, err := dbConn.Prepare("UPDATE live_safe SET logo = ?, logo_url = ?, logo_position = ?, logo_transparency = ?, lamp = ?, lamp_type = ?, lamp_text = ?, lamp_font_size = ?, lamp_transparency = ? WHERE lid = ?")
	if err != nil {
		log.Printf("%s",err)
		return nil, err
	}

	_,err = stmtUpa.Exec(logo, logo_url, logo_position, logo_transparency, lamp, lamp_type, lamp_text, lamp_font_size, lamp_transparency, lid)
	if err != nil {
		return nil, err
	}

	log.Printf(" Update success")

	defer stmtUpa.Close()
	LRSf := &defs.LiveRoomSafeIdentity{}
	LRSf.Lid = lid
	LRSf.Logo = logo
	LRSf.LogoUrl = logo_url
	LRSf.LogoPosition = logo_position
	LRSf.LogoTransparency = logo_transparency
	LRSf.Lamp = lamp
	LRSf.LampType = lamp_type
	LRSf.LampText = lamp_text
	LRSf.LampFontSize = lamp_font_size
	LRSf.LampTransparency = lamp_transparency

	return LRSf, nil
}

func RetrieveLRSafeByLid(Lid string) (*defs.LiveRoomSafeIdentity, error) {
	stmtOut, err := dbConn.Prepare("SELECT logo, logo_url, logo_position, logo_transparency, lamp, lamp_type, lamp_text, lamp_font_size, lamp_transparency FROM live_safe WHERE lid = ?")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	var logo_url, lamp_text string
	var logo, logo_position, logo_transparency, lamp, lamp_type, lamp_font_size, lamp_transparency int
	stmtOut.QueryRow(Lid).Scan(&logo, &logo_url, &logo_position, &logo_transparency, &lamp, &lamp_type, &lamp_text, &lamp_font_size, &lamp_transparency)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	LRSf := &defs.LiveRoomSafeIdentity{Lid: Lid, Logo: logo, LogoUrl: logo_url, LogoPosition: logo_position, LogoTransparency: logo_transparency, Lamp: lamp, LampType: lamp_type, LampText: lamp_text, LampFontSize: lamp_font_size, LampTransparency: lamp_transparency}
	defer stmtOut.Close()
	return LRSf, nil
}
