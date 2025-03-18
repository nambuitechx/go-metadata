package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

const TestConnectionDefinitionString string = "testConnectionDefinition"

// Test connection definition entity
type TestConnectionDefinitionEntity struct {
	ID						string						`db:"id" json:"id"`
	Name					string						`db:"name" json:"name"`
	FullyQualifiedName		string						`db:"fullyqualifiedname" json:"fullyQualifiedName"`
	Json					*TestConnectionDefinition	`db:"json" json:"json"`
	UpdatedAt				int64						`db:"updatedat" json:"updatedAt"`
	UpdatedBy				string						`db:"updatedby" json:"updatedBy"`
	Deleted					bool						`db:"deleted" json:"deleted"`
	NameHash				string						`db:"namehash" json:"nameHash"`
}

// Test connection definition
type TestConnectionDefinition struct {
	ID						string					`json:"id"`
	Name					string					`json:"name"`
	FullyQualifiedName		string					`json:"fullyQualifiedName"`

	DisplayName				string					`json:"displayName"`
	Description				string					`json:"description"`

	Steps					[]*TestConnectionStep	`json:"steps"`

	Deleted					bool					`json:"deleted"`
}

func (s TestConnectionDefinition) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *TestConnectionDefinition) Scan(value interface{}) error {
	val, ok := value.([]byte)

	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(val, &s)
}

type TestConnectionStep struct {
	Name			string			`json:"name"`
	Description		string			`json:"description"`
	ErrorMessage	string			`json:"errorMessage"`
	Mandatory		bool			`json:"mandatory"`
	ShortCircuit	bool			`json:"shortCircuit"`
}

// APIs
type GetTestConnectionDefinitionEntitiesQuery struct {
	Limit int	`form:"limit"`
	Offset int	`form:"offset"`
}

type GetTestConnectionDefinitionEntityByIdParam struct {
	ID string	`uri:"id" binding:"required"`
}

type GetTestConnectionDefinitionEntityByFqnParam struct {
	FQN string	`uri:"fqn" binding:"required"`
}
