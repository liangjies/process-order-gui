package initialize

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Sql Server初始化
func MSSQLGorm() *gorm.DB {
	return GormSqlServer()
}

//@author: liangjies
//@function: GormSqlServer
//@description: 初始化SqlServer数据库
//@return: *gorm.DB

func GormSqlServer() *gorm.DB {
	// 固化SQL Server配置
	Username := "root"
	Password := "root"
	Path := "127.0.0.1:1433"
	Dbname := "master"

	dsn := "sqlserver://" + Username + ":" + Password + "@" + Path + "?database=" + Dbname + "&encrypt=DISABLE"
	if db, err := gorm.Open(sqlserver.Open(dsn), config()); err != nil {
		fmt.Println("SqlServer连接异常")
		os.Exit(0)
		return nil
	} else {
		return db
	}
}

type writer struct {
	logger.Writer
}

// NewWriter writer 构造函数
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf 格式化打印日志
func (w *writer) Printf(message string, data ...interface{}) {
	w.Writer.Printf(message, data...)
}

func config() *gorm.Config {
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	_default := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})
	config.Logger = _default.LogMode(logger.Info)
	return config
}
