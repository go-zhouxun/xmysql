package xmysql

import (
	"database/sql"
	"fmt"

	"github.com/zhouxun1995/xlog"
)

type MySQL struct {
	database *sql.DB
	logger   xlog.XLog
}

func (mysql *MySQL) GetConnection() *XDBSession {
	return mysql.NewDBSession(false, mysql.logger)
}

func (mysql *MySQL) Begin() *XDBSession {
	return mysql.NewDBSession(true, mysql.logger)
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
