package dbop

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-api-server/api/defs"
	"go-api-server/api/utils"
	"log"
	"strings"
)

func UploadResourseByCom(aid string, cid string, name string, rtype string, size float64, label []string, res_url string, pic_url string) (*defs.ResourseIdentity,error) {
	time, _ := getCurrentTime()
	stmtIns, err := dbConn.Prepare("INSERT INTO resourse (rid, aid, cid, name, rtype, size, label, time, res_url, pic_url) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	rid, _ := utils.NewUUID()
	var store string
	store = label[0]
	for i := 1; i < len(label); i++ {
		store += ";" + label[i]
	}
	_,err = stmtIns.Exec(rid, aid, cid, name, rtype, size, store, time, res_url, pic_url)
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
	res.ResUrl = res_url
	res.PicUrl = pic_url

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

func UpdateResourse(rid string, name string, label []string, pic_url string) (*defs.ResourseIdentity,error) {
	stmtUpa, err := dbConn.Prepare("UPDATE resourse SET name = ?, label = ?, pic_url = ? WHERE rid = ?")
	if err != nil {
		log.Printf("%s",err)
		return nil, err
	}

	var store string
	store = label[0]
	for i := 1; i < len(label); i++ {
		store += ";" + label[i]
	}

	_,err = stmtUpa.Exec(name, store, pic_url, rid)
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
	res.ResUrl = Res.ResUrl
	res.PicUrl = pic_url
	return res, nil
}

func RetrieveResourseByRid(Rid string) (*defs.ResourseIdentity, error) {
	stmtOut, err := dbConn.Prepare("SELECT aid, cid, name, rtype, size, label, time, res_url, pic_url FROM resourse WHERE rid = ?")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	var aid, cid, name, rtype, llabel, time, res_url, pic_url string
	var label []string
	var size float64
	stmtOut.QueryRow(Rid).Scan(&aid, &cid, &name, &rtype, &size, &llabel, &time, &res_url, &pic_url)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	label = strings.Split(llabel, ";")
	res := &defs.ResourseIdentity{Rid: Rid, Aid: aid, Cid: cid, Name: name, Rtype: rtype, Size: size, Label: label, Time: time, ResUrl: res_url, PicUrl: pic_url}
	defer stmtOut.Close()
	return res, nil
}

func RetrieveResourseByCid(Cid string) ([] defs.ResourseIdentity, error) {  //查找同公司文件并以切片的形式返回查询文件的结果
	var resourse [] defs.ResourseIdentity
	stmtOut, err := dbConn.Prepare("SELECT rid, aid, name, rtype, size, label, time, res_url, pic_url FROM resourse WHERE cid = ?")
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
		var rid, aid, cid, name, rtype, llabel, time, res_url, pic_url string
		var label []string
		var size float64
		if er := rows.Scan(&rid, &aid, &name, &rtype, &size, &llabel, &time, &res_url, &pic_url); er != nil {
			log.Printf("Retrieve resourse error: %s", er)
			return nil, err
		}
		label = strings.Split(llabel, ";")
		Res := defs.ResourseIdentity{Rid: rid, Aid: aid, Cid: cid, Name: name, Rtype: rtype, Size: size, Label: label, Time: time, ResUrl: res_url, PicUrl: pic_url}
		resourse = append(resourse, Res)
	}
	return resourse, nil
}
