package service

import (
	"github.com/sirupsen/logrus"
	"github.com/sjohna/personal-finance/repo"
)

type CurrencyService struct {
	Repo *repo.Repo
}

func (svc *CurrencyService) CreateCurrency(logger *logrus.Entry, name string, abbreviation string, magnitude int) (*repo.Currency, error) {
	log := serviceFunctionLogger(logger, "CreateCurrency")
	defer logServiceReturn(log)

	var currency *repo.Currency

	err := svc.Repo.SerializableTx(log, func(tx *repo.TxDAO) error {
		txLog := tx.Logger()

		params := repo.CreateCurrencyParams{
			name,
			abbreviation,
			magnitude,
		}

		id, time, err := repo.HandleCreateSingleEntityFromApiCall(tx, "create", "currency", params)
		if err != nil {
			return err
		}

		currency, err = repo.CreateCurrency(tx, id, params)
		if err != nil {
			txLog.WithError(err).Error()
			return err
		}

		currency.CreatedAt = time
		currency.UpdatedAt = time

		return nil
	})

	if err != nil {
		return nil, err
	}

	return currency, nil
}

func (svc *CurrencyService) GetCurrency(logger *logrus.Entry, accountID int64) (*repo.Currency, error) {
	log := serviceFunctionLogger(logger, "GetCurrency")
	defer logServiceReturn(log)

	dao := svc.Repo.NonTx(log)

	currency, err := repo.GetCurrency(dao, accountID)
	if err != nil {
		log.WithError(err).Error()
	}

	return currency, err
}

// TODO: pagination
func (svc *CurrencyService) GetCurrencies(logger *logrus.Entry) ([]*repo.Currency, error) {
	log := serviceFunctionLogger(logger, "GetCurrencies")
	defer logServiceReturn(log)

	dao := svc.Repo.NonTx(log)

	currencies, err := repo.GetCurrencies(dao)
	if err != nil {
		log.WithError(err).Error()
	}

	return currencies, err
}
