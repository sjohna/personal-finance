package service

import (
	"github.com/sirupsen/logrus"
	"github.com/sjohna/personal-finance/repo"
)

type AccountService struct {
	Repo *repo.Repo
}

func (svc *AccountService) CreateAccount(logger *logrus.Entry, accountName string, accountDesc string) (*repo.Account, error) {
	log := serviceFunctionLogger(logger, "CreateAccount")
	defer logServiceReturn(log)

	// todo: refactor to make the tx be in a single function...
	tx, err := svc.Repo.SerializableTx(log)
	if err != nil {
		log.WithError(err).Error("Error creating DAO")
		return nil, err
	}
	defer tx.Close()

	action_id, err := repo.CreateAction(tx, "api-call")
	if err != nil {
		log.WithError(err).Error("Error creating action")
		return nil, err
	}

	id, err := repo.GetNextEntityId(tx)
	if err != nil {
		log.WithError(err).Error()
		return nil, err
	}

	params := repo.CreateAccountParams{
		id,
		accountName,
		accountDesc,
	}

	err = repo.CreateEvent(tx, action_id, "create", "account", params)
	if err != nil {
		log.WithError(err).Error("Error creating event")
		return nil, err
	}

	account, err := repo.CreateAccount(tx, params)
	if err != nil {
		log.WithError(err).Error()
		return nil, err
	}

	return account, nil
}

func (svc *AccountService) GetAccount(logger *logrus.Entry, accountID int64) (*repo.Account, error) {
	log := serviceFunctionLogger(logger, "GetAccount")
	defer logServiceReturn(log)

	dao := svc.Repo.NonTx(log)
	defer dao.Close()
	account, err := repo.GetAccount(dao, accountID)
	if err != nil {
		log.WithError(err).Error()
	}

	return account, err
}

// TODO: pagination
func (svc *AccountService) GetAccounts(logger *logrus.Entry) ([]*repo.Account, error) {
	log := serviceFunctionLogger(logger, "GetAccounts")
	defer logServiceReturn(log)

	dao := svc.Repo.NonTx(log)
	defer dao.Close()
	accounts, err := repo.GetAccounts(dao)
	if err != nil {
		log.WithError(err).Error()
	}

	return accounts, err
}
