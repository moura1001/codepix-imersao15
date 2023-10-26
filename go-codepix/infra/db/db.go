package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/moura1001/codepix/domain/model"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
		config  *gorm.Config
	)

	if os.Getenv("DEBUG") == "true" {
		config = &gorm.Config{Logger: logger.Default.LogMode(logger.Info)}
	} else {
		config = &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)}
	}

	if envMode != "dev" && envMode != "test" {
		dsn = os.Getenv("DSN")
		sqlDB, errDb := sql.Open("pgx", dsn)
		if errDb == nil {
			db, err = gorm.Open(postgres.New(postgres.Config{
				Conn: sqlDB,
			}), config)
		} else {
			err = errDb
		}
	} else {
		db, err = GetDBConnectionTest()
	}

	if err != nil {
		log.Fatalf("error to setup database. Details: '%s'", err)
	}

	if os.Getenv("AUTO_MIGRATE_DB") == "true" {
		db.AutoMigrate(&model.Bank{}, &model.Account{}, &model.PixKey{}, &model.Transaction{})
	}

	return db
}

func GetDBConnectionTest() (*gorm.DB, error) {
	var (
		dsn    string = os.Getenv("DSN_TEST")
		db     *gorm.DB
		err    error
		config *gorm.Config
	)

	if os.Getenv("DEBUG") == "true" {
		config = &gorm.Config{Logger: logger.Default.LogMode(logger.Info)}
	} else {
		config = &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)}
	}

	db, err = gorm.Open(sqlite.Open(dsn), config)

	if err != nil {
		log.Fatalf("error connecting to test database. Details: '%s'", err)
	}

	if os.Getenv("AUTO_MIGRATE_DB") == "true" {
		err = db.AutoMigrate(&model.Bank{}, &model.Account{}, &model.PixKey{}, &model.Transaction{})
	}

	return db, err
}
