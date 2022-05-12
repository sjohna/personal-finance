package repo

type Account struct {
	Id          int64  `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
}

type CreateAccountParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CreateAccount(dao DAO, id int64, params CreateAccountParams) (*Account, error) {
	log := repoFunctionLogger(dao.Logger(), "CreateAccount")
	defer logRepoReturn(log)

	SQL := `--sql
		INSERT INTO account (id, name, description)
		VALUES ($1, $2, $3)
		RETURNING *`

	var createdAccount Account
	err := dao.Get(&createdAccount, SQL, id, params.Name, params.Description)
	if err != nil {
		log.WithError(err).Error()
	}

	return &createdAccount, err
}

func GetAccount(dao DAO, id int64) (*Account, error) {
	log := repoFunctionLogger(dao.Logger(), "GetAccount")
	defer logRepoReturn(log)

	SQL := `--sql
		SELECT * FROM account
		WHERE account.id = $1`

	var account Account
	err := dao.Get(&account, SQL, id)
	if err != nil {
		log.WithError(err).Error()
	}

	return &account, err
}

// TODO: pagination
func GetAccounts(dao DAO) ([]*Account, error) {
	log := repoFunctionLogger(dao.Logger(), "GetAccounts")
	defer logRepoReturn(log)

	SQL := `--sql
		SELECT * FROM account`

	accounts := make([]*Account, 0)
	err := dao.Select(&accounts, SQL)
	if err != nil {
		log.WithError(err).Error()
	}

	return accounts, err
}
