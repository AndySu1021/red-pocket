package db

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseType string

const (
	MySQL  DatabaseType = "mysql"
	SQLite DatabaseType = "sqlite"
)

type Config struct {
	Debug              bool
	Type               DatabaseType
	Host               string
	Port               int
	Username           string
	Password           string
	DBName             string
	MaxIdleConnections int `mapstructure:"max_idle_connections"`
	MaxOpenConnections int `mapstructure:"max_open_connections"`
	MaxLifetimeSec     int
	WithColor          bool `mapstructure:"with_color"`
}

func GetConnectionStr(cfg *Config) (connectionString string, err error) {
	switch cfg.Type {
	case MySQL:
		connectionString = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&multiStatements=true&parseTime=true", cfg.Username, cfg.Password, cfg.Host+":"+strconv.Itoa(cfg.Port), cfg.DBName)
	case SQLite:
		if cfg.Host == "" {
			connectionString = path.Join(os.Getenv("PROJ_DIR"), "test/.data", "sqlite.db?cache=shared")
		} else {
			connectionString = cfg.Host
		}
	default:
		return "", errors.New("not support driver")
	}

	return
}

func NewDatabase(cfg *Config) (db *gorm.DB, err error) {
	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = time.Duration(180) * time.Second

	var dialector gorm.Dialector

	dsn, err := GetConnectionStr(cfg)
	if err != nil {
		return nil, err
	}

	switch cfg.Type {
	case MySQL:
		dialector = mysql.Open(dsn)
	case SQLite:
		dialector = sqlite.Open(dsn)
	default:
		return nil, errors.New("not support db type")
	}

	log.Debug().Msgf("main: database connection string: %s", dsn)

	colorful := false
	logLevel := logger.Silent
	if cfg.Debug {
		logLevel = logger.Info
		colorful = cfg.WithColor
	}

	newLogger := NewLogger(logger.Config{
		SlowThreshold: time.Second, // Slow SQL threshold
		LogLevel:      logLevel,    // Log level
		Colorful:      colorful,    // Disable color
	})

	err = backoff.Retry(func() error {
		db, err = gorm.Open(dialector, &gorm.Config{
			Logger:                                   newLogger,
			PrepareStmt:                              true,
			DisableForeignKeyConstraintWhenMigrating: true,
		})
		if err != nil {
			log.Error().Msgf("Fail to open conn, err: %+v", err)
			return err
		}

		sqlDB, err := db.DB()
		if err != nil {
			log.Error().Msgf("Fail to get DB, err: %+v", err)
			return err
		}

		err = sqlDB.Ping()
		return err
	}, bo)

	if err != nil {
		log.Error().Msgf("main: database connect err: %s", err.Error())
		return nil, err
	}

	log.Info().Msgf("database ping success")

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if cfg.MaxIdleConnections != 0 {
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConnections)
	} else {
		sqlDB.SetMaxIdleConns(2)
	}

	if cfg.MaxOpenConnections != 0 {
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConnections)
	} else {
		sqlDB.SetMaxOpenConns(5)
	}

	if cfg.MaxLifetimeSec != 0 {
		sqlDB.SetConnMaxLifetime(time.Duration(cfg.MaxLifetimeSec) * time.Second)
	} else {
		sqlDB.SetConnMaxLifetime(14400 * time.Second)
	}

	return db, nil
}
