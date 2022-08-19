package repo

import (
	"time"

	sq "github.com/elgris/sqrl"
)

type Account struct {
	Id               int64         `db:"id" json:"id"`
	Name             string        `db:"name" json:"name"`
	Description      string        `db:"description" json:"description"`
	CreatedAt        time.Time     `db:"created_at" json:"createdAt"`
	UpdatedAt        time.Time     `db:"updated_at" json:"updatedAt"`
	DebitsAndCredits []DebitCredit `json:"debitsAndCredits,omitempty"`
}

type CreateAccountParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DebitCreditType int64

const (
	Debit  DebitCreditType = 0
	Credit DebitCreditType = 1
)

type DebitCredit struct {
	Type       DebitCreditType `db:"type" json:"type"`
	Id         int64           `db:"id" json:"id"`
	Amount     int64           `db:"amount" json:"amount"`
	CurrencyId int64           `db:"currency_id" json:"currencyId"`
	Time       time.Time       `db:"time" json:"time"`
	AccountId  int64           `db:"account_id" json:"accountId"`
	CreatedAt  time.Time       `db:"created_at" json:"createdAt"`
	UpdatedAt  time.Time       `db:"updated_at" json:"updatedAt"`
}

type CreateDebitCreditParams struct {
	Amount     int64     `json:"amount"`
	CurrencyId int64     `json:"currencyId"`
	Time       time.Time `json:"time"`
	AccountId  int64     `json:"accountId"`
}

func accountSelect() *sq.SelectBuilder {
	return sq.Select("account.*", "create_action.time as created_at", "update_action.time as updated_at").
		From("account").
		Join("event create_event on create_event.entity_id = account.id and create_event.event_type = 'create'").
		Join("action create_action on create_event.action_id = create_action.id").
		JoinClause(`join lateral
							(
								select action.time 
								from event
								join action on event.action_id = action.id
								where event.entity_id = account.id
								order by time desc
								limit 1
							) update_action on true`)
}

func debitCreditSelect() *sq.SelectBuilder {
	return sq.Select("debits_and_credits.*", "create_action.time as created_at", "update_action.time as updated_at").
		From("debits_and_credits").
		Join("event create_event on create_event.entity_id = debits_and_credits.id and create_event.event_type = 'create'").
		Join("action create_action on create_event.action_id = create_action.id").
		JoinClause(`join lateral
							(
								select action.time 
								from event
								join action on event.action_id = action.id
								where event.entity_id = debits_and_credits.id
								order by time desc
								limit 1
							) update_action on true`)
}

func CreateAccount(dao DAO, id int64, params CreateAccountParams) (*Account, error) {
	log := repoFunctionLogger(dao.Logger(), "CreateAccount")
	defer logRepoReturn(log)

	// language=SQL
	SQL := `
		insert into account (id, name, description)
		values ($1, $2, $3)
		returning *`

	var createdAccount Account
	err := dao.Get(&createdAccount, SQL, id, params.Name, params.Description)
	if err != nil {
		log.WithError(err).Error()
		return nil, err
	}

	return &createdAccount, err
}

func GetAccount(dao DAO, id int64) (*Account, error) {
	log := repoFunctionLogger(dao.Logger(), "GetAccount")
	defer logRepoReturn(log)

	SQL, _, err := accountSelect().Where("account.id = $1").ToSql()
	if err != nil {
		log.WithError(err).Fatal("SQL error initializing query")
		return nil, err
	}

	var account Account
	err = dao.Get(&account, SQL, id)
	if err != nil {
		log.WithError(err).Error()
		return nil, err
	}

	return &account, err
}

// TODO: pagination
func GetAccounts(dao DAO) ([]*Account, error) {
	log := repoFunctionLogger(dao.Logger(), "GetAccounts")
	defer logRepoReturn(log)

	SQL, _, err := accountSelect().ToSql()
	if err != nil {
		log.WithError(err).Fatal("SQL error initializing query")
		return nil, err
	}

	accounts := make([]*Account, 0)
	err = dao.Select(&accounts, SQL)
	if err != nil {
		log.WithError(err).Error()
		return nil, err
	}

	return accounts, err
}

func CreateDebit(dao DAO, id int64, params CreateDebitCreditParams) (*DebitCredit, error) {
	log := repoFunctionLogger(dao.Logger(), "CreateDebit")
	defer logRepoReturn(log)

	// language=SQL
	SQL := `
		insert into debit (id, amount, currency_id, time, account_id)
		values ($1, $2, $3, $4, $5)
		returning *, 0 as type`

	var createdDebit DebitCredit
	err := dao.Get(&createdDebit, SQL, id, params.Amount, params.CurrencyId, params.Time, params.AccountId)
	if err != nil {
		log.WithError(err).Error()
		return nil, err
	}

	return &createdDebit, err
}

func CreateCredit(dao DAO, id int64, params CreateDebitCreditParams) (*DebitCredit, error) {
	log := repoFunctionLogger(dao.Logger(), "CreateCredit")
	defer logRepoReturn(log)

	// language=SQL
	SQL := `
		insert into credit (id, amount, currency_id, time, account_id)
		values ($1, $2, $3, $4, $5)
		returning *, 1 as type`

	var createdCredit DebitCredit
	err := dao.Get(&createdCredit, SQL, id, params.Amount, params.CurrencyId, params.Time, params.AccountId)
	if err != nil {
		log.WithError(err).Error()
		return nil, err
	}

	return &createdCredit, err
}

func GetDebitsAndCreditsForAccount(dao DAO, accountId int64) ([]DebitCredit, error) {
	log := repoFunctionLogger(dao.Logger(), "GetDebitsAndCreditsForAccount")
	defer logRepoReturn(log)

	SQL, _, err := debitCreditSelect().Where("account_id = $1").OrderBy("time asc").ToSql()
	if err != nil {
		log.WithError(err).Fatal("SQL error initializing query")
	}

	debitsAndCredits := make([]DebitCredit, 0)
	err = dao.Select(&debitsAndCredits, SQL, accountId)
	if err != nil {
		log.WithError(err).Error()
		return nil, err
	}

	return debitsAndCredits, err
}
