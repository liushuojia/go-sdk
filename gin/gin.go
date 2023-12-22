package GIN

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func Body(c *gin.Context, arr ...interface{}) ([]byte, error) {
	body, err := c.GetRawData()
	if err != nil {
		return nil, err
	}

	if arr != nil && len(arr) > 0 {
		for _, v := range arr {
			if err := json.Unmarshal(body, v); err != nil {
				return nil, err
			}
		}
	}

	return body, err
}

func Int(c *gin.Context, field string) (int, error) {
	value := c.Query(field)
	if value == "" {
		return 0, errors.New("参数为空")
	}
	return strconv.Atoi(value)
}
func Int64(c *gin.Context, field string) (int64, error) {
	value := c.Query(field)
	if value == "" {
		return 0, errors.New("参数为空")
	}
	return strconv.ParseInt(value, 10, 64)
}
func String(c *gin.Context, field string) string {
	return strings.Trim(c.Query(field), " ")
}

func IntSlice(c *gin.Context, field string) []int {
	var arr []int
	if q := String(c, field); q != "" {
		for _, v := range strings.Split(q, ",") {
			if n, err := strconv.Atoi(v); err == nil {
				arr = append(arr, n)
			}
		}
	}
	if _arr := c.QueryArray(field + "[]"); len(_arr) > 0 {
		for _, v := range _arr {
			if n, err := strconv.Atoi(v); err == nil {
				arr = append(arr, n)
			}
		}
	}
	return arr
}
func Int64Slice(c *gin.Context, field string) []int64 {
	var arr []int64
	if q := String(c, field); q != "" {
		for _, v := range strings.Split(q, ",") {
			if n, err := strconv.ParseInt(v, 10, 64); err == nil {
				arr = append(arr, n)
			}
		}
	}
	if _arr := c.QueryArray(field + "[]"); len(_arr) > 0 {
		for _, v := range _arr {
			if n, err := strconv.ParseInt(v, 10, 64); err == nil {
				arr = append(arr, n)
			}
		}
	}
	return arr
}
func StringSlice(c *gin.Context, field string) []string {
	var arr []string
	if q := String(c, field); q != "" {
		for _, v := range strings.Split(q, ",") {
			if v != "" {
				arr = append(arr, v)
			}
		}
	}
	if _arr := c.QueryArray(field + "[]"); len(_arr) > 0 {
		for _, v := range _arr {
			if v != "" {
				arr = append(arr, v)
			}
		}
	}

	return arr
}

func ParamInt(c *gin.Context, field string) (int, error) {
	value := c.Param(field)
	if value == "" {
		return 0, errors.New("参数为空")
	}
	return strconv.Atoi(value)
}
func ParamInt64(c *gin.Context, field string) (int64, error) {
	value := c.Param(field)
	if value == "" {
		return 0, errors.New("参数为空")
	}
	return strconv.ParseInt(value, 10, 64)
}
func ParamString(c *gin.Context, field string) string {
	return strings.Trim(c.Param(field), " ")
}

func FormInt(c *gin.Context, field string) (int, error) {
	value := c.PostForm(field)
	if value == "" {
		return 0, errors.New("参数为空")
	}
	return strconv.Atoi(value)
}
func FormInt64(c *gin.Context, field string) (int64, error) {
	value := c.PostForm(field)
	if value == "" {
		return 0, errors.New("参数为空")
	}
	return strconv.ParseInt(value, 10, 64)
}
func FormString(c *gin.Context, field string) string {
	return strings.Trim(c.PostForm(field), " ")
}
