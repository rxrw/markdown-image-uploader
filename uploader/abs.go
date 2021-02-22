package uploader

//Uploader 接口类
type Uploader interface {
	Connect() error
	UploadFile(localFile string, remoteFile string) (string, error)
	UploadString(content string, remoteFile string) (string, error)
	FileExists(remoteFile string) bool
}
