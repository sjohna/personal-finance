package repo

type Currency struct {
	Id           int64  `db:"id" json:"id"`
	Name         string `db:"name" json:"name"`
	Abbreviation string `db:"abbreviation" json:"abbreviation"`
	Magnitude    int    `db:"magnitude" json:"magnitude"`
}

type CreateCurrencyParams struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	Magnitude    int    `json:"magnitude"`
}

func CreateCurrency(dao DAO, params CreateCurrencyParams) (*Currency, error) {
	log := repoFunctionLogger(dao.Logger(), "CreateCurrency")
	defer logRepoReturn(log)

	SQL := `--sql
		INSERT INTO currency (id, name, abbreviation, magnitude)
		VALUES ($1, $2, $3, $4)
		RETURNING *`

	var createdCurrency Currency
	err := dao.Get(&createdCurrency, SQL, params.Id, params.Name, params.Abbreviation, params.Magnitude)
	if err != nil {
		log.WithError(err).Error()
	}

	return &createdCurrency, err
}
