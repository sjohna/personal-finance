package repo

type Account struct {
	Id          int64  `db:"id" json:"id"`
	AccountName string `db:"name" json:"accountName"`
	AccountDesc string `db:"description" json:"accountDesc"`
}

type CreateAccountParams struct {
	Id          int64  `json:"id"`
	AccountName string `json:"accountName"`
	AccountDesc string `json:"accountDesc"`
}

func CreateAccount(dao DAO, params CreateAccountParams) (*Account, error) {
	log := repoFunctionLogger(dao.Logger(), "CreateAccount")
	defer logRepoReturn(log)

	SQL := `--sql
		INSERT INTO account (id, name, description)
		VALUES ($1, $2, $3)
		RETURNING *`

	var createdAccount Account
	err := dao.Get(&createdAccount, SQL, params.Id, params.AccountName, params.AccountDesc)
	if err != nil {
		log.WithError(err).Error()
	}

	return &createdAccount, err
}

func GetAccount(dao DAO, accountID int64) (*Account, error) {
	log := repoFunctionLogger(dao.Logger(), "GetAccount")
	defer logRepoReturn(log)

	SQL := `--sql
		SELECT * FROM account
		WHERE account.id = $1`

	var account Account
	err := dao.Get(&account, SQL, accountID)
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
