package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func DateToTime(datetime string) int64 {
	datetime = strings.Replace(datetime, "T", " ", -1)
	tmp1 := strings.Split(datetime, "+")
	datetime = tmp1[0]

	timeLayout := "2006-01-02 15:04:05"  //转化所需模板
	loc, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation(timeLayout, datetime, loc)

	return tmp.Unix() //转化为时间戳 类型是int64
}
func TimeToDate(timestamp int64, timeLayout string) string {
	//时间戳转化为日期
	if timestamp <= 0 {
		return ""
	}
	return time.Unix(timestamp, 0).Local().Format(timeLayout)
}

func TodayLastTime() int64 {
	tmpTime := time.Now()
	day := fmt.Sprintf("%02d-%02d-%02d 23:59:59", tmpTime.Year(), tmpTime.Month(), tmpTime.Day())
	return DateToTime(day)
}
func GetFirstDateOfYear(d time.Time) time.Time {
	return time.Date(d.Year(), 1, 1, 0, 0, 0, 0, d.Location())
}
func GetLastDateOfYear(d time.Time) time.Time {
	return GetFirstDateOfYear(d).AddDate(1, 0, 0).Add(-1 * time.Second)
}
func GetFirstDateOfMonth(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), 1, 0, 0, 0, 0, d.Location())
}
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, 0).Add(-1 * time.Second)
}
func GetFirstSecondOfDate(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}
func GetLastSecondOfDate(d time.Time) time.Time {
	return GetFirstSecondOfDate(d).AddDate(0, 0, 1).Add(-1 * time.Second)
}

func TimeFomatShow(datetime time.Time) string {
	if datetime.Format("2006-01-02") == "1970-01-01" {
		return "-"
	}
	havePassTime := time.Now().Unix() - datetime.Unix()
	if havePassTime <= 10 {
		return "刚刚"
	}
	if havePassTime <= 60 {
		return strconv.Itoa(int(havePassTime)) + "秒前"
	}

	if havePassTime < 60*60 {
		return strconv.Itoa(int(havePassTime/60)) + "分钟前"
	}

	if havePassTime < 24*60*60 {
		return strconv.Itoa(int(havePassTime/60/60)) + "小时前"
	}

	if havePassTime < 10*24*60*60 {
		return strconv.Itoa(int(havePassTime/60/60/24)) + "天前"
	}
	return datetime.Format("2006.01.02 15:04")
}
