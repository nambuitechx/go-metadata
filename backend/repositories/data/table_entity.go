package repositories

import (
	"github.com/jmoiron/sqlx"
	baseModels "github.com/nambuitechx/go-metadata/models/base"
	dataModels "github.com/nambuitechx/go-metadata/models/data"
)

type TableEntityRepository struct {
	DB *sqlx.DB
}

func NewTableEntityRepository(db *sqlx.DB) *TableEntityRepository {
	return &TableEntityRepository{ DB: db }
}

func (r *TableEntityRepository) SelectTableEntities(limit int, offset int) ([]dataModels.TableEntity, error) {
	tableEntities := []dataModels.TableEntity{}
	var err error
	
	if limit < 0 {
		statement := "SELECT * FROM table_entity"
		err = r.DB.Select(&tableEntities, statement)
	} else {
		statement := "SELECT * FROM table_entity LIMIT $1 OFFSET $2"
		err = r.DB.Select(&tableEntities, statement, limit, offset)
	}

	return tableEntities, err
}

func (r *TableEntityRepository) SelectCountTableEntities() (*baseModels.EntityTotal, error) {
	entityTotal := &baseModels.EntityTotal{}
	statement := "SELECT COUNT(id) as total FROM table_entity"
	err := r.DB.Get(entityTotal, statement)
	return entityTotal, err
}

func (r *TableEntityRepository) SelectTableEntityById(id string) (*dataModels.TableEntity, error) {
	tableEntity := &dataModels.TableEntity{}
	statement := "SELECT * FROM table_entity WHERE id = $1"
	err := r.DB.Get(tableEntity, statement, id)
	return tableEntity, err
}

func (r *TableEntityRepository) SelectTableEntityByFqn(fqn string) (*dataModels.TableEntity, error) {
	tableEntity := &dataModels.TableEntity{}
	statement := "SELECT * FROM table_entity WHERE json->>'fullyQualifiedName' = $1"
	err := r.DB.Get(tableEntity, statement, fqn)
	return tableEntity, err
}

func (r *TableEntityRepository) InsertTableEntity(payload *dataModels.TableEntity) (*dataModels.TableEntity, error) {
	var tableEntity = dataModels.TableEntity{}
	statement := `
		INSERT INTO table_entity(id, name, json, updatedat, updatedby, deleted, fqnhash)
		VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING *
	`
	err := r.DB.Get(
		&tableEntity,
		statement,
		payload.ID,
		payload.Name,
		payload.Json,
		payload.UpdatedAt,
		payload.UpdatedBy,
		payload.Deleted,
		payload.FqnHash,
	)
	return &tableEntity, err
}

func (r *TableEntityRepository) UpdateTableEntity(payload *dataModels.TableEntity) (*dataModels.TableEntity, error) {
	var tableEntity = dataModels.TableEntity{}
	statement := `
		UPDATE table_entity
		SET name = $2, json = $3, updatedat = $4, updatedby = $5, deleted = $6, fqnhash = $7
		WHERE id = $1 RETURNING *
	`
	err := r.DB.Get(
		&tableEntity,
		statement,
		payload.ID,
		payload.Name,
		payload.Json,
		payload.UpdatedAt,
		payload.UpdatedBy,
		payload.Deleted,
		payload.FqnHash,
	)
	return &tableEntity, err
}

func (r *TableEntityRepository) DeleteTableEntityById(id string) error {
	statement := "DELETE FROM table_entity WHERE id = $1"
	_, err := r.DB.Exec(statement, id)
	return err
}

func (r *TableEntityRepository) DeleteTableEntityByFqn(fqn string) error {
	statement := "DELETE FROM table_entity WHERE json->>'fullyQualifiedName' = $1"
	_, err := r.DB.Exec(statement, fqn)
	return err
}
