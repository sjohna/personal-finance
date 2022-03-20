package repo

type Account struct {
	Id          int    `db:"id" json:"id"`
	AccountName string `db:"account_name" json:"accountName"`
	AccountDesc string `db:"account_desc" json:"accountDesc"`
}

func CreateAccount(dao DAO, accountName string, accountDesc string) (*Account, error) {
	log := repoFunctionLogger(dao.Logger(), "CreateAccount")
	defer logRepoReturn(log)

	SQL := `--sql
		INSERT INTO account (account_name, account_desc)
		VALUES ($1, $2)
		RETURNING *`

	var createdAccount Account
	err := dao.Get(&createdAccount, SQL, accountName, accountDesc)
	if err != nil {
		log.WithError(err).Error()
	}

	return &createdAccount, err
}

func GetAccount(dao DAO, accountID int) (*Account, error) {
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
