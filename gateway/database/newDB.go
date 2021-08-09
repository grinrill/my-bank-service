package database

import (
	"database/sql"

	"github.com/grinrill/my-bank-service/domain/bankRepo"

	// import to enable "mysql" database driver
	_ "github.com/go-sql-driver/mysql"
)

func newDB() (db *sql.DB, cancel bankRepo.CancelFunc, err error) {
	db, err = sql.Open("mysql", "/bankAccounts.db")
	if err != nil {return}

	cancel = func() error {
		return db.Close()
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS accounts (
			accountID INT,
			currency CHAR(3) DEFAULT 'SBP',
			balance FLOAT DEFAULT 0,
			PRIMARY KEY (accountID)
		)`,
	)
	if err != nil {return}

	// I've never work with sql database in go
	// so I really dont't know do I really need these
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return
}