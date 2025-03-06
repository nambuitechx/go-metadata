package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/nambuitechx/go-metadata/models/entities"
)

type DBServiceEntityRepository struct {
	DB *sqlx.DB
}

func NewDBServiceEntityRepository(db *sqlx.DB) *DBServiceEntityRepository {
	return &DBServiceEntityRepository{ DB: db }
}

func (r *DBServiceEntityRepository) SelectDBServiceEntities(limit int, offset int) ([]models.DBServiceEntity, error) {
	dbserviceEntities := []models.DBServiceEntity{}
	var err error
	
	if limit < 0 {
		statement := "SELECT * FROM dbservice_entity"
		err = r.DB.Select(&dbserviceEntities, statement)
	} else {
		statement := "SELECT * FROM dbservice_entity LIMIT $1 OFFSET $2"
		err = r.DB.Select(&dbserviceEntities, statement, limit, offset)
	}

	return dbserviceEntities, err
}

func (r *DBServiceEntityRepository) SelectDBServiceEntityById(id string) (*models.DBServiceEntity, error) {
	dbserviceEntity := &models.DBServiceEntity{}
	statement := "SELECT * FROM dbservice_entity WHERE id = $1"
	err := r.DB.Get(dbserviceEntity, statement, id)
	return dbserviceEntity, err
}

func (r *DBServiceEntityRepository) SelectDBServiceEntityByFqn(fqn string) (*models.DBServiceEntity, error) {
	dbserviceEntity := &models.DBServiceEntity{}
	statement := "SELECT * FROM dbservice_entity WHERE json->>'fullyQualifiedName' = $1"
	err := r.DB.Get(dbserviceEntity, statement, fqn)
	return dbserviceEntity, err
}

func (r *DBServiceEntityRepository) InsertDBServiceEntity(payload *models.DBServiceEntity) (*models.DBServiceEntity, error) {
	var dbserviceEntity = models.DBServiceEntity{}
	statement := `
		INSERT INTO dbservice_entity(id, name, servicetype, json, updatedat, updatedby, deleted, namehash)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *
	`
	err := r.DB.Get(
		&dbserviceEntity,
		statement,
		payload.ID,
		payload.Name,
		payload.ServiceType,
		payload.Json,
		payload.UpdatedAt,
		payload.UpdatedBy,
		payload.Deleted,
		payload.NameHash,
	)
	return &dbserviceEntity, err
}

func (r *DBServiceEntityRepository) DeleteDBServiceEntityById(id string) error {
	statement := "DELETE FROM dbservice_entity WHERE id = $1"
	_, err := r.DB.Exec(statement, id)
	return err
}

func (r *DBServiceEntityRepository) DeleteDBServiceEntityByFqn(fqn string) error {
	statement := "DELETE FROM dbservice_entity WHERE json->>'fullyQualifiedName' = $1"
	_, err := r.DB.Exec(statement, fqn)
	return err
}
