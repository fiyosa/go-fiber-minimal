package lib

import (
	"fmt"
	"go-fiber-ddd/config"
	"go-fiber-ddd/database/entity"
	"log"
	"os"
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
		`host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=Asia/Jakarta`,
		config.Db.Host,
		config.Db.User,
		config.Db.Password,
		config.Db.Name,
		config.Db.Port,
		config.Db.SSLMode,
	)

	var setLogger logger.Interface
	if config.Env.APP_ENV != "development" {
		setLogger = logger.Default.LogMode(logger.Silent)
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

	return logger.New(
		log.New(file, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             0,           // Disable slow threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
}
