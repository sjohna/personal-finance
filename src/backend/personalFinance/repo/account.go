package repo

import (
	"time"

	sq "github.com/elgris/sqrl"
)

type Account struct {
	Id          int64     `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt   time.Time `db:"updated_at" json:"updatedAt"`
}

type CreateAccountParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
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

func CreateAccount(dao DAO, id int64, params CreateAccountParams) (*Account, error) {
	log := repoFunctionLogger(dao.Logger(), "CreateAccount")
	defer logRepoReturn(log)

	SQL := `--sql
		insert into account (id, name, description)
		values ($1, $2, $3)
		returning *`

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

	SQL, _, err := accountSelect().Where("account.id = $1").ToSql()
	if err != nil {
		log.WithError(err).Fatal("SQL error initializing query")
	}

	var account Account
	err = dao.Get(&account, SQL, id)
	if err != nil {
		log.WithError(err).Error()
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
	}

	accounts := make([]*Account, 0)
	err = dao.Select(&accounts, SQL)
	if err != nil {
		log.WithError(err).Error()
	}

	return accounts, err
}
