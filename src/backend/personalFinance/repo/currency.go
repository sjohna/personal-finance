package repo

import (
	"time"

	sq "github.com/elgris/sqrl"
)

type Currency struct {
	Id           int64     `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	Abbreviation string    `db:"abbreviation" json:"abbreviation"`
	Magnitude    int       `db:"magnitude" json:"magnitude"`
	CreatedAt    time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt    time.Time `db:"updated_at" json:"updatedAt"`
}

type CreateCurrencyParams struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	Magnitude    int    `json:"magnitude"`
}

func currencySelect() *sq.SelectBuilder {
	return sq.Select("currency.*", "create_action.time as created_at", "update_action.time as updated_at").
		From("currency").
		Join("event create_event on create_event.entity_id = currency.id and create_event.event_type = 'create'").
		Join("action create_action on create_event.action_id = create_action.id").
		JoinClause(`join lateral
							(
								select action.time 
								from event
								join action on event.action_id = action.id
								where event.entity_id = currency.id
								order by time desc
								limit 1
							) update_action on true`)
}

func CreateCurrency(dao DAO, id int64, params CreateCurrencyParams) (*Currency, error) {
	log := repoFunctionLogger(dao.Logger(), "CreateCurrency")
	defer logRepoReturn(log)

	// language=SQL
	SQL := `
		INSERT INTO currency (id, name, abbreviation, magnitude)
		VALUES ($1, $2, $3, $4)
		RETURNING *`

	var createdCurrency Currency
	err := dao.Get(&createdCurrency, SQL, id, params.Name, params.Abbreviation, params.Magnitude)
	if err != nil {
		log.WithError(err).Error()
	}

	return &createdCurrency, err
}

func GetCurrency(dao DAO, id int64) (*Currency, error) {
	log := repoFunctionLogger(dao.Logger(), "GetCurrency")
	defer logRepoReturn(log)

	SQL, _, err := currencySelect().Where("currency.id = $1").ToSql()
	if err != nil {
		log.WithError(err).Fatal("SQL error initializing query")
	}

	var currency Currency
	err = dao.Get(&currency, SQL, id)
	if err != nil {
		log.WithError(err).Error()
	}

	return &currency, err
}

// TODO: pagination
func GetCurrencies(dao DAO) ([]*Currency, error) {
	log := repoFunctionLogger(dao.Logger(), "GetCurrencies")
	defer logRepoReturn(log)

	SQL, _, err := currencySelect().ToSql()
	if err != nil {
		log.WithError(err).Fatal("SQL error initializing query")
	}

	currencies := make([]*Currency, 0)
	err = dao.Select(&currencies, SQL)
	if err != nil {
		log.WithError(err).Error()
	}

	return currencies, err
}
