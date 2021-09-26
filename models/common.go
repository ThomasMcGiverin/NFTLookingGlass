package models

import (
	"github.com/lib/pq"
	"github.com/thomasmcgiverin/NFTLookingGlass/server/config"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
	gormtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorm.io/gorm.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
)

var (
	db         *gorm.DB
	readOnlyDB *gorm.DB
	//redisClient *redis.Client

	dbLock       sync.Mutex
	readOnlyLock sync.Mutex
	//redisLock    sync.Mutex
)

func DB() (*gorm.DB, error) {
	dbLock.Lock()
	defer dbLock.Unlock()

	if db != nil {
		return db, nil
	}

	db, err := ConnectDB(config.Cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ConnectDB(url string) (*gorm.DB, error) {
	sqltrace.Register("postgres", &pq.Driver{}, sqltrace.WithServiceName(config.Cfg.ServiceName))
	dbConn, err := sqltrace.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	conn, err := gormtrace.Open(
		postgres.New(postgres.Config{Conn: dbConn}),
		&gorm.Config{Logger: logger.Discard})

	if err != nil {
		return nil, err
	}

	genericDB, err := conn.DB()
	if err != nil {
		return nil, err
	}

	genericDB.SetMaxIdleConns(config.Cfg.MaxIdleConn)
	genericDB.SetMaxOpenConns(config.Cfg.MaxOpenConn)
	return conn, nil
}
