package repositories

import (
	"github.com/jmoiron/sqlx"
	automationsModels "github.com/nambuitechx/go-metadata/models/automations"
)

type WorkflowEntityRepository struct {
	DB *sqlx.DB
}

func NewWorkflowEntityRepository(db *sqlx.DB) *WorkflowEntityRepository {
	return &WorkflowEntityRepository{ DB: db }
}

func (r *WorkflowEntityRepository) SelectWorkflowEntities(limit int, offset int) ([]automationsModels.WorkflowEntity, error) {
	workflowEntities := []automationsModels.WorkflowEntity{}
	var err error
	
	if limit < 0 {
		statement := "SELECT * FROM automations_workflow"
		err = r.DB.Select(&workflowEntities, statement)
	} else {
		statement := "SELECT * FROM automations_workflow LIMIT $1 OFFSET $2"
		err = r.DB.Select(&workflowEntities, statement, limit, offset)
	}

	return workflowEntities, err
}

func (r *WorkflowEntityRepository) SelectWorkflowEntityById(id string) (*automationsModels.WorkflowEntity, error) {
	workflowEntity := &automationsModels.WorkflowEntity{}
	statement := "SELECT * FROM automations_workflow WHERE id = $1"
	err := r.DB.Get(workflowEntity, statement, id)
	return workflowEntity, err
}

func (r *WorkflowEntityRepository) SelectWorkflowEntityByFqn(fqn string) (*automationsModels.WorkflowEntity, error) {
	workflowEntity := &automationsModels.WorkflowEntity{}
	statement := "SELECT * FROM automations_workflow WHERE json->>'fullyQualifiedName' = $1"
	err := r.DB.Get(workflowEntity, statement, fqn)
	return workflowEntity, err
}

func (r *WorkflowEntityRepository) InsertWorkflowEntity(payload *automationsModels.WorkflowEntity) (*automationsModels.WorkflowEntity, error) {
	var workflowEntity = automationsModels.WorkflowEntity{}
	statement := `
		INSERT INTO automations_workflow(id, name, workflowtype, status, json, updatedat, updatedby, deleted, namehash)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *
	`
	err := r.DB.Get(
		&workflowEntity,
		statement,
		payload.ID,
		payload.Name,
		payload.WorkflowType,
		payload.Status,
		payload.Json,
		payload.UpdatedAt,
		payload.UpdatedBy,
		payload.Deleted,
		payload.NameHash,
	)
	return &workflowEntity, err
}

func (r *WorkflowEntityRepository) UpdateWorkflowEntity(payload *automationsModels.WorkflowEntity) (*automationsModels.WorkflowEntity, error) {
	var workflowEntity = automationsModels.WorkflowEntity{}
	statement := `
		UPDATE automations_workflow
		SET name = $2, workflowtype = $3, status = $4, json = $5, updatedat = $6, updatedby = $7, deleted = $8, namehash = $9
		WHERE id = $1 RETURNING *
	`
	err := r.DB.Get(
		&workflowEntity,
		statement,
		payload.ID,
		payload.Name,
		payload.WorkflowType,
		payload.Status,
		payload.Json,
		payload.UpdatedAt,
		payload.UpdatedBy,
		payload.Deleted,
		payload.NameHash,
	)
	return &workflowEntity, err
}

func (r *WorkflowEntityRepository) DeleteWorkflowEntityById(id string) error {
	statement := "DELETE FROM automations_workflow WHERE id = $1"
	_, err := r.DB.Exec(statement, id)
	return err
}

func (r *WorkflowEntityRepository) DeleteWorkflowEntityByFqn(fqn string) error {
	statement := "DELETE FROM automations_workflow WHERE json->>'fullyQualifiedName' = $1"
	_, err := r.DB.Exec(statement, fqn)
	return err
}
