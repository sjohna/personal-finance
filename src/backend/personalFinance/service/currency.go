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

		id, err := repo.HandleCreateSingleEntityFromApiCall(tx, "create", "currency", params)
		if err != nil {
			return err
		}

		currency, err = repo.CreateCurrency(tx, id, params)
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
