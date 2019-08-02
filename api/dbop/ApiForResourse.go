package dbop

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-api-server/api/defs"
	"go-api-server/api/utils"
	"log"
)

func UploadResourseByCom(aid string, cid string, name string, rtype string, size float64, label string) (*defs.ResourseIdentity,error) {
	time, _ := getCurrentTime()
	stmtIns, err := dbConn.Prepare("INSERT INTO resourse (rid, aid, cid, name, rtype, size, label, time) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	rid, _ := utils.NewUUID()
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
	res.Rtype = Res.Rtype
	res.Size = Res.Size
	res.Label = label
	res.Time = Res.Time
	return res, nil
}

func RetrieveResourseByRid(Rid string) (*defs.ResourseIdentity, error) {
	stmtOut, err := dbConn.Prepare("SELECT aid, cid, name, rtype, size, label, time FROM resourse WHERE rid = ?")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	var aid, cid, name, rtype, label, time string
	var size float64
	stmtOut.QueryRow(Rid).Scan(&aid, &cid, &name, &rtype, &size, &label, &time)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	res := &defs.ResourseIdentity{Rid: Rid, Aid: aid, Cid: cid, Name: name, Rtype: rtype, Size: size, Label: label, Time: time}
	defer stmtOut.Close()
	return res, nil
}

func RetrieveResourseByCid(Cid string) ([] defs.ResourseIdentity, error) {  //查找同公司文件并以切片的形式返回查询文件的结果
	var resourse [] defs.ResourseIdentity
	stmtOut, err := dbConn.Prepare("SELECT rid, aid, name, rtype, size, label, time FROM resourse WHERE cid = ?")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	//cid := Cid
	rows, err := stmtOut.Query(Cid)
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	for rows.Next() {
		var rid, aid, cid, name, rtype, label, time string
		var size float64
		if er := rows.Scan(&rid, &aid, &name, &rtype, &size, &label, &time); er != nil {
			log.Printf("Retrieve resourse error: %s", er)
			return nil, err
		}

		Res := defs.ResourseIdentity{Rid: rid, Aid: aid, Cid: cid, Name: name, Rtype: rtype, Size: size, Label: label, Time: time}
		resourse = append(resourse, Res)
	}
	return resourse, nil
}