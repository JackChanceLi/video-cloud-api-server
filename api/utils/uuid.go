package utils

import (
	"crypto/rand"
	"fmt"
	"github.com/rs/xid"
	"github.com/segmentio/ksuid"
	"io"
)

func NewUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits
	uuid[8] = uuid [8] &^ 0xc0 | 0x80
	// version 4 (pseudo-random)
	uuid[6] = uuid[6] &^ 0xf0 | 0x40

	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:12], uuid[12:16]), nil
}

func NewIDByUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits
	uuid[8] = uuid [8] &^ 0xc0 | 0x80
	// version 4 (pseudo-random)
	uuid[6] = uuid[6] &^ 0xf0 | 0x40

	return fmt.Sprintf("%x", uuid[0:5]), nil
}

func NewStreamID() (string, error) {
	//uuid := make([]byte, 16)
	//n, err := io.ReadFull(rand.Reader, uuid)
	//if n != len(uuid) || err != nil {
	//	return "", err
	//}
	//// variant bits
	//uuid[8] = uuid [8] &^ 0xc0 | 0x80
	//// version 4 (pseudo-random)
	//uuid[6] = uuid[6] &^ 0xf0 | 0x40
	//
	//return fmt.Sprintf("%x", uuid[0:5]), nil

	//id := xid.New().String()
	//nid := string(id[12:20])
	//return nid, nil

	id := ksuid.New().String()
	nid := string(id[17:27])
	return nid, nil
}

func NewID() string {  //另一种根据时间生成随机ID的函数
	id := xid.New()
	//fmt.Println(id)
	return id.String()
}
