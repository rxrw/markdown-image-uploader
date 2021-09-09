package uploader

import (
	"fmt"
	"os"
)

var client Uploader

//Uploader 接口类
type Uploader interface {
	Connect() error
	UploadFile(localFile string, remoteFile string) (string, error)
	UploadString(content string, remoteFile string) (string, error)
	FileExists(remoteFile string) bool
}

//NewClient 返回个uploader client
func NewClient() Uploader {
	if client != nil {
		return client
	}

	clientType := os.Getenv("CLIENT_NAME")
	fmt.Printf("使用%s作为上传客户端", clientType)
	switch clientType {
	case "aliyun":
		client = NewAliyunClient()
	case "qiniu":
		client = NewQiniuClient()
	default:
		client = NewAliyunClient()
	}

	err := client.Connect()

	if err != nil {
		fmt.Println(err)
	}

	return client

}
