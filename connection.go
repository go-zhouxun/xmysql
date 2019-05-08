package xmysql

import (
	"database/sql"

	"github.com/zhouxun1995/xlog"
)

type XDBSession struct {
	db          *sql.DB
	Transaction bool
	tx          *sql.Tx
	Finished    bool

	Logger xlog.XLog
}

func (mysql MySQL) NewDBSession(tx bool, logger xlog.XLog) *XDBSession {
	if tx {
		transaction, err := mysql.database.Begin()
		if err != nil {
			return nil
		}
		return &XDBSession{
			db:          mysql.database,
			Transaction: true,
			Finished:    false,
			tx:          transaction,
			Logger:      logger,
		}
	}
	return &XDBSession{
		db:          mysql.database,
		Transaction: false,
		tx:          nil,
		Logger:      logger,
	}
}

func (xmysql *XDBSession) Exec(query string, args ...interface{}) (sql.Result, error) {
	if xmysql.Transaction && !xmysql.Finished {
		return xmysql.tx.Exec(query, args...)
	} else {
		return xmysql.db.Exec(query, args...)
	}
}

func (xmysql *XDBSession) QueryRow(query string, args ...interface{}) *sql.Row {
	if xmysql.Transaction && !xmysql.Finished {
		return xmysql.tx.QueryRow(query, args...)
	} else {
		return xmysql.db.QueryRow(query, args...)
	}
}

func (xmysql *XDBSession) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if xmysql.Transaction && !xmysql.Finished {
		return xmysql.tx.Query(query, args...)
	} else {
		return xmysql.db.Query(query, args...)
	}
}

func (xmysql *XDBSession) Commit() error {
	if xmysql.Transaction && !xmysql.Finished {
		err := xmysql.tx.Commit()
		if err == nil {
			xmysql.Finished = true
		}
		return err
	}
	return nil
}

func (xmysql *XDBSession) Rollback() error {
	if xmysql.Transaction && !xmysql.Finished {
		err := xmysql.tx.Rollback()
		if err == nil {
			xmysql.Finished = true
		}
		return err
	}
	return nil
}
