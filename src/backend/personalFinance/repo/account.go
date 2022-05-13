package repo

import "time"

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
		select account.*,
	         create_action.time as created_at,
					 update_action.time as updated_at
		from account
		join event create_event on create_event.entity_id = account.id and create_event.event_type = 'create'
		join action create_action on create_event.action_id = create_action.id
		join lateral
		(
		  select action.time 
			from event
			join action on event.action_id = action.id
			where event.entity_id = account.id
			order by time desc
			limit 1
		) update_action on true
		where account.id = $1`

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
		select account.*,
	         create_action.time as created_at,
					 update_action.time as updated_at
		from account
		join event create_event on create_event.entity_id = account.id and create_event.event_type = 'create'
		join action create_action on create_event.action_id = create_action.id
		join lateral
		(
		  select action.time 
			from event
			join action on event.action_id = action.id
			where event.entity_id = account.id
			order by time desc
			limit 1
		) update_action on true`

	accounts := make([]*Account, 0)
	err := dao.Select(&accounts, SQL)
	if err != nil {
		log.WithError(err).Error()
	}

	return accounts, err
}
