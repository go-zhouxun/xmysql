package xmysql

import (
	"database/sql"
	"fmt"
)

type MySQL struct {
	database *sql.DB
}

func (mysql *MySQL) GetConnection() *XConnection {
	return mysql.NewXMySQL(false)
}

func (mysql *MySQL) Begin() *XConnection {
	return mysql.NewXMySQL(true)
}

func NewMySQl(config XMySQLConfig) *MySQL {
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", config.User, config.Password, config.Address, config.Port, config.DBName)
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil
	}
	db.SetMaxIdleConns(config.MaxIdle)
	db.SetMaxOpenConns(config.MaxConn)
	return &MySQL{database: db}
}
