package repo

import "time"

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

func CreateCurrency(dao DAO, id int64, params CreateCurrencyParams) (*Currency, error) {
	log := repoFunctionLogger(dao.Logger(), "CreateCurrency")
	defer logRepoReturn(log)

	SQL := `--sql
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
