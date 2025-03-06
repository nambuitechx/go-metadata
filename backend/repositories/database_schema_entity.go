package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/nambuitechx/go-metadata/models"
)

type DatabaseSchemaEntityRepository struct {
	DB *sqlx.DB
}

func NewDatabaseSchemaEntityRepository(db *sqlx.DB) *DatabaseSchemaEntityRepository {
	return &DatabaseSchemaEntityRepository{ DB: db }
}

func (r *DatabaseSchemaEntityRepository) SelectDatabaseSchemaEntities(limit int, offset int) ([]models.DatabaseSchemaEntity, error) {
	databaseSchemaEntities := []models.DatabaseSchemaEntity{}
	var err error
	
	if limit < 0 {
		statement := "SELECT * FROM database_schema_entity"
		err = r.DB.Select(&databaseSchemaEntities, statement)
	} else {
		statement := "SELECT * FROM database_schema_entity LIMIT $1 OFFSET $2"
		err = r.DB.Select(&databaseSchemaEntities, statement, limit, offset)
	}

	return databaseSchemaEntities, err
}

func (r *DatabaseSchemaEntityRepository) SelectDatabaseSchemaEntityById(id string) (*models.DatabaseSchemaEntity, error) {
	databaseSchemaEntity := &models.DatabaseSchemaEntity{}
	statement := "SELECT * FROM database_schema_entity WHERE id = $1"
	err := r.DB.Get(databaseSchemaEntity, statement, id)
	return databaseSchemaEntity, err
}

func (r *DatabaseSchemaEntityRepository) SelectDatabaseSchemaEntityByFqn(fqn string) (*models.DatabaseSchemaEntity, error) {
	databaseSchemaEntity := &models.DatabaseSchemaEntity{}
	statement := "SELECT * FROM database_schema_entity WHERE json->>'fullyQualifiedName' = $1"
	err := r.DB.Get(databaseSchemaEntity, statement, fqn)
	return databaseSchemaEntity, err
}

func (r *DatabaseSchemaEntityRepository) InsertDatabaseSchemaEntity(payload *models.DatabaseSchemaEntity) (*models.DatabaseSchemaEntity, error) {
	var databaseSchemaEntity = models.DatabaseSchemaEntity{}
	statement := `
		INSERT INTO database_schema_entity(id, name, json, updatedat, updatedby, deleted, fqnhash)
		VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING *
	`
	err := r.DB.Get(
		&databaseSchemaEntity,
		statement,
		payload.ID,
		payload.Name,
		payload.Json,
		payload.UpdatedAt,
		payload.UpdatedBy,
		payload.Deleted,
		payload.FqnHash,
	)
	return &databaseSchemaEntity, err
}

func (r *DatabaseSchemaEntityRepository) DeleteDatabaseSchemaEntityById(id string) error {
	statement := "DELETE FROM database_schema_entity WHERE id = $1"
	_, err := r.DB.Exec(statement, id)
	return err
}

func (r *DatabaseSchemaEntityRepository) DeleteDatabaseSchemaEntityByFqn(fqn string) error {
	statement := "DELETE FROM database_schema_entity WHERE json->>'fullyQualifiedName' = $1"
	_, err := r.DB.Exec(statement, fqn)
	return err
}
