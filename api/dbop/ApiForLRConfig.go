package dbop

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-api-server/api/defs"
	"log"
)

func InsertLRConfigByCom(lid string, live_pic string, danmu int, chat int, share int, share_text string, advertisement int, ad_jump_url string, ad_pic_url string, ad_text string) (*defs.LiveRoomConfigIdentity,error) {
	stmtIns, err := dbConn.Prepare("INSERT INTO live_config (lid, live_pic, danmu, chat, share, share_text, advertisement, ad_jump_url, ad_pic_url, ad_text) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Printf("%s",err)
		return nil, err
	}

	_,err = stmtIns.Exec(lid, live_pic, danmu, chat, share, share_text, advertisement, ad_jump_url, ad_pic_url, ad_text)
	if err != nil {
		return nil, err
	}

	log.Printf(" Insert success")

	defer stmtIns.Close()

	LRCon := &defs.LiveRoomConfigIdentity{}
	LRCon.Lid = lid
	LRCon.LivePic = live_pic
	LRCon.Danmu = danmu
	LRCon.Chat = chat
	LRCon.Share = share
	LRCon.ShareText = share_text
	LRCon.Advertisement = advertisement
	LRCon.AdJumpUrl = ad_jump_url
	LRCon.AdPicUrl = ad_pic_url
	LRCon.AdText = ad_text

	return LRCon, nil
}

func UpdateLRConfig(lid string, live_pic string, danmu int, chat int, share int, share_text string, advertisement int, ad_jump_url string, ad_pic_url string, ad_text string) (*defs.LiveRoomConfigIdentity,error) {
	stmtUpa, err := dbConn.Prepare("UPDATE live_config SET live_pic = ?, danmu = ?, chat = ?, share = ?, share_text = ?, advertisement = ?, ad_jump_url = ?, ad_pic_url = ?, ad_text = ? WHERE lid = ?")
	if err != nil {
		log.Printf("%s",err)
		return nil, err
	}

	_,err = stmtUpa.Exec(live_pic, danmu, chat, share, share_text, advertisement, ad_jump_url, ad_pic_url, ad_text, lid)
	if err != nil {
		return nil, err
	}

	log.Printf(" Update success")

	defer stmtUpa.Close()

	LRCon := &defs.LiveRoomConfigIdentity{}
	LRCon.Lid = lid
	LRCon.LivePic = live_pic
	LRCon.Danmu = danmu
	LRCon.Chat = chat
	LRCon.Share = share
	LRCon.ShareText = share_text
	LRCon.Advertisement = advertisement
	LRCon.AdJumpUrl = ad_jump_url
	LRCon.AdPicUrl = ad_pic_url
	LRCon.AdText = ad_text

	return LRCon, nil
}

func RetrieveLRConfigByLid(Lid string) (*defs.LiveRoomConfigIdentity, error) {
	stmtOut, err := dbConn.Prepare("SELECT live_pic, danmu, chat, share, share_text, advertisement, ad_jump_url, ad_pic_url, ad_text FROM live_config WHERE lid = ?")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	var live_pic, share_text, ad_jump_url, ad_pic_url, ad_text string
	var danmu, chat, share, advertisement int
	stmtOut.QueryRow(Lid).Scan(&live_pic, &danmu, &chat, &share, &share_text, &advertisement, &ad_jump_url, &ad_pic_url, &ad_text)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	LRCon := &defs.LiveRoomConfigIdentity{Lid: Lid, LivePic: live_pic, Danmu: danmu, Chat: chat, Share: share, ShareText: share_text, Advertisement: advertisement, AdJumpUrl: ad_jump_url, AdPicUrl: ad_pic_url, AdText: ad_text}
	defer stmtOut.Close()
	return LRCon, nil
}
