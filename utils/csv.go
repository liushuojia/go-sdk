package utils

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"strconv"
	"strings"
)

func WriteCsvString(c *gin.Context, b *bytes.Buffer) {

	if string, err := UTF82GBK(b.String()); err == nil {
		c.String(200, string)
	}
	/*
		if f, ok := c.Writer.(http.Flusher); ok {
			f.Flush()
		}
	*/
	b.Reset()
	return
}
func GetCsvMessage(dataMap map[string]interface{}, keyMap []string) (returnMap []string) {
	for _, key := range keyMap {
		if value, ok := dataMap[key]; ok {
			setString := ""
			switch str := value.(type) {
			case string:
				setString = str
			case int:
				setString = strconv.FormatInt(int64(str), 10)
			case int8:
				setString = strconv.FormatInt(int64(str), 10)
			case int16:
				setString = strconv.FormatInt(int64(str), 10)
			case int32:
				setString = strconv.FormatInt(int64(str), 10)
			case int64:
				setString = strconv.FormatInt(int64(str), 10)
			case float32:
				setString = strconv.FormatFloat(float64(str), 'f', -1, 32)
			case float64:
				setString = strconv.FormatFloat(str, 'f', -1, 64)
			case bool:
				if str {
					setString = "true"
				} else {
					setString = "false"
				}
			}
			returnMap = append(returnMap, setString)
		} else {
			returnMap = append(returnMap, "")
		}
	}
	return
}
func UTF82GBK(src string) (string, error) {
	reader := transform.NewReader(strings.NewReader(src), simplifiedchinese.GBK.NewEncoder())
	if buf, err := ioutil.ReadAll(reader); err != nil {
		return "", err
	} else {
		return string(buf), nil
	}
}
