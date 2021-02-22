package uploader

import (
	"fmt"
	"os"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var client *AliyunClient

//AliyunClient is the
type AliyunClient struct {
	bucket *oss.Bucket
}

//Connect connects
func (a *AliyunClient) Connect() error {

	if a.bucket != nil {
		return nil
	}

	var err error

	// oss.UseCname(true)为开启CNAME。CNAME是指将自定义域名绑定到存储空间上。
	client, err := oss.New(os.Getenv("ALIYUN_OSS_ENDPOINT"), os.Getenv("ALIYUN_OSS_ACCESSKEY_ID"), os.Getenv("ALIYUN_OSS_ACCESSKEY_SECRET"))

	bucketName := os.Getenv("ALIYUN_OSS_BUCKET_NAME")

	a.bucket, err = client.Bucket(bucketName)

	if err != nil {
		return err
	}

	return nil
}

//UploadFile 返回上传的url err remoteF
func (a *AliyunClient) UploadFile(localFile string, remoteFile string) (string, error) {

	if a.bucket == nil {
		a.Connect()
	}

	if a.FileExists(remoteFile) {
		return fmt.Sprintf("%s/%s", os.Getenv("ALIYUN_OSS_VISIT_URL"), remoteFile), nil
	}

	err := a.bucket.PutObjectFromFile(remoteFile, localFile)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return fmt.Sprintf("%s/%s", os.Getenv("ALIYUN_OSS_VISIT_URL"), remoteFile), nil
}

//UploadString 返回上传的url err remoteF
func (a *AliyunClient) UploadString(content string, remoteFile string) (string, error) {

	if a.bucket == nil {
		a.Connect()
	}

	if a.FileExists(remoteFile) {
		return fmt.Sprintf("%s/%s", os.Getenv("ALIYUN_OSS_VISIT_URL"), remoteFile), nil
	}

	err := a.bucket.PutObject(remoteFile, strings.NewReader(content))

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return fmt.Sprintf("%s/%s", os.Getenv("ALIYUN_OSS_VISIT_URL"), remoteFile), nil
}

//FileExists 文件是否存在
func (a *AliyunClient) FileExists(remoteFile string) bool {

	if a.bucket == nil {
		a.Connect()
	}

	res, _ := a.bucket.IsObjectExist(remoteFile)
	return res
}

//NewClient 返回个阿里云client
func NewClient() *AliyunClient {
	if client != nil {
		return client
	}

	client = &AliyunClient{}

	err := client.Connect()

	if err != nil {
		fmt.Println(err)
	}

	return client

}
