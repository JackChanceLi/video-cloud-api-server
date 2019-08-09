package dbop

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-api-server/api/defs"
	"log"
)

func InsertLRQualityByCom(lid string, delay int, transcode int, transcode_type []int, record int, record_type int) (*defs.LiveRoomQualityIdentity,error) {
	stmtIns, err := dbConn.Prepare("INSERT INTO live_quality (lid, delay, transcode, transcode_type, record, record_type) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Printf("%s",err)
		return nil, err
	}

	var tmp int
	for i := 0; i < len(transcode_type); i++ {
		tmp = tmp * 10 + transcode_type[i]
	}

	_,err = stmtIns.Exec(lid, delay, transcode, tmp, record, record_type)
	if err != nil {
		return nil, err
	}

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

func UpdateLRQuality(lid string, delay int, transcode int, transcode_type []int, record int, record_type int) (*defs.LiveRoomQualityIdentity,error) {
	stmtUpa, err := dbConn.Prepare("UPDATE live_quality SET delay = ?, transcode = ?, transcode_type = ?, record = ?, record_type = ? WHERE lid = ?")
	if err != nil {
		log.Printf("%s",err)
		return nil, err
	}
	var tmp int
	for i := 0; i < len(transcode_type); i++ {
		tmp = tmp * 10 + transcode_type[i]
	}

	_,err = stmtUpa.Exec(delay, transcode, tmp, record, record_type, lid)
	if err != nil {
		return nil, err
	}

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

	var delay, transcode, transcode_type, tmp, record, record_type,len int

	stmtOut.QueryRow(Lid).Scan(&delay, &transcode, &transcode_type, &record, &record_type)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	tmp = transcode_type
	//求长度
	for ;; len++ {
		if tmp == 0 {
			break
		}
		tmp /= 10
	}
	tType := make([]int,len)
	for i:=0;;i++ {
		if transcode_type == 0{
			break
		}
		tType[i] = transcode_type % 10
		transcode_type /= 10
	}
	fmt.Println(tType)

	LRQua := &defs.LiveRoomQualityIdentity{Lid: Lid, Delay: delay, Transcode: transcode, TranscodeType: tType, Record: record, RecordType: record_type}
	defer stmtOut.Close()
	return LRQua, nil
}
