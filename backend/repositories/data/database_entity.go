package repositories

import (
	"github.com/jmoiron/sqlx"
	baseModels "github.com/nambuitechx/go-metadata/models/base"
	dataModels "github.com/nambuitechx/go-metadata/models/data"
)

type DatabaseEntityRepository struct {
	DB *sqlx.DB
}

func NewDatabaseEntityRepository(db *sqlx.DB) *DatabaseEntityRepository {
	return &DatabaseEntityRepository{ DB: db }
}

func (r *DatabaseEntityRepository) SelectDatabaseEntities(service string, limit int, offset int) ([]dataModels.DatabaseEntity, error) {
	databaseEntities := []dataModels.DatabaseEntity{}
	var err error
	
	if limit < 0 {
		statement := "SELECT * FROM database_entity WHERE json->>'fullyQualifiedName' LIKE ($1 || '.%')"
		err = r.DB.Select(&databaseEntities, statement, service)
	} else {
		statement := "SELECT * FROM database_entity WHERE json->>'fullyQualifiedName' LIKE ($1 || '.%') LIMIT $2 OFFSET $3"
		err = r.DB.Select(&databaseEntities, statement, service, limit, offset)
	}

	return databaseEntities, err
}

func (r *DatabaseEntityRepository) SelectCountDatabaseEntities(service string) (*baseModels.EntityTotal, error) {
	entityTotal := &baseModels.EntityTotal{}
	statement := "SELECT COUNT(id) as total FROM database_entity WHERE json->>'fullyQualifiedName' LIKE ($1 || '.%')"
	err := r.DB.Get(entityTotal, statement, service)
	return entityTotal, err
}

func (r *DatabaseEntityRepository) SelectDatabaseEntityById(id string) (*dataModels.DatabaseEntity, error) {
	databaseEntity := &dataModels.DatabaseEntity{}
	statement := "SELECT * FROM database_entity WHERE id = $1"
	err := r.DB.Get(databaseEntity, statement, id)
	return databaseEntity, err
}

func (r *DatabaseEntityRepository) SelectDatabaseEntityByFqn(fqn string) (*dataModels.DatabaseEntity, error) {
	databaseEntity := &dataModels.DatabaseEntity{}
	statement := "SELECT * FROM database_entity WHERE json->>'fullyQualifiedName' = $1"
	err := r.DB.Get(databaseEntity, statement, fqn)
	return databaseEntity, err
}

func (r *DatabaseEntityRepository) InsertDatabaseEntity(payload *dataModels.DatabaseEntity) (*dataModels.DatabaseEntity, error) {
	var databaseEntity = dataModels.DatabaseEntity{}
	statement := `
		INSERT INTO database_entity(id, name, json, updatedat, updatedby, deleted, fqnhash)
		VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING *
	`
	err := r.DB.Get(
		&databaseEntity,
		statement,
		payload.ID,
		payload.Name,
		payload.Json,
		payload.UpdatedAt,
		payload.UpdatedBy,
		payload.Deleted,
		payload.FqnHash,
	)
	return &databaseEntity, err
}

func (r *DatabaseEntityRepository) UpdateDatabaseEntity(payload *dataModels.DatabaseEntity) (*dataModels.DatabaseEntity, error) {
	var databaseEntity = dataModels.DatabaseEntity{}
	statement := `
		UPDATE database_entity
		SET name = $2, json = $3, updatedat = $4, updatedby = $5, deleted = $6, fqnhash = $7
		WHERE id = $1 RETURNING *
	`
	err := r.DB.Get(
		&databaseEntity,
		statement,
		payload.ID,
		payload.Name,
		payload.Json,
		payload.UpdatedAt,
		payload.UpdatedBy,
		payload.Deleted,
		payload.FqnHash,
	)
	return &databaseEntity, err
}

func (r *DatabaseEntityRepository) DeleteDatabaseEntityById(id string) error {
	statement := "DELETE FROM database_entity WHERE id = $1"
	_, err := r.DB.Exec(statement, id)
	return err
}

func (r *DatabaseEntityRepository) DeleteDatabaseEntityByFqn(fqn string) error {
	statement := "DELETE FROM database_entity WHERE json->>'fullyQualifiedName' = $1"
	_, err := r.DB.Exec(statement, fqn)
	return err
}
