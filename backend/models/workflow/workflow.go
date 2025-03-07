package models

import (
	"github.com/nambuitechx/go-metadata/models/entities"
)

type Workflow struct {
	ID						string					`json:"id"`
	Name					string					`json:"name"`
	FullyQualifiedName		string					`json:"fullyQualifiedName"`

	DisplayName				string					`json:"displayName"`
	Description				string					`json:"description"`

	WorkflowType			string					`json:"workflowType"`
	Status					string					`json:"status"`

	Request					*TestServiceConnection	`json:"request"`
	Response				*TestConnectionResult	`json:"response"`

	Deleted					bool					`json:"deleted"`
}

// Type and status
var WorkflowType = map[string]int {"TEST_CONNECTION": 0}
var WorkflowStatus = map[string]int {"Pending": 0, "Successful": 1, "Failed": 2, "Running": 3}

// Test service connection request
type TestServiceConnection struct {
	ServiceType			string							`json:"serviceType"`
	ServiceName			string							`json:"serviceName"`

	ConnectionType		string							`json:"connectionType"`
	Connection			*models.DatabaseConnection		`json:"connection"`
}

// Test connection result
type TestConnectionResult struct {
	LastUpdatedAt		string							`json:"lastUpdatedAt"`
	Status				string							`json:"status"`
	Steps				[]*TestConnectionStepResult		`json:"steps"`
}

var StatusType = map[string]int {"Successful": 0, "Failed": 1, "Running": 2}

type TestConnectionStepResult struct {
	Name				string		`json:"name"`
	Mandatory			bool		`json:"mandatory"`
	Passed				bool		`json:"passed"`
	Message				string		`json:"message"`
	ErrorLog			string		`json:"errorLog"`
}

// APIs
type CreateWorkflowRequest struct {
	Name				string					`json:"name"`

	DisplayName			string					`json:"displayName"`
	Description			string					`json:"description"`

	WorkflowType		string					`json:"workflowType"`
	Status				string					`json:"status"`

	Request				*TestServiceConnection	`json:"request"`
	Response			*TestConnectionResult	`json:"response"`
}
