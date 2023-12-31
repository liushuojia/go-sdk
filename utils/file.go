package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/http"
	"os"
)

func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, nil
}
func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		return false
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func WriteFile(name string, content []byte) error {
	return os.WriteFile(name, content, 0644)
}
func ReadFile(name string) (data []byte, err error) {
	return os.ReadFile(name)
}

// CopyFile 文件复制
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		return
	}
	defer srcFile.Close()

	//打开dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return
	}
	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}

// FileMd5 获取文件哈希值
func FileMd5(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(md5hash.Sum(nil)), nil
}

// FileContentType 获取文件类型
func FileContentType(filePath string) (string, error) {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	if err != nil {
		return "", err
	}

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)
	if _, err := file.Read(buffer); err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
