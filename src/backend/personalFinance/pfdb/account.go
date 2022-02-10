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

	tx, err := db.Beginx()
	if err != nil {
		return nil, err
	}

	result := tx.QueryRowx(SQL, accountName, accountDesc)
	var createdAccount Account
	err = result.StructScan(&createdAccount)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &createdAccount, nil
}

func GetAccount(db *sqlx.DB, accountID int) (*Account, error) {
	SQL := `
		SELECT * FROM account
		WHERE account.id = $1`

	tx, err := db.Beginx()
	if err != nil {
		return nil, err
	}

	var account Account
	err = tx.Get(&account, SQL, accountID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &account, nil
}

// TODO: pagination
func GetAccounts(db *sqlx.DB) ([]*Account, error) {
	SQL := `
		SELECT * FROM account`

	tx, err := db.Beginx()
	if err != nil {
		return nil, err
	}

	accounts := make([]*Account, 0)
	if err = tx.Select(&accounts, SQL); err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return accounts, nil
}
