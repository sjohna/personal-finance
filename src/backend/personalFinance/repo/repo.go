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

func (repo *Repo) SerializableTx(logger *logrus.Entry) (*TxDAO, error) {
	dao, err := newTXDAO(repo.DB, logger)
	if err != nil {
		logger.WithError(err).Error("Error creating TxDAO in SerializableTx")
		return nil, err
	}

	_, err = dao.Exec("set transaction isolation level serializable")
	if err != nil {
		dao.Logger().WithError(err).Error("Failed to set transaction isolation level serializable")
		dao.Close()
		return nil, err
	}

	return dao, nil
}
