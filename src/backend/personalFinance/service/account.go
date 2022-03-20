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

	dao := svc.Repo.NonTx(log)
	defer dao.Close()
	account, err := repo.CreateAccount(dao, accountName, accountDesc)
	if err != nil {
		log.WithError(err).Error()
	}

	return account, err
}

func (svc *AccountService) GetAccount(logger *logrus.Entry, accountID int) (*repo.Account, error) {
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
