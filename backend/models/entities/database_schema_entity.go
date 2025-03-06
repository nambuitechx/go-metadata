package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// Database schema entity
type DatabaseSchemaEntity struct {
	ID					string				`db:"id" json:"id"`
	Name				string				`db:"name" json:"name"`
	Json				*DatabaseSchema		`db:"json" json:"json"`
	UpdatedAt			int64				`db:"updatedat" json:"updatedAt"`
	UpdatedBy			string				`db:"updatedby" json:"updatedBy"`
	Deleted				bool				`db:"deleted" json:"deleted"`
	FqnHash				string				`db:"fqnhash" json:"fqnHash"`
}

// Database schema
type DatabaseSchema struct {
	ID					string				`json:"id"`
	Name				string				`json:"name"`
	FullyQualifiedName	string				`json:"fullyQualifiedName"`
	
	DisplayName			string				`json:"displayName"`
	Description			string				`json:"description"`

	ServiceType			string				`json:"serviceType"`
	Service				*EntityReference	`json:"service"`
	Database			*EntityReference	`json:"database"`

	Deleted				bool				`json:"deleted"`
}

func (s DatabaseSchema) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *DatabaseSchema) Scan(value interface{}) error {
	val, ok := value.([]byte)

	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(val, &s)
}

func (s *DatabaseSchema) ToEntityReference() *EntityReference {
	entityRef := &EntityReference{
		ID: s.ID,
		Type: "databaseSchema",
		Name: s.Name,
		FullyQualifiedName: s.FullyQualifiedName,
		DisplayName: s.DisplayName,
		Description: s.Description,
		Deleted: s.Deleted,
	}

	return entityRef
}

// APIs
type GetDatabaseSchemaEntitiesQuery struct {
	Limit int	`form:"limit"`
	Offset int	`form:"offset"`
}

type GetDatabaseSchemaEntityParam struct {
	ID string	`uri:"id" binding:"required"`
}

type CreateDatabaseSchemaEntityPayload struct {
	Name			string				`json:"name" binding:"required"`
	DisplayName		string				`json:"displayName"`
	Description		string				`json:"description"`

	Database		string				`json:"database" binding:"required"`
}
