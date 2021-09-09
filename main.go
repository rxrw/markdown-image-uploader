package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"iuv520/pic-upload/uploader"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
)

var (
	dir string
)

func main() {
	dir = getPath()

	log.Printf("开始执行... \n 目录：%s", dir)

	res := scanDirs(dir)

	modifyFile(res)
}

func scanDirs(dirName string) []string {
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Println(err)
	}
	var fileList []string
	for _, file := range files {
		if file.IsDir() {
			fileList = append(fileList, scanDirs(dirName+string(os.PathSeparator)+file.Name())...)
		} else {
			if path.Ext(file.Name()) == ".md" {
				fileList = append(fileList, dirName+string(os.PathSeparator)+file.Name())
			}
		}
	}
	return fileList
}

func modifyFile(files []string) {
	for _, file := range files {
		first, err := ioutil.ReadFile(file)
		content := string(first[:])
		if err != nil {
			fmt.Println(err)
		}
		// content
		newContent := findImage(content, file)

		if newContent != "" {

			ioutil.WriteFile(file, []byte(newContent), 0600)
		}
	}
}

func findImage(content string, fileName string) string {
	pattern := regexp.MustCompile(`!\[.*?\]\((.*?)\)|<img.*?src=['"](.*?)['"].*?>`)

	matched := pattern.FindAllStringSubmatch(content, -1)

	result := make(map[string]string)

	if len(matched) > 0 {
		for _, match := range matched {
			value, ok := result[match[1]]
			if ok && value != "" {
				//处理同文件里的重复图片
				continue
			}
			res := replaceImage(match[1], fileName)
			if match[1] != res && res != "" {
				result[match[1]] = res
			}
		}
	}

	if len(result) > 0 {
		log.Printf("%s中需要替换的图片：%v", fileName, result)
		for k, v := range result {
			content = strings.ReplaceAll(content, k, v)
		}
		return content
	}
	log.Printf("%s中没有需要替换的图片。", fileName)

	return ""
}

func replaceImage(originImage string, fileName string) string {
	path := getPath()

	var client uploader.Uploader = uploader.NewClient()

	remoteName := ""

	remoteName = strings.TrimLeft(strings.Replace(fileName, path, "", 1), string(os.PathSeparator))
	originPathArr := strings.Split(remoteName, string(os.PathSeparator))
	originPath := strings.Join(originPathArr[:len(originPathArr)-1], string(os.PathSeparator))

	remoteName = strings.Replace(remoteName, ".md", "", 1) + string(os.PathSeparator)

	var remote string

	//如果是一个本地图片，直接上传
	if !regexp.MustCompile(`((http(s?))|(ftp))://.*`).MatchString(originImage) {

		remoteName = remoteName + originImage

		originImage = path + originPath + string(os.PathSeparator) + originImage

		remote, _ = client.UploadFile(originImage, remoteName)

		if os.Getenv("DELETE_ORIGIN_URL") != "" {
			log.Printf("删除原有图片文件... <%s>", originImage)
			os.Remove(originImage)
		}
	} else {

		temp := strings.Split(originImage, "/")
		temp = strings.Split(temp[len(temp)-1], "?")
		urlName := temp[len(temp)-1]
		extArr := strings.Split(urlName, ".")
		ext := extArr[len(extArr)-1]
		if len(ext) > 4 {
			ext = "jpg"
		}
		h := md5.New()
		h.Write([]byte(urlName))
		cipherStr := h.Sum(nil)
		remoteName = remoteName + hex.EncodeToString(cipherStr) + "." + ext

		if strings.Contains(originImage, os.Getenv("VISIT_URL")) {
			return ""
		}

		// 这种备案导致的cdn用不了，解决一下先
		if os.Getenv("OLD_URL") != "" && strings.Contains(originImage, os.Getenv("OLD_URL")) {
			return strings.Replace(originImage, os.Getenv("OLD_URL"), os.Getenv("VISIT_URL"), 1)
		}

		fmt.Printf("GETTING FROM WEB: %s\n", originImage)
		resp, err := http.Get(originImage)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		if resp.StatusCode != 200 {
			fmt.Println(resp.StatusCode)
			fmt.Println(resp.Request.URL.String())
			return ""
		}
		body, _ := ioutil.ReadAll(resp.Body)

		remote, _ = client.UploadString(string(body), remoteName)
	}

	return remote
}

func getPath() string {
	path := os.Args[1]

	if path != "" {

		return fmt.Sprintf("%s%s", strings.TrimRight(path, string(os.PathSeparator)), string(os.PathSeparator))
	}

	return ""
}
