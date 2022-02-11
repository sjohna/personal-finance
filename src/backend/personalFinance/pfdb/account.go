package pfdb

import "github.com/jmoiron/sqlx"

type Account struct {
	Id          int    `db:"id" json:"id"`
	AccountName string `db:"account_name" json:"accountName"`
	AccountDesc string `db:"account_desc" json:"accountDesc"`
}

// TODO: database transaction helper function

func CreateAccount(db *sqlx.DB, accountName string, accountDesc string) (*Account, error) {
	SQL := `
		INSERT INTO account (account_name, account_desc)
		VALUES ($1, $2)
		RETURNING *`

	var createdAccount Account

	return &createdAccount, Tx(db, func(tx *sqlx.Tx) error {
		result := tx.QueryRowx(SQL, accountName, accountDesc)

		return result.StructScan(&createdAccount)
	})
}

func GetAccount(db *sqlx.DB, accountID int) (*Account, error) {
	SQL := `
		SELECT * FROM account
		WHERE account.id = $1`

	var account Account

	return &account, Tx(db, func(tx *sqlx.Tx) error {
		return tx.Get(&account, SQL, accountID)
	})
}

// TODO: pagination
func GetAccounts(db *sqlx.DB) ([]*Account, error) {
	SQL := `
		SELECT * FROM account`

	accounts := make([]*Account, 0)

	return accounts, Tx(db, func(tx *sqlx.Tx) error {
		return tx.Select(&accounts, SQL)
	})
}
