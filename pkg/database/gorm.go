package database

import (
	"database/sql"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type gormConnector struct {
	gormMysql *gorm.DB
}

type GormConnector interface {
	GormMysql() *gorm.DB
}

func (c *gormConnector) GormMysql() *gorm.DB {
	return c.gormMysql
}

func InitGorm() GormConnector {

	gormConn := new(gormConnector)

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsnHost := os.Getenv("MYSQL_HOST")
	sqlDB, err := sql.Open("mysql", dsnHost)
	if err != nil {
		panic(err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	//  check connection
	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	gormConn.gormMysql = db

	return gormConn
}
