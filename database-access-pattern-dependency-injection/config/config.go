package config

import (
	"database/sql"
	"errors"
	"sync"
)

type Env struct {
	db   *sql.DB
	once sync.Once
}

func NewEnv() *Env {
	return &Env{}
}

func (env *Env) InitDB(dataSourceName string) error {
	var err error
	env.once.Do(func() {
		env.db, err = sql.Open("postgres", dataSourceName)
		if err != nil {
			return
		}
		err = env.db.Ping()
	})
	return err
}

func (env *Env) GetDB() (*sql.DB, error) {
	if env.db == nil {
		return nil, errors.New("database not initialized - call InitDB first")
	}
	return env.db, nil
}

func (env *Env) CloseDB() error {
	if env.db != nil {
		return env.db.Close()
	}
	return nil
}
