package goutils

import (
	"fmt"

	"errors"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

func NewMysqlDB(host string, port int, user, pwd, dbName string, maxIdle int) (*sqlx.DB, error) {
	str := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8", user, pwd, host, port, dbName)

	db := sqlx.MustOpen("mysql", str)
	if err := db.Ping(); err != nil { // 检查连接性
		return nil, errors.New(fmt.Sprintf("%v %v", err, str))
	}
	db.SetMaxIdleConns(maxIdle)
	return db, nil
}

func NewPGDB(host string, port int, user, pwd, dbName string, maxIdle int) (*sqlx.DB, error) {
	str := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pwd, dbName)

	db := sqlx.MustOpen("postgres", str)
	if err := db.Ping(); err != nil { // 检查连接性
		return nil, errors.New(fmt.Sprintf("%v %v", err, str))
	}
	db.SetMaxIdleConns(maxIdle)
	return db, nil
}