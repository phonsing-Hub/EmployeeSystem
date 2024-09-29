package db

import (
	"fmt"
	"os"
	"time"
	"github.com/phonsing-Hub/EmployeeSystem/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type Db struct {
	DB *gorm.DB
}

func New(dbuser string, dbpass string, dbhost string, dbname string) (*Db, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbuser, dbpass, dbhost, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: loggerConfig(true)})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(
		&models.Region{},
		&models.Country{},
		&models.Location{},
		&models.Job{},
		&models.Department{},
		&models.Employee{},
		&models.Dependent{},
		&models.AuthUser{},
		&models.Token{},
	)

	if err != nil {
		log.Fatalf("ไม่สามารถ migrate ตารางได้: %v", err)
	}
	return &Db{DB: db}, nil
}

func loggerConfig(enable bool) logger.Interface {
	if enable {
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Set log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound errors
				ParameterizedQueries:      true,        // Don't include raw SQL queries in logs
				Colorful:                  true,        // Colorize logs
			},
		)
		return newLogger
	}

	// Default silent logger if not enabled
	return logger.Default.LogMode(logger.Silent)
}
