package utils

import (
	"bytes"
	"io"
	"net/http"
)

func httpDo(method string, url string, msg string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(
		method,
		url,
		bytes.NewBuffer([]byte(msg)),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

// HttpDoPost 方式
func HttpDoPost(url string, msg string) ([]byte, error) {
	return httpDo("POST", url, msg)
}

// HttpDoGet 方式
func HttpDoGet(url string) ([]byte, error) {
	return httpDo("GET", url, "")
}
