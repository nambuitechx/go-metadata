package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	typeModels "github.com/nambuitechx/go-metadata/models/type"
)

// Stored procedure entity
type StoredProcedureEntity struct {
	ID					string				`db:"id" json:"id"`
	Name				string				`db:"name" json:"name"`
	Json				*StoredProcedure	`db:"json" json:"json"`
	UpdatedAt			int64				`db:"updatedat" json:"updatedAt"`
	UpdatedBy			string				`db:"updatedby" json:"updatedBy"`
	Deleted				bool				`db:"deleted" json:"deleted"`
	FqnHash				string				`db:"fqnhash" json:"fqnHash"`
}

// StoredProcedure
type StoredProcedure struct {
	ID						string							`json:"id"`
	Name					string							`json:"name"`
	FullyQualifiedName		string							`json:"fullyQualifiedName"`
	
	DisplayName				string							`json:"displayName"`
	Description				string							`json:"description"`

	StoredProcedureCode		*StoredProcedureCode			`json:"storedProcedureCode"`
	StoredProcedureType		string							`json:"storedProcedureType"`

	ServiceType				string							`json:"serviceType"`
	Service					*typeModels.EntityReference		`json:"service"`
	Database				*typeModels.EntityReference		`json:"database"`
	DatabaseSchema			*typeModels.EntityReference		`json:"databaseSchema"`

	Deleted					bool							`json:"deleted"`
}

func (s StoredProcedure) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *StoredProcedure) Scan(value interface{}) error {
	val, ok := value.([]byte)

	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(val, &s)
}

type StoredProcedureCode struct {
	Language		string			`json:"language"`
	Code			string			`json:"code"`
}

// APIs
type GetStoredProcedureEntitiesQuery struct {
	Limit int	`form:"limit"`
	Offset int	`form:"offset"`
}

type GetStoredProcedureEntityByIdParam struct {
	ID string	`uri:"id" binding:"required"`
}

type GetStoredProcedureEntityByFqnParam struct {
	FQN string	`uri:"fqn" binding:"required"`
}

type CreateStoredProcedureEntityPayload struct {
	Name					string							`json:"name" binding:"required"`
	DisplayName				string							`json:"displayName"`
	Description				string							`json:"description"`

	StoredProcedureCode		*StoredProcedureCode			`json:"storedProcedureCode"`
	StoredProcedureType		string							`json:"storedProcedureType"`

	DatabaseSchema			string							`json:"databaseSchema" binding:"required"`
}
