package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type AccountRepo struct {
	DB *sqlx.DB
}

type Account struct {
	Id          int    `db:"id" json:"id"`
	AccountName string `db:"account_name" json:"accountName"`
	AccountDesc string `db:"account_desc" json:"accountDesc"`
}

func (repo *AccountRepo) CreateAccount(parentLog *logrus.Entry, accountName string, accountDesc string) (*Account, error) {
	log := RepoFunctionLogger(parentLog, "CreateAccount")
	defer log.Info("Returned")

	SQL := `--sql
		INSERT INTO account (account_name, account_desc)
		VALUES ($1, $2)
		RETURNING *`

	var createdAccount Account

	return &createdAccount, Tx(repo.DB, func(tx *sqlx.Tx) error {
		result := tx.QueryRowx(SQL, accountName, accountDesc)

		return result.StructScan(&createdAccount)
	})
}

func (repo *AccountRepo) GetAccount(parentLog *logrus.Entry, accountID int) (*Account, error) {
	log := RepoFunctionLogger(parentLog, "GetAccount")
	defer log.Info("Returned")

	SQL := `--sql
		SELECT * FROM account
		WHERE account.id = $1`

	var account Account

	return &account, Tx(repo.DB, func(tx *sqlx.Tx) error {
		return tx.Get(&account, SQL, accountID)
	})
}

// TODO: pagination
func (repo *AccountRepo) GetAccounts(parentLog *logrus.Entry) ([]*Account, error) {
	log := RepoFunctionLogger(parentLog, "GetAccounts")
	defer log.Info("Returned")

	SQL := `--sql
		SELECT * FROM account`

	accounts := make([]*Account, 0)

	return accounts, Tx(repo.DB, func(tx *sqlx.Tx) error {
		return tx.Select(&accounts, SQL)
	})
}