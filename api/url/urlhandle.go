package url

import (
	"crypto/md5"
	"fmt"
	"go-api-server/api/utils"
)
func NewRtmpUrl() (string, error){
	streamId, _ := utils.NewStreamID()
	key := utils.NewID()
	//将key和stream_id进行MD5加密
	str := key + streamId
	data := []byte(str)
	has := md5.Sum(data)
	authKey := fmt.Sprintf("%x", has) //将[]byte转成16进制
	fmt.Println("authKey:" +  authKey)
	vhost := "default"
	url := "rtmp://" + IP + PORT + LiveName + streamId + "/?auth_key=" + authKey + "&vhost=" + vhost
	return url, nil
}

func NewFlvUrl() (string, error) {
	streamId, _ := utils.NewStreamID()
	url := "http://" + Host + ":8090/" + "live/" + streamId + ".flv"
	return url, nil
}

func NewHlsUrl() (string, error) {
	streamId, _ := utils.NewStreamID()
	url := "http://" + Host + ":8090/" + LiveName +streamId + ".m3u8"
	return url, nil
}

func NewDisplayUrl() (string, error) {
	channelId, _ := utils.NewStreamID()
	url := "http://" + Host + ":80/live/player/?channel_id=" + channelId
	return url, nil
}
