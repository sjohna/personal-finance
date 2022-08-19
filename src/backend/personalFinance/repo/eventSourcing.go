package repo

import (
	"encoding/json"
	"time"

	"gopkg.in/guregu/null.v3"
)

func GetNextEntityId(dao DAO) (int64, error) {
	log := repoFunctionLogger(dao.Logger(), "GetNextEntityId")
	defer logRepoReturn(log)

	// language=SQL
	SQL := `
		select nextEntityId();`

	var nextId int64
	err := dao.Get(&nextId, SQL)
	if err != nil {
		log.WithError(err).Error()
	}

	return nextId, err
}

type Action struct {
	Id     int64       `db:"id"`
	Time   time.Time   `db:"time"`
	Origin string      `db:"origin"`
	Notes  null.String `db:"notes"`
}

// todo: make origin an enum, or have different functions
func CreateAction(dao DAO, actionOrigin string) (Action, error) {
	log := repoFunctionLogger(dao.Logger(), "CreateAction")
	defer logRepoReturn(log)

	// language=SQL
	SQL := `
		insert into action(origin)
		values($1)
		returning *;`

	var action Action
	err := dao.Get(&action, SQL, actionOrigin)
	if err != nil {
		log.WithError(err).Error()
	}

	return action, err
}

// todo: enums or functions
func CreateEvent(dao DAO, actionId int64, eventType string, entityType string, entityId int64, params interface{}) error {
	log := repoFunctionLogger(dao.Logger(), "CreateEvent")
	defer logRepoReturn(log)

	bytes, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Error("Failed to marshal params to JSON")
		return err
	}

	jsonString := string(bytes)

	// language=SQL
	SQL := `--sql
	insert into event(event_type, entity_type, entity_id, parameters, action_id)
	values($1, $2, $3, $4, $5);`

	// todo: check result?
	_, err = dao.Exec(SQL, eventType, entityType, entityId, jsonString, actionId)
	if err != nil {
		log.WithError(err).Error()
		return err
	}

	return nil
}

func HandleCreateSingleEntityFromApiCall(dao DAO, eventType string, entityType string, params interface{}) (int64, time.Time, error) {
	log := dao.Logger()

	entityId, err := GetNextEntityId(dao)
	if err != nil {
		log.WithError(err).Error("Error getting next entity ID")
		return 0, time.Time{}, err
	}

	action, err := CreateAction(dao, "api-call")
	if err != nil {
		log.WithError(err).Error("Error creating action")
		return 0, time.Time{}, err
	}

	err = CreateEvent(dao, action.Id, eventType, entityType, entityId, params)
	if err != nil {
		log.WithError(err).Error("Error creating event")
		return 0, time.Time{}, err
	}

	return entityId, action.Time, nil
}
