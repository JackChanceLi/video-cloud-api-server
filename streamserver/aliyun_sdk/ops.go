package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"net/http"
	"os"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
)

func createBucket(endpoint, accessKeyId, accessKeySecret, bucketName string) {
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		handleError(err)
		return
	}
	// 创建存储空间。
	err = client.CreateBucket(bucketName)
	if err != nil {
		handleError(err)
		return
	}
	fmt.Println("Create bucket success")
}

func uploadFile(endpoint, accessKeyId, accessKeySecret, bucketName, objectName, fileName string) {
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}


	// 上传本地文件。
	err = bucket.PutObjectFromFile(objectName, fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	fmt.Println("Upload success")
}

func getAssumeRole(endpoint, accessKeyId, accessKeySecret, bucketName, objectName, fileName string) {

}
