package dbop

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-api-server/api/defs"
	"log"
)

func InsertLRQualityByCom(lid string, delay int, transcode int, transcode_type int, record int, record_type int) (*defs.LiveRoomQualityIdentity,error) {
	stmtIns, err := dbConn.Prepare("INSERT INTO live_quality (lid, delay, transcode, transcode_type, record, record_type) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Printf("%s",err)
		return nil, err
	}

	_,err = stmtIns.Exec(lid, delay, transcode, transcode_type, record, record_type)
	if err != nil {
		return nil, err
	}

	log.Printf(" Insert success")

	defer stmtIns.Close()
	LRQua := &defs.LiveRoomQualityIdentity{}
	LRQua.Lid = lid
	LRQua.Delay = delay
	LRQua.Transcode = transcode
	LRQua.TranscodeType = transcode_type
	LRQua.Record = record
	LRQua.RecordType = record_type

	return LRQua, nil
}

func UpdateLRQuality(lid string, delay int, transcode int, transcode_type int, record int, record_type int) (*defs.LiveRoomQualityIdentity,error) {
	stmtUpa, err := dbConn.Prepare("UPDATE live_quality SET delay = ?, transcode = ?, transcode_type = ?, record = ?, record_type = ? WHERE lid = ?")
	if err != nil {
		log.Printf("%s",err)
		return nil, err
	}

	_,err = stmtUpa.Exec(delay, transcode, transcode_type, record, record_type, lid)
	if err != nil {
		return nil, err
	}

	log.Printf(" Update success")

	defer stmtUpa.Close()

	LRQua := &defs.LiveRoomQualityIdentity{}
	LRQua.Lid = lid
	LRQua.Delay = delay
	LRQua.Transcode = transcode
	LRQua.TranscodeType = transcode_type
	LRQua.Record = record
	LRQua.RecordType = record_type

	return LRQua, nil
}

func RetrieveLRQualityByLid(Lid string) (*defs.LiveRoomQualityIdentity, error) {
	stmtOut, err := dbConn.Prepare("SELECT delay, transcode, transcode_type, record, record_type FROM live_quality WHERE lid = ?")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	var delay, transcode, transcode_type, record, record_type int
	stmtOut.QueryRow(Lid).Scan(&delay, &transcode, &transcode_type, &record, &record_type)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	LRQua := &defs.LiveRoomQualityIdentity{Lid: Lid, Delay: delay, Transcode: transcode, TranscodeType: transcode_type, Record: record, RecordType: record_type}
	defer stmtOut.Close()
	return LRQua, nil
}
