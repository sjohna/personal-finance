package pfdb

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Account struct {
	Id          int    `db:"id" json:"id"`
	AccountName string `db:"account_name" json:"accountName"`
	AccountDesc string `db:"account_desc" json:"accountDesc"`
}

func repoFunctionLogger(log *logrus.Entry, repoFunction string) *logrus.Entry {
	log = log.WithField("repoFunction", repoFunction)
	log.Info("Called")
	return log
}

func CreateAccount(parentLog *logrus.Entry, db *sqlx.DB, accountName string, accountDesc string) (*Account, error) {
	log := repoFunctionLogger(parentLog, "CreateAccount")
	defer log.Info("Returned")

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

func GetAccount(parentLog *logrus.Entry, db *sqlx.DB, accountID int) (*Account, error) {
	log := repoFunctionLogger(parentLog, "GetAccount")
	defer log.Info("Returned")

	SQL := `
		SELECT * FROM account
		WHERE account.id = $1`

	var account Account

	return &account, Tx(db, func(tx *sqlx.Tx) error {
		return tx.Get(&account, SQL, accountID)
	})
}

// TODO: pagination
func GetAccounts(parentLog *logrus.Entry, db *sqlx.DB) ([]*Account, error) {
	log := repoFunctionLogger(parentLog, "GetAccounts")
	defer log.Info("Returned")

	SQL := `
		SELECT * FROM account`

	accounts := make([]*Account, 0)

	return accounts, Tx(db, func(tx *sqlx.Tx) error {
		return tx.Select(&accounts, SQL)
	})
}
