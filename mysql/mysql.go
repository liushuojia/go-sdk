package mysqlConn

import (
	"context"
	"fmt"
	"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	sql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net"
	"time"
)

const (
	max_id   = 10
	max_open = 100
)

type MySQL struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func (m *MySQL) ConnWithSSH(client *ssh.Client) (*gorm.DB, error) {

	// 注册ssh代理
	mysql.RegisterDialContext("mysql+ssh", func(_ context.Context, addr string) (net.Conn, error) {
		return client.Dial("tcp", addr)
	})

	conn, err := gorm.Open(
		sql.Open(
			fmt.Sprintf(
				"%s:%s@mysql+ssh(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
				m.User, m.Password, m.Host, m.Port, m.Database,
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
func (m *MySQL) Conn() (*gorm.DB, error) {
	// 填写注册的mysql网络
	conn, err := gorm.Open(
		sql.Open(
			fmt.Sprintf(
				"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				m.User, m.Password, m.Host, m.Port, m.Database,
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
