package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/url"
	"os"
	"strings"
)

func FileToMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	_, _ = io.Copy(hash, file)
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// HashFileMd5 return md5 value of file
func HashFileMd5(filePath string) (md5Str string, fileStat os.FileInfo, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	fileStat, err = file.Stat()
	hash := md5.New()
	if _, err = io.Copy(hash, file); err != nil {
		return
	}
	hashInBytes := hash.Sum(nil)[:16]
	md5Str = hex.EncodeToString(hashInBytes)
	return
}

func GetFileType(fileName string) string {
	// 获取文件名后缀
	index := strings.LastIndex(fileName, ".")
	if index == -1 {
		// 文件名中没有后缀，无法确定文件类型
		return "unknown"
	}
	// 提取文件后缀并转换为小写
	fileSuffix := strings.ToLower(fileName[index+1:])
	// 根据文件后缀来判断文件类型
	switch fileSuffix {
	case "jpg", "jpeg", "png", "gif":
		return "image"
	case "mp4", "avi", "mov":
		return "video"
	case "mp3", "wav", "flac":
		return "audio"
	case "doc", "docx", "pdf", "txt", "ppt", "pptx":
		return "doc"
	case "zip", "rar", "7z":
		return "zip"
	default:
		return "other"
	}
}

// RemoveDomainName 删除前缀域名部分
func RemoveDomainName(urlStr string) string {
	// 解析URL
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return ""
	}
	// 获取域名部分
	domain := parsedURL.Hostname()
	// 判断协议类型
	protocol := "http"
	if parsedURL.Scheme == "https" {
		protocol = "https"
	}
	// 截取域名后面的内容
	return strings.TrimPrefix(urlStr, protocol+"://"+domain)
}
