package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func RepoFunctionLogger(log *logrus.Entry, repoFunction string) *logrus.Entry {
	log = log.WithField("repo-function", repoFunction)
	log.Info("Called")
	return log
}

func Tx(db *sqlx.DB, txFunc func(tx *sqlx.Tx) error) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	// TODO: think about error handling and logging here
	err = txFunc(tx)

	if err == nil {
		err = tx.Commit()
	}

	if err != nil {
		tx.Rollback()
	}

	return err
}
