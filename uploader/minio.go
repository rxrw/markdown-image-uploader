package uploader

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient struct {
	client     *minio.Client
	bucketName string
}

//Connect connects
func (a *MinioClient) Connect() (err error) {

	endpoint := os.Getenv("MINIO_ENDPOINT")

	a.client, err = minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(os.Getenv("ACCESS_KEY"), os.Getenv("ACCESS_SECRET"), ""),
	})

	if err != nil {
		return err
	}

	a.bucketName = os.Getenv("BUCKET_NAME")

	return err
}

//UploadFile 返回上传的url err remoteF
func (a *MinioClient) UploadFile(filePath string, objectName string) (string, error) {

	if a.FileExists(filePath) {
		return fmt.Sprintf("%s/%s", os.Getenv("VISIT_URL"), objectName), nil
	}

	file, _ := os.Open(objectName)
	ct, _ := GetFileContentType(file)

	_, err := a.client.FPutObject(context.Background(), a.bucketName, objectName, filePath, minio.PutObjectOptions{
		ContentType: ct,
	})

	return fmt.Sprintf("%s/%s", os.Getenv("VISIT_URL"), objectName), err
}

//UploadString 返回上传的url err remoteF
func (a *MinioClient) UploadString(content string, remoteFile string) (string, error) {

	if a.FileExists(remoteFile) {
		return fmt.Sprintf("%s/%s", os.Getenv("VISIT_URL"), remoteFile), nil
	}

	ct := http.DetectContentType([]byte(content))

	_, err := a.client.PutObject(context.Background(), a.bucketName, remoteFile, strings.NewReader(content), -1, minio.PutObjectOptions{
		ContentType: ct,
	})

	return fmt.Sprintf("%s/%s", os.Getenv("VISIT_URL"), remoteFile), err
}

//FileExists 文件是否存在
func (a *MinioClient) FileExists(remoteFile string) bool {

	fileInfo, _ := a.client.GetObject(context.Background(), a.bucketName, remoteFile, minio.GetObjectOptions{})

	stat, _ := fileInfo.Stat()

	return stat.Size != 0
}

func GetFileContentType(out *os.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

//NewQiniuClient en
func NewMinioClient() *MinioClient {
	return &MinioClient{}
}
