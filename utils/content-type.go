package utils

import (
	"fmt"
	"mime"
	"path/filepath"
)

// ContentType2Extension 根据Content-Type返回文件后缀名
func ContentType2Extension(contentType string) []string {
	// 解析Content-Type
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return nil
	}
	// 根据mediaType获取对应的文件后缀
	extensions, err := mime.ExtensionsByType(mediaType)
	if err != nil || len(extensions) == 0 {
		return nil
	}
	fmt.Println(extensions)

	// 通常取第一个后缀名
	return extensions
}

func GetContentType(filePath string) (string, error) {
	ext := filepath.Ext(filePath)
	if ext == "" {
		return "", fmt.Errorf("unable to detect content type: no file extension")
	}
	contentType := mime.TypeByExtension(ext)
	return contentType, nil
}
