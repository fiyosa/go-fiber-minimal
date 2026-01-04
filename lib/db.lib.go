package lib

import (
	"context"
	"fmt"
	"go-fiber-minimal/config"
	"go-fiber-minimal/database/entity"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB dbManager

type dbManager struct {
	Run *gorm.DB
}

func (*dbManager) Entity() []any {
	return []any{
		&entity.User{},
		&entity.Auth{},
		&entity.Role{},
		&entity.Permission{},
		&entity.UserHasRole{},
		&entity.RoleHasPermission{},
	}
}

func (m *dbManager) Init() {
	config.Db.Init()

	dsn := fmt.Sprintf(
		`host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=Asia/Jakarta search_path=%v`,
		config.Db.Host,
		config.Db.User,
		config.Db.Password,
		config.Db.Name,
		config.Db.Port,
		config.Db.SSLMode,
		config.Db.Schema,
	)

	var setLogger logger.Interface
	if config.Env.APP_ENV != "local" {
		setLogger = logger.Default.LogMode(logger.Silent) // disable logging
	} else {
		setLogger = gormLogger()
	}

	connect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: config.Db.Schema + ".",
			// SingularTable: true,
		},
		SkipDefaultTransaction: true,
		Logger:                 setLogger,
		NowFunc: func() time.Time {
			return time.Now().Local() // timestamps
		},
	})

	if err != nil {
		fmt.Printf("Error access database: %v \n\n", err.Error())
		os.Exit(1)
	}

	m.Run = connect
}

func gormLogger() logger.Interface {
	dirLogs := "./logs"

	_, err := os.Stat(dirLogs)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(dirLogs, 0755); err != nil {
			fmt.Printf("Error creating directory logs: %v \n\n", err.Error())
			os.Exit(1)
		}
	}

	currentDate := time.Now().Format("2006-01-02")
	file, err := os.OpenFile(dirLogs+"/gorm.log"+currentDate+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error opening file: %v \n\n", err.Error())
		os.Exit(1)
	}

	return &gormCustomLogger{
		Writer: log.New(file, "", 0),
	}
}

type gormCustomLogger struct {
	Writer *log.Logger
}

func (l *gormCustomLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

func (l *gormCustomLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.log("INFO", fmt.Sprintf(msg, data...))
}

func (l *gormCustomLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.log("WARN", fmt.Sprintf(msg, data...))
}

func (l *gormCustomLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.log("ERROR", fmt.Sprintf(msg, data...))
}

func (l *gormCustomLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()

	msg := fmt.Sprintf("[%.3fms] [%d rows] %s", float64(elapsed.Nanoseconds())/1e6, rows, sql)
	level := "INFO"
	if err != nil {
		level = "ERROR"
		msg = fmt.Sprintf("%v %s", err, msg)
	}
	l.log(level, msg)
}

func (l *gormCustomLogger) log(level string, msg string) {
	now := time.Now().Format("2006:01:02 15:04:05")

	fileInfo := "unknown:0"
	for i := 2; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok && !strings.Contains(file, "gorm.io") && filepath.Base(file) != "db.lib.go" {
			fileInfo = fmt.Sprintf("%s:%d", filepath.Base(file), line)
			break
		}
	}

	l.Writer.Printf("[%s] %s %s %s\n", level, now, fileInfo, msg)
}
