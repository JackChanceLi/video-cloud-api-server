package dbop

import (
	_ "github.com/go-sql-driver/mysql"
	"go-api-server/api/defs"
	"log"
)

func UploadResourseByCom(rid string, aid string, cid string, name string, rtype string, size float64, label string, time string) (*defs.ResourseIdentity,error) {
	stmtIns, err := dbConn.Prepare("INSERT INTO resourse (rid, aid, cid, name, rtype, size, label, time) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	_,err = stmtIns.Exec(rid, aid, cid, name, rtype, size, label, time)
	if err != nil {
		return nil, err
	}

	log.Printf(" Insert success")

	defer stmtIns.Close()
	res := &defs.ResourseIdentity{}
	res.Rid = rid
	res.Aid = aid
	res.Cid = cid
	res.Name = name
	res.Rtype = rtype
	res.Size = size
	res.Label = label
	res.Time = time

	return res, nil
}

func DeleteResourse(rid string) error {
	stmtOut, err := dbConn.Prepare("DELETE FROM resourse WHERE rid = ?")
	if err != nil {
		log.Printf("%s",err)
		return err
	}

	if _, err := stmtOut.Query(rid); err != nil {
		return err
	}

	log.Printf(" Delete success")

	defer stmtOut.Close()
	return nil
}

func UpdateResourse(rid string, name string, label string) (*defs.ResourseIdentity,error) {
	stmtUpa, err := dbConn.Prepare("UPDATE resourse SET name = ?, label = ? WHERE rid = ?")
	if err != nil {
		log.Printf("%s",err)
		return nil, err
	}

	_,err = stmtUpa.Exec(name, label, rid)
	if err != nil {
		return nil, err
	}

	log.Printf(" Update success")

	defer stmtUpa.Close()
	Res, err := RetrieveResourseByRid(rid)
	if err != nil {
		return nil, err
	}
	res := &defs.ResourseIdentity{}
	res.Rid = rid
	res.Aid = Res.Aid
	res.Cid = Res.Cid
	res.Name = name
	res.Rtype = Res.
	res.Size = size
	res.Label = label
	res.Time = time
	return res, nil
}

func RetrieveResourseByRid(Rid string) (*defs.UserInformation, error) {
	stmtOut, err := dbConn.Prepare("SELECT * FROM resourse WHERE rid = ?")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	row, err := stmtOut.Query(Rid)
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	var rid, aid, cid, name, rtype, label, time string
	var size float64
	if er := row.Scan(&rid, &aid, &cid, &name, &rtype, &size, &label, &time); er != nil {
		log.Printf("Retrieve resourse error: %s", er)
	}

	Ad := &defs.UserInformation{Aid: aid, Cid: cid, Name: uname, Email:email, Auth: auth}
	return Ad, nil
}