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
	switch clientType {
	case "aliyun":
		client = NewAliyunClient()
	case "qiniu":
		client = NewQiniuClient()
	}

	err := client.Connect()

	if err != nil {
		fmt.Println(err)
	}

	return client

}
