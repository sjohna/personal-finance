package pfdb

// initialization functions for database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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
	account_id int NOT NULL,
	CONSTRAINT fk_account 
      FOREIGN KEY(account_id)
        REFERENCES account(id)
);`

var createDebitTable = `
CREATE TABLE IF NOT EXISTS debit (
	id serial PRIMARY KEY,
	debitAmount DECIMAL(20,10),
	time timestamp,
	account_id int NOT NULL,
	CONSTRAINT fk_account 
      FOREIGN KEY(account_id)
        REFERENCES account(id)
);`

var createTransactionTable = `
CREATE TABLE IF NOT EXISTS transaction (
	id serial PRIMARY KEY,
	credit_id int NOT NULL UNIQUE,
	debit_id int NOT NULL UNIQUE,
	CONSTRAINT fk_credit 
      FOREIGN KEY(credit_id)
        REFERENCES credit(id),
	CONSTRAINT fk_debit 
      FOREIGN KEY(debit_id)
        REFERENCES debit(id)
);`

func createTable(db *sqlx.DB, SQL string, tableName string) error {
	if _, err := db.Exec(createAccountTable); err != nil {
		return err
	}

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
