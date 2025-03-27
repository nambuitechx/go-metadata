package repositories

import (
	"github.com/jmoiron/sqlx"
	baseModels "github.com/nambuitechx/go-metadata/models/base"
	dataModels "github.com/nambuitechx/go-metadata/models/data"
)

type StoredProcedureEntityRepository struct {
	DB *sqlx.DB
}

func NewStoredProcedureEntityRepository(db *sqlx.DB) *StoredProcedureEntityRepository {
	return &StoredProcedureEntityRepository{ DB: db }
}

func (r *StoredProcedureEntityRepository) SelectStoredProcedureEntities(limit int, offset int) ([]dataModels.StoredProcedureEntity, error) {
	storedProcedureEntities := []dataModels.StoredProcedureEntity{}
	var err error
	
	if limit < 0 {
		statement := "SELECT * FROM stored_procedure_entity"
		err = r.DB.Select(&storedProcedureEntities, statement)
	} else {
		statement := "SELECT * FROM stored_procedure_entity LIMIT $1 OFFSET $2"
		err = r.DB.Select(&storedProcedureEntities, statement, limit, offset)
	}

	return storedProcedureEntities, err
}

func (r *StoredProcedureEntityRepository) SelectCountStoredProcedureEntities() (*baseModels.EntityTotal, error) {
	entityTotal := &baseModels.EntityTotal{}
	statement := "SELECT COUNT(id) as total FROM stored_procedure_entity"
	err := r.DB.Get(entityTotal, statement)
	return entityTotal, err
}

func (r *StoredProcedureEntityRepository) SelectStoredProcedureEntityById(id string) (*dataModels.StoredProcedureEntity, error) {
	storedProcedureEntity := &dataModels.StoredProcedureEntity{}
	statement := "SELECT * FROM stored_procedure_entity WHERE id = $1"
	err := r.DB.Get(storedProcedureEntity, statement, id)
	return storedProcedureEntity, err
}

func (r *StoredProcedureEntityRepository) SelectStoredProcedureEntityByFqn(fqn string) (*dataModels.StoredProcedureEntity, error) {
	storedProcedureEntity := &dataModels.StoredProcedureEntity{}
	statement := "SELECT * FROM stored_procedure_entity WHERE json->>'fullyQualifiedName' = $1"
	err := r.DB.Get(storedProcedureEntity, statement, fqn)
	return storedProcedureEntity, err
}

func (r *StoredProcedureEntityRepository) InsertStoredProcedureEntity(payload *dataModels.StoredProcedureEntity) (*dataModels.StoredProcedureEntity, error) {
	var storedProcedureEntity = dataModels.StoredProcedureEntity{}
	statement := `
		INSERT INTO stored_procedure_entity(id, name, json, updatedat, updatedby, deleted, fqnhash)
		VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING *
	`
	err := r.DB.Get(
		&storedProcedureEntity,
		statement,
		payload.ID,
		payload.Name,
		payload.Json,
		payload.UpdatedAt,
		payload.UpdatedBy,
		payload.Deleted,
		payload.FqnHash,
	)
	return &storedProcedureEntity, err
}

func (r *StoredProcedureEntityRepository) UpdateStoredProcedureEntity(payload *dataModels.StoredProcedureEntity) (*dataModels.StoredProcedureEntity, error) {
	var storedProcedureEntity = dataModels.StoredProcedureEntity{}
	statement := `
		UPDATE stored_procedure_entity
		SET name = $2, json = $3, updatedat = $4, updatedby = $5, deleted = $6, fqnhash = $7
		WHERE id = $1 RETURNING *
	`
	err := r.DB.Get(
		&storedProcedureEntity,
		statement,
		payload.ID,
		payload.Name,
		payload.Json,
		payload.UpdatedAt,
		payload.UpdatedBy,
		payload.Deleted,
		payload.FqnHash,
	)
	return &storedProcedureEntity, err
}

func (r *StoredProcedureEntityRepository) DeleteStoredProcedureEntityById(id string) error {
	statement := "DELETE FROM stored_procedure_entity WHERE id = $1"
	_, err := r.DB.Exec(statement, id)
	return err
}

func (r *StoredProcedureEntityRepository) DeleteStoredProcedureEntityByFqn(fqn string) error {
	statement := "DELETE FROM stored_procedure_entity WHERE json->>'fullyQualifiedName' = $1"
	_, err := r.DB.Exec(statement, fqn)
	return err
}
