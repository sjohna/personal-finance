package service

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/sjohna/personal-finance/repo"
)

type AccountService struct {
	Repo *repo.Repo
}

func (svc *AccountService) CreateAccount(logger *logrus.Entry, name string, description string) (*repo.Account, error) {
	log := serviceFunctionLogger(logger, "CreateAccount")
	defer logServiceReturn(log)

	var account *repo.Account

	err := svc.Repo.SerializableTx(log, func(tx *repo.TxDAO) error {
		txLog := tx.Logger()

		params := repo.CreateAccountParams{
			name,
			description,
		}

		id, time, err := repo.HandleCreateSingleEntityFromApiCall(tx, "create", "account", params)
		if err != nil {
			return err
		}

		account, err = repo.CreateAccount(tx, id, params)
		if err != nil {
			txLog.WithError(err).Error()
			return err
		}

		account.CreatedAt = time
		account.UpdatedAt = time

		return nil
	})

	if err != nil {
		return nil, err
	}

	return account, nil
}

func (svc *AccountService) GetAccount(logger *logrus.Entry, accountId int64) (*repo.Account, error) {
	log := serviceFunctionLogger(logger, "GetAccount")
	defer logServiceReturn(log)

	dao := svc.Repo.NonTx(log)

	account, err := repo.GetAccount(dao, accountId)
	if err != nil {
		log.WithError(err).Error()
		return nil, err
	}

	return account, err
}

func (svc *AccountService) GetAccountWithDebitsAndCredits(logger *logrus.Entry, accountId int64) (*repo.Account, error) {
	log := serviceFunctionLogger(logger, "GetAccountWithDebitsAndCredits")
	defer logServiceReturn(log)

	dao := svc.Repo.NonTx(log)

	account, err := repo.GetAccount(dao, accountId)
	if err != nil {
		log.WithError(err).Error()
		return nil, err
	}

	debitsAndCredits, err := repo.GetDebitsAndCreditsForAccount(dao, accountId)
	if err != nil {
		log.WithError(err).Error()
		return nil, err
	}

	account.DebitsAndCredits = debitsAndCredits

	return account, err
}

// TODO: pagination
func (svc *AccountService) GetAccounts(logger *logrus.Entry) ([]*repo.Account, error) {
	log := serviceFunctionLogger(logger, "GetAccounts")
	defer logServiceReturn(log)

	dao := svc.Repo.NonTx(log)

	accounts, err := repo.GetAccounts(dao)
	if err != nil {
		log.WithError(err).Error()
		return nil, err
	}

	return accounts, err
}

func (svc *AccountService) CreateDebit(logger *logrus.Entry, accountId int64, amount int64, currencyId int64, time time.Time) (*repo.DebitCredit, error) {
	log := serviceFunctionLogger(logger, "CreateDebit")
	defer logServiceReturn(log)

	var debit *repo.DebitCredit

	err := svc.Repo.SerializableTx(log, func(tx *repo.TxDAO) error {
		txLog := tx.Logger()

		params := repo.CreateDebitCreditParams{
			amount,
			currencyId,
			time,
			accountId,
		}

		id, _, err := repo.HandleCreateSingleEntityFromApiCall(tx, "create", "debit", params)
		if err != nil {
			return err
		}

		debit, err = repo.CreateDebit(tx, id, params)
		if err != nil {
			txLog.WithError(err).Error()
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return debit, nil
}

func (svc *AccountService) CreateCredit(logger *logrus.Entry, accountId int64, amount int64, currencyId int64, time time.Time) (*repo.DebitCredit, error) {
	log := serviceFunctionLogger(logger, "CreateCredit")
	defer logServiceReturn(log)

	var debit *repo.DebitCredit

	err := svc.Repo.SerializableTx(log, func(tx *repo.TxDAO) error {
		txLog := tx.Logger()

		params := repo.CreateDebitCreditParams{
			amount,
			currencyId,
			time,
			accountId,
		}

		id, _, err := repo.HandleCreateSingleEntityFromApiCall(tx, "create", "credit", params)
		if err != nil {
			return err
		}

		debit, err = repo.CreateCredit(tx, id, params)
		if err != nil {
			txLog.WithError(err).Error()
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return debit, nil
}
