package main

import (
	"fmt"
	"os"
)
func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}
func main() {
	// Endpoint以杭州为例，其它Region请按实际情况填写。
	endpoint := "http://oss-cn-beijing.aliyuncs.com"
	// 阿里云主账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM账号进行API访问或日常运维，请登录 https://ram.console.aliyun.com 创建RAM账号。
	accessKeyId := "LTAItzz13YeySswG"
	accessKeySecret := "yojnuT3XyXhvx4jOVneZT5i919W3jk"
	bucketName := "video-bupt"


	// 创建OSSClient实例。
	//createBucket(endpoint, accessKeyId, accessKeySecret, bucketName)

	//上传文件测试
	objectName := "123.mp4"
	fileName := "test.mp4"
	uploadFile(endpoint, accessKeyId, accessKeySecret, bucketName, objectName, fileName)
}
