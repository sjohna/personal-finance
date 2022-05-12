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

		id, err := repo.GetNextEntityId(tx)
		if err != nil {
			txLog.WithError(err).Error()
			return err
		}

		params := repo.CreateCurrencyParams{
			id,
			name,
			abbreviation,
			magnitude,
		}

		if err = repo.HandleCreateSingleEntityFromApiCall(tx, "create", "currency", params); err != nil {
			return err
		}

		currency, err = repo.CreateCurrency(tx, params)
		if err != nil {
			txLog.WithError(err).Error()
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return currency, nil
}
