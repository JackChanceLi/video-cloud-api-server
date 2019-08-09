package url

import (
	"crypto/md5"
	"fmt"
	"go-api-server/api/utils"
)
func NewRtmpUrl(streamId string) (string, error){
	key := utils.NewID()
	//将key和stream_id进行MD5加密
	str := key + streamId
	data := []byte(str)
	has := md5.Sum(data)
	authKey := fmt.Sprintf("%x", has) //将[]byte转成16进制
	fmt.Println("authKey:" +  authKey)
	//vhost := "default"
	url := "rtmp://" + IP + PORT + LiveName + streamId + "/?auth_key=" + authKey// + "&vhost=" + vhost
	return url, nil
}

func NewFlvUrl(streamId string) (string, error) {
	url := "http://" + Host + ":8090/" + "live/" + streamId + ".flv"
	return url, nil
}

func NewHlsUrl(streamId string) (string, error) {
	url := "http://" + Host + ":8090/" + LiveName +streamId + ".m3u8"
	return url, nil
}

func NewDisplayUrl(channelID string) (string, error) {
	url := "http://" + DisplayHost + ":8082/live/player/?channel_id=" + channelID
	return url, nil
}
