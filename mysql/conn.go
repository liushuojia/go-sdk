package mysqlConn

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

const (
	max_id   = 10
	max_open = 100
)

func GormDB(username, password, host string, port int, database string) (*gorm.DB, error) {
	conn, err := gorm.Open(
		mysql.Open(
			fmt.Sprintf(
				"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				username, password, host, port, database,
			),
		),
		&gorm.Config{
			Logger: logger.New(
				log.New(), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
				logger.Config{
					SlowThreshold: time.Second, // 慢 SQL 阈值
					LogLevel:      logger.Info, // 日志级别
					//IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
					Colorful: false, // 禁用彩色打印
				},
			),
		},
	)
	if err != nil {
		return nil, err
	}

	db, err := conn.DB()
	if err != nil {
		return nil, err

	}
	db.SetMaxIdleConns(max_id)            // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.SetMaxOpenConns(max_open)          // SetMaxOpenConns sets the maximum number of open connections to the database.
	db.SetConnMaxLifetime(24 * time.Hour) // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db.SetConnMaxIdleTime(time.Hour)
	return conn, nil
}
func Close(db *gorm.DB) {
	if db, err := db.DB(); err == nil && db.Ping() == nil {
		db.Close()
	}
}
