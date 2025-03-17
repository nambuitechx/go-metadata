package repositories

import (
	"github.com/jmoiron/sqlx"
	servicesModels "github.com/nambuitechx/go-metadata/models/services"
)

type DBServiceEntityRepository struct {
	DB *sqlx.DB
}

func NewDBServiceEntityRepository(db *sqlx.DB) *DBServiceEntityRepository {
	return &DBServiceEntityRepository{ DB: db }
}

func (r *DBServiceEntityRepository) SelectDBServiceEntities(limit int, offset int) ([]servicesModels.DBServiceEntity, error) {
	dbserviceEntities := []servicesModels.DBServiceEntity{}
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

func (r *DBServiceEntityRepository) SelectDBServiceEntityById(id string) (*servicesModels.DBServiceEntity, error) {
	dbserviceEntity := &servicesModels.DBServiceEntity{}
	statement := "SELECT * FROM dbservice_entity WHERE id = $1"
	err := r.DB.Get(dbserviceEntity, statement, id)
	return dbserviceEntity, err
}

func (r *DBServiceEntityRepository) SelectDBServiceEntityByFqn(fqn string) (*servicesModels.DBServiceEntity, error) {
	dbserviceEntity := &servicesModels.DBServiceEntity{}
	statement := "SELECT * FROM dbservice_entity WHERE json->>'fullyQualifiedName' = $1"
	err := r.DB.Get(dbserviceEntity, statement, fqn)
	return dbserviceEntity, err
}

func (r *DBServiceEntityRepository) InsertDBServiceEntity(payload *servicesModels.DBServiceEntity) (*servicesModels.DBServiceEntity, error) {
	var dbserviceEntity = servicesModels.DBServiceEntity{}
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

func (r *DBServiceEntityRepository) UpdateDBServiceEntity(payload *servicesModels.DBServiceEntity) (*servicesModels.DBServiceEntity, error) {
	var dbserviceEntity = servicesModels.DBServiceEntity{}
	statement := `
		UPDATE dbservice_entity
		SET name = $2, servicetype = $3, json = $4, updatedat = $5, updatedby = $6, deleted = $7, namehash = $8
		WHERE id = $1 RETURNING *
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
