package mysql

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type client struct {
	driver      string
	address     string
	maxOpenConn int
	maxIdleConn int
	maxLifeTime int
	db          *sqlx.DB
}

func (c *client) Init() {
	db := sqlx.MustOpen(c.driver, c.address)
	db.SetConnMaxLifetime(time.Duration(c.maxLifeTime) * time.Second)
	db.SetMaxIdleConns(c.maxIdleConn)
	db.SetMaxOpenConns(c.maxOpenConn)
	c.db = db

	err := db.Ping()
	if err != nil {
		panic(err)
	}
}
