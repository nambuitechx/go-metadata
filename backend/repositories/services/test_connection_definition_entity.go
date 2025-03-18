package repositories

import (
	"github.com/jmoiron/sqlx"
	servicesModels "github.com/nambuitechx/go-metadata/models/services"
)

type TestConnectionDefinitionEntityRepository struct {
	DB *sqlx.DB
}

func NewTestConnectionDefinitionEntityRepository(db *sqlx.DB) *TestConnectionDefinitionEntityRepository {
	return &TestConnectionDefinitionEntityRepository{ DB: db }
}

func (r *TestConnectionDefinitionEntityRepository) SelectTestConnectionDefinitionEntities(limit int, offset int) ([]servicesModels.TestConnectionDefinitionEntity, error) {
	testConnectionDefinitionEntities := []servicesModels.TestConnectionDefinitionEntity{}
	var err error
	
	if limit < 0 {
		statement := "SELECT * FROM test_connection_definition"
		err = r.DB.Select(&testConnectionDefinitionEntities, statement)
	} else {
		statement := "SELECT * FROM test_connection_definition LIMIT $1 OFFSET $2"
		err = r.DB.Select(&testConnectionDefinitionEntities, statement, limit, offset)
	}

	return testConnectionDefinitionEntities, err
}

func (r *TestConnectionDefinitionEntityRepository) SelectTestConnectionDefinitionEntityById(id string) (*servicesModels.TestConnectionDefinitionEntity, error) {
	testConnectionDefinitionEntity := &servicesModels.TestConnectionDefinitionEntity{}
	statement := "SELECT * FROM test_connection_definition WHERE id = $1"
	err := r.DB.Get(testConnectionDefinitionEntity, statement, id)
	return testConnectionDefinitionEntity, err
}

func (r *TestConnectionDefinitionEntityRepository) SelectTestConnectionDefinitionEntityByFqn(fqn string) (*servicesModels.TestConnectionDefinitionEntity, error) {
	testConnectionDefinitionEntity := &servicesModels.TestConnectionDefinitionEntity{}
	statement := "SELECT * FROM test_connection_definition WHERE json->>'fullyQualifiedName' = $1"
	err := r.DB.Get(testConnectionDefinitionEntity, statement, fqn)
	return testConnectionDefinitionEntity, err
}

func (r *TestConnectionDefinitionEntityRepository) InsertTestConnectionDefinitionEntity(payload *servicesModels.TestConnectionDefinitionEntity) (*servicesModels.TestConnectionDefinitionEntity, error) {
	var testConnectionDefinitionEntity = servicesModels.TestConnectionDefinitionEntity{}
	statement := `
		INSERT INTO test_connection_definition(id, name, fullyqualifiedname, json, updatedat, updatedby, deleted, namehash)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *
	`
	err := r.DB.Get(
		&testConnectionDefinitionEntity,
		statement,
		payload.ID,
		payload.Name,
		payload.FullyQualifiedName,
		payload.Json,
		payload.UpdatedAt,
		payload.UpdatedBy,
		payload.Deleted,
		payload.NameHash,
	)
	return &testConnectionDefinitionEntity, err
}

func (r *TestConnectionDefinitionEntityRepository) UpdateTestConnectionDefinitionEntity(payload *servicesModels.TestConnectionDefinitionEntity) (*servicesModels.TestConnectionDefinitionEntity, error) {
	var testConnectionDefinitionEntity = servicesModels.TestConnectionDefinitionEntity{}
	statement := `
		UPDATE test_connection_definition
		SET name = $2, fullyqualifiedname = $3, json = $4, updatedat = $5, updatedby = $6, deleted = $7, namehash = $8
		WHERE id = $1 RETURNING *
	`
	err := r.DB.Get(
		&testConnectionDefinitionEntity,
		statement,
		payload.ID,
		payload.Name,
		payload.FullyQualifiedName,
		payload.Json,
		payload.UpdatedAt,
		payload.UpdatedBy,
		payload.Deleted,
		payload.NameHash,
	)
	return &testConnectionDefinitionEntity, err
}

func (r *TestConnectionDefinitionEntityRepository) DeleteTestConnectionDefinitionEntityById(id string) error {
	statement := "DELETE FROM test_connection_definition WHERE id = $1"
	_, err := r.DB.Exec(statement, id)
	return err
}

func (r *TestConnectionDefinitionEntityRepository) DeleteTestConnectionDefinitionEntityByFqn(fqn string) error {
	statement := "DELETE FROM test_connection_definition WHERE json->>'fullyQualifiedName' = $1"
	_, err := r.DB.Exec(statement, fqn)
	return err
}
