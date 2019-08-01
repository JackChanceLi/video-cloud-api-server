package dbop

import (
	"database/sql"
	"go-micro-server/api/defs"
	"log"
	"strconv"
	"sync"
)

func InsertSession(sid string, ttl int64, uid string) error {
	ttlStr := strconv.FormatInt(ttl,10)
	stmtIns,err := dbConn.Prepare("INSERT INTO sessions (session_id, TTL, uid) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	_,err = stmtIns.Exec(sid, ttlStr, uid)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return  nil
}

func RetrieveSession(sid string) (*defs.Session, error) {
	ss := &defs.Session{}
	stmtOut, err := dbConn.Prepare("SELECT TTL, uid FROM sessions WHERE session_id = ?")
	if err != nil {
		return nil,err
	}

	var ttl string
	var uname string
	stmtOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if res, err := strconv.ParseInt(ttl, 10,64); err == nil {
		ss.TTL = res
		ss.Uid = uname
	} else {
		return nil, err
	}

	defer stmtOut.Close()
	return ss,nil
}

func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("SELECT * FROM sessions")
	if err != nil {
		log.Printf("%s", err)
		return nil,err
	}
	rows, err := stmtOut.Query()
	if err != nil {
		log.Printf("%s", err)
		return nil,err
	}
	log.Println("All sessions in DB are:")
	for rows.Next() {
		var id string
		var ttlstr string
		var uid string
		if er := rows.Scan(&id, &ttlstr, &uid); er != nil {
			log.Printf("retrieve sessions error: %s", er)
			break
		}
		if ttl, err1 := strconv.ParseInt(ttlstr, 10, 64); err1 == nil {
			ss := &defs.Session{Uid: uid, TTL:ttl}
			m.Store(id,ss)
			log.Printf("Session id:%s, ttl:%d", id, ss.TTL)
		}

	}
	return m,nil
}

func DeleteSession(sid string) error {
	stmtOut, err := dbConn.Prepare("DELETE FROM sessions WHERE session_id = ?")
	if err != nil {
		log.Printf("%s",err)
		return err
	}

	if _, err := stmtOut.Query(sid); err != nil {
		return err
	}

	defer stmtOut.Close()
	return nil
}
func DeleteSessionByName(Uid string) error {
	stmtOut, err := dbConn.Prepare("DELETE FROM sessions WHERE uid = ?")
	if err != nil {
		log.Printf("%s",err)
		return err
	}

	if _, err := stmtOut.Query(Uid); err != nil {
		return err
	}

	defer stmtOut.Close()
	return nil
}
