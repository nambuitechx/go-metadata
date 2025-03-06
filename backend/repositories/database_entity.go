package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/nambuitechx/go-metadata/models"
)

type DatabaseEntityRepository struct {
	DB *sqlx.DB
}

func NewDatabaseEntityRepository(db *sqlx.DB) *DatabaseEntityRepository {
	return &DatabaseEntityRepository{ DB: db }
}

func (r *DatabaseEntityRepository) SelectDatabaseEntities(limit int, offset int) ([]models.DatabaseEntity, error) {
	databaseEntities := []models.DatabaseEntity{}
	var err error
	
	if limit < 0 {
		statement := "SELECT * FROM database_entity"
		err = r.DB.Select(&databaseEntities, statement)
	} else {
		statement := "SELECT * FROM database_entity LIMIT $1 OFFSET $2"
		err = r.DB.Select(&databaseEntities, statement, limit, offset)
	}

	return databaseEntities, err
}

func (r *DatabaseEntityRepository) SelectDatabaseEntityById(id string) (*models.DatabaseEntity, error) {
	databaseEntity := &models.DatabaseEntity{}
	statement := "SELECT * FROM database_entity WHERE id = $1"
	err := r.DB.Get(databaseEntity, statement, id)
	return databaseEntity, err
}

func (r *DatabaseEntityRepository) SelectDatabaseEntityByFqn(fqn string) (*models.DatabaseEntity, error) {
	databaseEntity := &models.DatabaseEntity{}
	statement := "SELECT * FROM database_entity WHERE json->>'fullyQualifiedName' = $1"
	err := r.DB.Get(databaseEntity, statement, fqn)
	return databaseEntity, err
}

func (r *DatabaseEntityRepository) InsertDatabaseEntity(payload *models.DatabaseEntity) (*models.DatabaseEntity, error) {
	var databaseEntity = models.DatabaseEntity{}
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
