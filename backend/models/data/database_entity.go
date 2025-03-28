package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	typeModels "github.com/nambuitechx/go-metadata/models/type"
)

// Database entity
type DatabaseEntity struct {
	ID					string				`db:"id" json:"id"`
	Name				string				`db:"name" json:"name"`
	Json				*Database			`db:"json" json:"json"`
	UpdatedAt			int64				`db:"updatedat" json:"updatedAt"`
	UpdatedBy			string				`db:"updatedby" json:"updatedBy"`
	Deleted				bool				`db:"deleted" json:"deleted"`
	FqnHash				string				`db:"fqnhash" json:"fqnHash"`
}

// Database
type Database struct {
	ID					string						`json:"id"`
	Name				string						`json:"name"`
	FullyQualifiedName	string						`json:"fullyQualifiedName"`
	
	DisplayName			string						`json:"displayName"`
	Description			string						`json:"description"`

	ServiceType			string						`json:"serviceType"`
	Service				*typeModels.EntityReference	`json:"service"`

	Deleted				bool						`json:"deleted"`
}

func (s Database) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *Database) Scan(value interface{}) error {
	val, ok := value.([]byte)

	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(val, &s)
}

func (s *Database) ToEntityReference() *typeModels.EntityReference {
	entityRef := &typeModels.EntityReference{
		ID: s.ID,
		Type: "database",
		Name: s.Name,
		FullyQualifiedName: s.FullyQualifiedName,
		DisplayName: s.DisplayName,
		Description: s.Description,
		Deleted: s.Deleted,
	}

	return entityRef
}

// APIs
type GetDatabaseEntitiesQuery struct {
	Service		string	`form:"service"`
	Limit 		int		`form:"limit"`
	Offset 		int		`form:"offset"`
}

type GetDatabaseEntityByIdParam struct {
	ID string	`uri:"id" binding:"required"`
}

type GetDatabaseEntityByFqnParam struct {
	FQN string	`uri:"fqn" binding:"required"`
}

type CreateDatabaseEntityPayload struct {
	Name			string				`json:"name" binding:"required"`
	DisplayName		string				`json:"displayName"`
	Description		string				`json:"description"`

	Service			string				`json:"service" binding:"required"`
}
