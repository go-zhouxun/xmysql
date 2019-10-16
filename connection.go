package xmysql

import (
	"database/sql"
)

type XConnection struct {
	db          *sql.DB
	Transaction bool
	tx          *sql.Tx
	Finished    bool
}

func (mysql MySQL) BeginTx() *XConnection {
	transaction, err := mysql.database.Begin()
	if err != nil {
		return nil
	}
	return &XConnection{
		db:          mysql.database,
		Transaction: true,
		Finished:    false,
		tx:          transaction,
	}
}

func (mysql MySQL) GetConnection() *XConnection {
	return &XConnection{
		db:          mysql.database,
		Transaction: false,
		tx:          nil,
	}
}

func (xmysql *XConnection) Exec(query string, args ...interface{}) (sql.Result, error) {
	if xmysql.Transaction && !xmysql.Finished {
		return xmysql.tx.Exec(query, args...)
	} else {
		return xmysql.db.Exec(query, args...)
	}
}

func (xmysql *XConnection) QueryRow(query string, args ...interface{}) *sql.Row {
	if xmysql.Transaction && !xmysql.Finished {
		return xmysql.tx.QueryRow(query, args...)
	} else {
		return xmysql.db.QueryRow(query, args...)
	}
}

func (xmysql *XConnection) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if xmysql.Transaction && !xmysql.Finished {
		return xmysql.tx.Query(query, args...)
	} else {
		return xmysql.db.Query(query, args...)
	}
}

func (xmysql *XConnection) Commit() error {
	if xmysql.Transaction && !xmysql.Finished {
		err := xmysql.tx.Commit()
		if err == nil {
			xmysql.Finished = true
		}
		return err
	}
	return nil
}

func (xmysql *XConnection) Rollback() error {
	if xmysql.Transaction && !xmysql.Finished {
		err := xmysql.tx.Rollback()
		if err == nil {
			xmysql.Finished = true
		}
		return err
	}
	return nil
}
