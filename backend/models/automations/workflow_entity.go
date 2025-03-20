package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	servicesModels "github.com/nambuitechx/go-metadata/models/services"
)

// Workflow entity
type WorkflowEntity struct {
	ID					string				`db:"id" json:"id"`
	Name				string				`db:"name" json:"name"`
	WorkflowType		string				`db:"workflowtype" json:"workflowType"`
	Status				string				`db:"status" json:"status"`
	Json				*Workflow			`db:"json" json:"json"`
	UpdatedAt			int64				`db:"updatedat" json:"updatedAt"`
	UpdatedBy			string				`db:"updatedby" json:"updatedBy"`
	Deleted				bool				`db:"deleted" json:"deleted"`
	NameHash			string				`db:"namehash" json:"nameHash"`
}

// Workflow
type Workflow struct {
	ID						string									`json:"id"`
	Name					string									`json:"name"`
	FullyQualifiedName		string									`json:"fullyQualifiedName"`

	DisplayName				string									`json:"displayName"`
	Description				string									`json:"description"`

	WorkflowType			string									`json:"workflowType"`
	Status					string									`json:"status"`

	Request					*TestServiceConnection					`json:"request"`
	Response				*servicesModels.TestConnectionResult	`json:"response"`

	Deleted					bool									`json:"deleted"`
}

func (s Workflow) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *Workflow) Scan(value interface{}) error {
	val, ok := value.([]byte)

	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(val, &s)
}

// Type and status
var WorkflowType = map[string]int {"TEST_CONNECTION": 0}
var WorkflowStatus = map[string]int {"Pending": 0, "Successful": 1, "Failed": 2, "Running": 3}

// Test service connection request
type TestServiceConnection struct {
	ServiceType			string									`json:"serviceType"`	// Ex: Database, Dashboard, Messaging, etc.
	ServiceName			string									`json:"serviceName"` 

	ConnectionType		string									`json:"connectionType"`	// Ex: Postgres, MySQL, Snowflake, etc.
	Connection			*servicesModels.DatabaseConnection		`json:"connection"`
}

// APIs
type GetWorkflowEntitiesQuery struct {
	Limit int	`form:"limit"`
	Offset int	`form:"offset"`
}

type GetWorkflowEntityByIdParam struct {
	ID string	`uri:"id" binding:"required"`
}

type GetWorkflowEntityByFqnParam struct {
	FQN string	`uri:"fqn" binding:"required"`
}

type CreateWorkflowRequest struct {
	Name				string									`json:"name"`

	DisplayName			string									`json:"displayName"`
	Description			string									`json:"description"`

	WorkflowType		string									`json:"workflowType"`
	Status				string									`json:"status"`

	Request				*TestServiceConnection					`json:"request"`
	Response			*servicesModels.TestConnectionResult	`json:"response"`
}

type PatchWorkflowRequest struct {
	DisplayName			*string									`json:"displayName"`
	Description			*string									`json:"description"`
	Status				*string									`json:"status"`
	Request				*TestServiceConnection					`json:"request"`
	Response			*servicesModels.TestConnectionResult	`json:"response"`
}
