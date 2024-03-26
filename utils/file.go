package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
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

func DeleteFile(path string) error {
	return os.Remove(path)
}

func DeleteFolder(path string) error {
	files, err := os.ReadDir(path) //读取目录下所有文件和子文件夹
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := fmt.Sprintf("%s/%s", path, file.Name())
		if file.IsDir() {
			if err := DeleteFolder(filePath); err != nil {
				return err
			}
		} else {
			if err := DeleteFile(filePath); err != nil {
				return err
			}
		}
	}

	// 最后删除空文件夹
	return os.RemoveAll(path)
}

func CutImage(oldFileName, newFileName string, width int) (err error) {

	file, err := os.Open(oldFileName)
	if err != nil {
		return
	}
	defer file.Close()

	c, _, err := image.DecodeConfig(file)
	if err != nil {
		return
	}

	if c.Width <= (width) {
		_, err = CopyFile(newFileName, oldFileName)
		return
	}

	//按照宽度进行等比例缩放
	src, err := imaging.Open(oldFileName, imaging.AutoOrientation(true))
	if err != nil {
		return
	}

	src = imaging.Resize(src, width, 0, imaging.Lanczos)
	err = imaging.Save(src, newFileName)
	return
}
