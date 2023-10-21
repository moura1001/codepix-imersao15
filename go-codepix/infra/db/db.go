package db

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/moura1001/codepix/domain/model"
	_ "gorm.io/driver/sqlite"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)

	err := godotenv.Load(filepath.Join(basePath, "..", "..", ".env.example"))
	if err != nil {
		log.Fatalf("error to load .env file. Details: '%s'", err)
	}
}

func GetDBConnection() *gorm.DB {
	var (
		envMode = os.Getenv("ENV_MODE")
		dsn     string
		db      *gorm.DB
		err     error
	)

	if envMode != "dev" && envMode != "test" {
		dsn = os.Getenv("DSN")
		db, err = gorm.Open(os.Getenv("DB_TYPE"), dsn)
	} else {
		dsn = os.Getenv("DSN_TEST")
		db, err = gorm.Open(os.Getenv("DB_TYPE_TEST"), dsn)
	}

	if err != nil {
		log.Fatalf("error connecting to database. Details: '%s'", err)
	}

	if os.Getenv("DEBUG") == "true" {
		db.LogMode(true)
	}

	if os.Getenv("AUTO_MIGRATE_DB") == "true" {
		db.AutoMigrate(&model.Bank{}, &model.Account{}, &model.PixKey{}, &model.Transaction{})
	}

	return db
}
