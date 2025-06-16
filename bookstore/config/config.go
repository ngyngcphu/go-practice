package config

import (
	"database/sql"
	"errors"
	"sync"
)

var (
	db   *sql.DB
	once sync.Once
)

func InitDB(dataSourceName string) error {
	var err error

	once.Do(func() {
		db, err = sql.Open("postgres", dataSourceName)
		if err != nil {
			return
		}
		err = db.Ping()
	})
	return err
}

func GetDB() (*sql.DB, error) {
	if db == nil {
		return nil, errors.New("database not initialized - call InitDB first")
	}
	return db, nil
}

func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
