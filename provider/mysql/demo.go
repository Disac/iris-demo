package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go-iris-frame-demo/etc/config"
)

var DemoDB = demo()

func demo() *sqlx.DB {
	c := &client{
		driver:      config.Val("mysql.driver"),
		address:     config.Val("mysql.address"),
		maxOpenConn: config.Int("mysql.maxOpenConn"),
		maxIdleConn: config.Int("mysql.maxIdleConn"),
		maxLifeTime: config.Int("mysql.maxLifeTime"),
	}
	c.Init()
	return c.db
}
