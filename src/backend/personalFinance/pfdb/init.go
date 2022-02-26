package pfdb

// initialization functions for database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// maybe add type column at some point
var createAccountTable = `
CREATE TABLE IF NOT EXISTS account (
    id serial PRIMARY KEY,
	account_name VARCHAR,
	account_desc VARCHAR
);`

var createCreditTable = `
CREATE TABLE IF NOT EXISTS credit (
	id serial PRIMARY KEY,
	creditAmount DECIMAL(20,10),
	time timestamp,
	account_id int NOT NULL REFERENCES account
);`

var createDebitTable = `
CREATE TABLE IF NOT EXISTS debit (
	id serial PRIMARY KEY,
	debitAmount DECIMAL(20,10),
	time timestamp,
	account_id int NOT NULL REFERENCES account
);`

var createTransactionTable = `
CREATE TABLE IF NOT EXISTS transaction (
	id serial PRIMARY KEY,
	credit_id int NOT NULL UNIQUE REFERENCES credit,
	debit_id int NOT NULL UNIQUE REFERENCES debit
);`

func createTable(db *sqlx.DB, SQL string, tableName string) error {
	log := logrus.WithFields(logrus.Fields{
		"startup":     true,
		"createTable": tableName,
	})

	log.Info("Creating table")
	_, err := db.Exec(SQL)
	if err != nil {
		log.WithError(err).WithField("SQL", SQL).Error("Error creating table")
		return err
	}
	log.Info("Successful")

	return nil
}

func CreateTables(db *sqlx.DB) error {
	// TODO: logging
	if err := createTable(db, createAccountTable, "account"); err != nil {
		return err
	}

	if err := createTable(db, createCreditTable, "credit"); err != nil {
		return err
	}

	if err := createTable(db, createDebitTable, "debit"); err != nil {
		return err
	}

	if err := createTable(db, createTransactionTable, "transaction"); err != nil {
		return err
	}

	return nil
}
