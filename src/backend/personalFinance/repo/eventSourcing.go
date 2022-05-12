package repo

import "encoding/json"

func GetNextEntityId(dao DAO) (int64, error) {
	log := repoFunctionLogger(dao.Logger(), "GetNextEntityId")
	defer logRepoReturn(log)

	SQL := `--sql
		select nextEntityId();`

	var nextId int64
	err := dao.Get(&nextId, SQL)
	if err != nil {
		log.WithError(err).Error()
	}

	return nextId, err
}

// todo: make origin an enum, or have different functions
func CreateAction(dao DAO, actionOrigin string) (int64, error) {
	log := repoFunctionLogger(dao.Logger(), "CreateAction")
	defer logRepoReturn(log)

	SQL := `--sql
		insert into action(origin)
		values($1)
		returning id;`

	var nextId int64
	err := dao.Get(&nextId, SQL, actionOrigin)
	if err != nil {
		log.WithError(err).Error()
	}

	return nextId, err
}

// todo: enums or functions
func CreateEvent(dao DAO, actionId int64, eventType string, entityType string, params interface{}) error {
	log := repoFunctionLogger(dao.Logger(), "CreateEvent")
	defer logRepoReturn(log)

	bytes, err := json.Marshal(params)
	if err != nil {
		log.WithError(err).Error("Failed to marshal params to JSON")
		return err
	}

	jsonString := string(bytes)

	SQL := `--sql
	insert into event(event_type, entity_type, parameters, action_id)
	values($1, $2, $3, $4);`

	// todo: check result?
	_, err = dao.Exec(SQL, eventType, entityType, jsonString, actionId)
	if err != nil {
		log.WithError(err).Error()
		return err
	}

	return nil
}

func HandleCreateSingleEntityFromApiCall(dao DAO, eventType string, entityType string, params interface{}) error {
	log := dao.Logger()

	action_id, err := CreateAction(dao, "api-call")
	if err != nil {
		log.WithError(err).Error("Error creating action")
		return err
	}

	err = CreateEvent(dao, action_id, eventType, entityType, params)
	if err != nil {
		log.WithError(err).Error("Error creating event")
		return err
	}

	return nil
}
