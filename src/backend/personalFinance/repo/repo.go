package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Repo struct {
	DB *sqlx.DB
}

func (repo *Repo) NonTx(logger *logrus.Entry) *DBDAO {
	return newDBDAO(repo.DB, logger)
}
