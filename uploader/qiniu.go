package uploader

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

//QiniuClient is the
type QiniuClient struct {
	zoneMap map[string]*storage.Zone

	config storage.Config

	upToken string

	client *qbox.Mac
}

//Connect connects
func (a *QiniuClient) Connect() error {

	a.zoneMap = map[string]*storage.Zone{"Hadong": &storage.ZoneHuadong, "Huabei": &storage.ZoneHuabei, "Huanan": &storage.ZoneHuanan, "Beimei": &storage.ZoneBeimei, "Xinjiapo": &storage.ZoneXinjiapo}

	a.config = storage.Config{}

	a.config.Zone = a.zoneMap[os.Getenv("ENDPOINT")]

	a.config.UseHTTPS = true

	a.config.UseCdnDomains = false

	putPolicy := storage.PutPolicy{
		Scope: os.Getenv("BUCKET_NAME"),
	}

	a.client = qbox.NewMac(os.Getenv("ACCESS_KEY"), os.Getenv("ACCESS_SECRET"))

	a.upToken = putPolicy.UploadToken(a.client)

	return nil
}

//UploadFile 返回上传的url err remoteF
func (a *QiniuClient) UploadFile(localFile string, remoteFile string) (string, error) {

	if a.upToken == "" {
		a.Connect()
	}

	if a.FileExists(remoteFile) {
		return fmt.Sprintf("%s/%s", os.Getenv("VISIT_URL"), remoteFile), nil
	}

	formUploader := storage.NewFormUploader(&a.config)

	ret := storage.PutRet{}

	putExtra := storage.PutExtra{}

	err := formUploader.PutFile(context.Background(), &ret, a.upToken, remoteFile, localFile, &putExtra)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return fmt.Sprintf("%s/%s", os.Getenv("VISIT_URL"), remoteFile), nil
}

//UploadString 返回上传的url err remoteF
func (a *QiniuClient) UploadString(content string, remoteFile string) (string, error) {

	if a.upToken == "" {
		a.Connect()
	}

	if a.FileExists(remoteFile) {
		return fmt.Sprintf("%s/%s", os.Getenv("VISIT_URL"), remoteFile), nil
	}
	formUploader := storage.NewFormUploader(&a.config)

	ret := storage.PutRet{}

	putExtra := storage.PutExtra{}

	err := formUploader.Put(context.Background(), &ret, a.upToken, remoteFile, bytes.NewReader([]byte(content)), int64(len(content)), &putExtra)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return fmt.Sprintf("%s/%s", os.Getenv("VISIT_URL"), remoteFile), nil
}

//FileExists 文件是否存在
func (a *QiniuClient) FileExists(remoteFile string) bool {

	if a.upToken == "" {
		a.Connect()
	}

	bucketManager := storage.NewBucketManager(a.client, &a.config)

	fileInfo, _ := bucketManager.Stat(os.Getenv("BUCKET_NAME"), remoteFile)

	return fileInfo.Fsize != 0
}

//NewQiniuClient en
func NewQiniuClient() *QiniuClient {
	return &QiniuClient{}
}
