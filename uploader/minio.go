package uploader

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient struct {
	client     *minio.Client
	bucketName string
	region     string
}

//Connect connects
func (a *MinioClient) Connect() (err error) {

	endpoint := os.Getenv("MINIO_ENDPOINT")

	a.client, err = minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(os.Getenv("ACCESS_KEY"), os.Getenv("ACCESS_SECRET"), ""),
	})

	a.bucketName = os.Getenv("BUCKET_NAME")
	a.region = os.Getenv("REGION")

	err = a.client.MakeBucket(context.Background(), a.bucketName, minio.MakeBucketOptions{
		Region: a.region,
	})

	return err
}

//UploadFile 返回上传的url err remoteF
func (a *MinioClient) UploadFile(filePath string, objectName string) (string, error) {

	if a.FileExists(filePath) {
		return fmt.Sprintf("%s/%s", os.Getenv("VISIT_URL"), objectName), nil
	}

	_, err := a.client.FPutObject(context.Background(), a.bucketName, objectName, filePath, minio.PutObjectOptions{})

	return fmt.Sprintf("%s/%s", os.Getenv("VISIT_URL"), objectName), err
}

//UploadString 返回上传的url err remoteF
func (a *MinioClient) UploadString(content string, remoteFile string) (string, error) {

	if a.FileExists(remoteFile) {
		return fmt.Sprintf("%s/%s", os.Getenv("VISIT_URL"), remoteFile), nil
	}

	err := a.client.MakeBucket(context.Background(), a.bucketName, minio.MakeBucketOptions{
		Region: a.region,
	})

	_, err = a.client.PutObject(context.Background(), a.bucketName, remoteFile, strings.NewReader(content), -1, minio.PutObjectOptions{})

	return fmt.Sprintf("%s/%s", os.Getenv("VISIT_URL"), remoteFile), err
}

//FileExists 文件是否存在
func (a *MinioClient) FileExists(remoteFile string) bool {

	fileInfo, _ := a.client.GetObject(context.Background(), a.bucketName, remoteFile, minio.GetObjectOptions{})

	stat, _ := fileInfo.Stat()

	if stat.Size == 0 {
		return false
	}

	return true
}

//NewQiniuClient en
func NewMinioClient() *MinioClient {
	return &MinioClient{}
}
