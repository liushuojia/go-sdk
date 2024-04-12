package utils

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

func LogFormat() {
	//dir, _ := os.Getwd()
	//logFilePath := dir + "/logs"
	//os.MkdirAll(logFilePath, os.ModePerm)
	//
	//logFileName := "log"
	//
	//// 日志文件
	//fileName := path.Join(logFilePath, logFileName)
	//
	//// 设置日志级别
	//log.SetLevel(log.DebugLevel)
	//
	//// 设置 rotatelogs
	//logWriter, _ := rotatelogs.New(
	//	// 分割后的文件名称
	//	fileName+".%Y%m%d.log",
	//
	//	// 生成软链，指向最新日志文件
	//	rotatelogs.WithLinkName(fileName),
	//
	//	// 设置最大保存时间(60天)
	//	rotatelogs.WithMaxAge(60*24*time.Hour),
	//
	//	// 设置日志切割时间间隔(1天)
	//	rotatelogs.WithRotationTime(24*time.Hour),
	//)
	//defer logWriter.Close()
	//writers := []io.Writer{
	//	logWriter,
	//	os.Stdout,
	//}
	//
	//// 写入文件
	//log.SetOutput(io.MultiWriter(writers...))

	//格式
	// JSONFormatter TextFormatter
	/*
		log.SetFormatter(&log.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	*/
	log.SetFormatter(new(MyFormatter))

	/*
		writeMap := lfshook.WriterMap{
			logrus.InfoLevel:  logWriter,
			logrus.FatalLevel: logWriter,
			logrus.DebugLevel: logWriter,
			logrus.WarnLevel:  logWriter,
			logrus.ErrorLevel: logWriter,
			logrus.PanicLevel: logWriter,
		}
		lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
			// 新增 Hook
			logger.AddHook(lfHook)

	*/

}

type MyFormatter struct{}

func (s *MyFormatter) Format(entry *log.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006/01/02 15:04:05")

	msg := ""
	if str, ok := entry.Data["req_uri"]; ok && str != "" {
		msg = fmt.Sprintf(
			"%s [GIN] | %d | %14s | %15s | %4s %s\n",
			timestamp, entry.Data["status_code"], entry.Data["latency_time"], entry.Data["client_ip"], entry.Data["req_method"], entry.Data["req_uri"],
		)
	} else {
		//msg = fmt.Sprintf("%s [%4s] %s\n", timestamp, strings.ToUpper(entry.Level.String()), entry.Message)
		msg = fmt.Sprintf("%s %s\n", timestamp, entry.Message)
	}

	return []byte(msg), nil
}
