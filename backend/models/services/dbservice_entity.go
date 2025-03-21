package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strings"

	baseUtils "github.com/nambuitechx/go-metadata/utils"
	securityModels "github.com/nambuitechx/go-metadata/models/security"
	typeModels "github.com/nambuitechx/go-metadata/models/type"
)

// Database service entity
type DBServiceEntity struct {
	ID					string				`db:"id" json:"id"`
	Name				string				`db:"name" json:"name"`
	ServiceType			string				`db:"servicetype" json:"serviceType"`
	Json				*DBService			`db:"json" json:"json"`
	UpdatedAt			int64				`db:"updatedat" json:"updatedAt"`
	UpdatedBy			string				`db:"updatedby" json:"updatedBy"`
	Deleted				bool				`db:"deleted" json:"deleted"`
	NameHash			string				`db:"namehash" json:"nameHash"`
}

// Database service
type DBService struct {
	ID						string					`json:"id"`
	Name					string					`json:"name"`
	FullyQualifiedName		string					`json:"fullyQualifiedName"`
	
	DisplayName				string					`json:"displayName"`
	Description				string					`json:"description"`

	ServiceType				string					`json:"serviceType"`			
	Connection				*DatabaseConnection		`json:"connection"`

	TestConnectionResult	*TestConnectionResult	`json:"testConnectionResult"`

	Deleted					bool					`json:"deleted"`
}

func (s DBService) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *DBService) Scan(value interface{}) error {
	val, ok := value.([]byte)

	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(val, &s)
}

func (s *DBService) ToEntityReference() *typeModels.EntityReference {
	entityRef := &typeModels.EntityReference{
		ID: s.ID,
		Type: "databaseService",
		Name: s.Name,
		FullyQualifiedName: s.FullyQualifiedName,
		DisplayName: s.DisplayName,
		Description: s.Description,
		Deleted: s.Deleted,
	}

	return entityRef
}

// Service type
var ServiceType = map[string]int {"Postgres": 0, "MySQL": 1}

func ValidateServiceType(serviceType string) (int, error) {
	idx, ok := ServiceType[serviceType]

	if !ok {
		return -1, errors.New("invalid service type")
	}

	return idx, nil
}

// All types of connections
type DatabaseConnection struct {
	Config				map[string]interface{}		`json:"config" binding:"required"`
}

// Postgres
type PostgresConnection struct {
	Type						*string						`json:"type"`
	Scheme						*string						`json:"scheme"`
	Username					string						`json:"username"`
	AuthType					map[string]interface{}		`json:"authType"`
	HostPort					string						`json:"hostPort"`
	Database					string						`json:"database"`
	IngestAllDatabases			*bool						`json:"ingestAllDatabases"`
	SSLMode						*string						`json:"sslMode"`
	SupportsMetadataExtraction	*bool						`json:"supportsMetadataExtraction"`
	SupportsUsageExtraction		*bool						`json:"supportsUsageExtraction"`
	SupportsLineageExtraction	*bool						`json:"supportsLineageExtraction"`
	SupportsDBTExtraction		*bool						`json:"supportsDBTExtraction"`
	SupportsProfiler			*bool						`json:"supportsProfiler"`
	SupportsDatabase			*bool						`json:"supportsDatabase"`
	SupportsQueryComment		*bool						`json:"supportsQueryComment"`
	SupportsDataDiff			*bool						`json:"supportsDataDiff"`
}

var PostgresType = map[string]int {"Postgres": 0}
var PostgresScheme = map[string]int {"postgresql+psycopg2": 0, "pgspider+psycopg2": 1}

func (c *PostgresConnection) SelfValidate() error {
	if c.Type == nil {
		v := "Postgres"
		c.Type = &v
	} else {
		if _, ok := PostgresType[*c.Type]; !ok {
			return errors.New("invalid postgres type")
		}
	}

	if c.Scheme == nil {
		v := "postgresql+psycopg2"
		c.Scheme = &v
	} else {
		if _, ok := PostgresScheme[*c.Scheme]; !ok {
			return errors.New("invalid postgres scheme")
		}
	}

	if strings.TrimSpace(c.Username) == "" {
		return errors.New("invalid postgres username")
	}

	if strings.TrimSpace(c.HostPort) == "" {
		return errors.New("invalid postgres hostPort")
	}

	if strings.TrimSpace(c.Database) == "" {
		return errors.New("invalid postgres database")
	}

	if c.IngestAllDatabases == nil {
		v := false
		c.IngestAllDatabases = &v
	}

	if c.SSLMode == nil {
		v := "disable"
		c.SSLMode = &v
	} else {
		if _, ok := securityModels.SSLMode[*c.SSLMode]; !ok {
			return errors.New("invalid ssl mode")
		}
	}

	if c.SupportsMetadataExtraction == nil {
		v := true
		c.SupportsMetadataExtraction = &v
	}

	if c.SupportsUsageExtraction == nil {
		v := true
		c.SupportsUsageExtraction = &v
	}

	if c.SupportsLineageExtraction == nil {
		v := true
		c.SupportsLineageExtraction = &v
	}

	if c.SupportsDBTExtraction == nil {
		v := true
		c.SupportsDBTExtraction = &v
	}

	if c.SupportsProfiler == nil {
		v := true
		c.SupportsProfiler = &v
	}

	if c.SupportsDatabase == nil {
		v := true
		c.SupportsDatabase = &v
	}

	if c.SupportsQueryComment == nil {
		v := true
		c.SupportsQueryComment = &v
	}

	if c.SupportsDataDiff == nil {
		v := true
		c.SupportsDataDiff = &v
	}

	return nil
}

func ValidatePostgresConnection(databaseConnection *DatabaseConnection) error {
	bytes, err := json.Marshal(databaseConnection.Config)

	if err != nil {
		return errors.New("failed to marshal postgres connection config")
	}

	var c *PostgresConnection = &PostgresConnection{}

	if err := json.Unmarshal(bytes, c); err != nil {
		return errors.New("failed to unmarshal to PostgresConnection")
	}

	validateErr := c.SelfValidate()

	if validateErr != nil {
		return validateErr
	}

	dataFromStruct, dataFromStructErr := baseUtils.StructToMap(c)

	if dataFromStructErr != nil {
		return dataFromStructErr
	}

	databaseConnection.Config = dataFromStruct
	return nil
}

// Mysql
type MysqlConnection struct {
	Type						*string					`json:"type"`
	Scheme						*string					`json:"scheme"`
	Username					string					`json:"username"`
	AuthType					map[string]interface{}	`json:"authType"`
	HostPort					string					`json:"hostPort"`
	DatabaseName				string					`json:"databaseName"`
	DatabaseSchema				string					`json:"databaseSchema"`
	SupportsMetadataExtraction	*bool					`json:"supportsMetadataExtraction"`
	SupportsDBTExtraction		*bool					`json:"supportsDBTExtraction"`
	SupportsProfiler			*bool					`json:"supportsProfiler"`
	SupportsQueryComment		*bool					`json:"supportsQueryComment"`
	SupportsDataDiff			*bool					`json:"supportsDataDiff"`
	SupportsUsageExtraction		*bool					`json:"supportsUsageExtraction"`
	SupportsLineageExtraction	*bool					`json:"supportsLineageExtraction"`
}

var MysqlType = map[string]int {"Mysql": 0}
var MysqlScheme = map[string]int {"mysql+pymysql": 0}

func (c *MysqlConnection) SelfValidate() error {
	if c.Type == nil {
		v := "Mysql"
		c.Type = &v
	} else {
		if _, ok := MysqlType[*c.Type]; !ok {
			return errors.New("invalid mysql type")
		}
	}

	if c.Scheme == nil {
		v := "mysql+pymysql"
		c.Scheme = &v
	} else {
		if _, ok := MysqlScheme[*c.Scheme]; !ok {
			return errors.New("invalid mysql scheme")
		}
	}

	if strings.TrimSpace(c.Username) == "" {
		return errors.New("invalid mysql username")
	}

	if strings.TrimSpace(c.HostPort) == "" {
		return errors.New("invalid mysql hostPort")
	}

	if strings.TrimSpace(c.DatabaseName) == "" {
		return errors.New("invalid mysql databaseName")
	}

	if strings.TrimSpace(c.DatabaseSchema) == "" {
		return errors.New("invalid mysql databaseSchema")
	}

	if c.SupportsMetadataExtraction == nil {
		v := true
		c.SupportsMetadataExtraction = &v
	}

	if c.SupportsDBTExtraction == nil {
		v := true
		c.SupportsDBTExtraction = &v
	}

	if c.SupportsProfiler == nil {
		v := true
		c.SupportsProfiler = &v
	}

	if c.SupportsQueryComment == nil {
		v := true
		c.SupportsQueryComment = &v
	}

	if c.SupportsDataDiff == nil {
		v := true
		c.SupportsDataDiff = &v
	}

	if c.SupportsUsageExtraction == nil {
		v := true
		c.SupportsUsageExtraction = &v
	}
	
	if c.SupportsLineageExtraction == nil {
		v := true
		c.SupportsLineageExtraction = &v
	}

	return nil
}

func ValidateMysqlConnection(databaseConnection *DatabaseConnection) error {
	bytes, err := json.Marshal(databaseConnection.Config)
	
	if err != nil {
		return errors.New("failed to marshal mysql connection config")
	}

	var c *MysqlConnection = &MysqlConnection{}

	if err := json.Unmarshal(bytes, c); err != nil {
		return errors.New("failed to unmarshal to MysqlConnection")
	}

	validateErr := c.SelfValidate()

	if validateErr != nil {
		return validateErr
	}

	dataFromStruct, dataFromStructErr := baseUtils.StructToMap(c)

	if dataFromStructErr != nil {
		return dataFromStructErr
	}

	databaseConnection.Config = dataFromStruct
	return nil
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
type GetDBServiceEntitiesQuery struct {
	Limit int	`form:"limit"`
	Offset int	`form:"offset"`
}

type GetDBServiceEntityByIdParam struct {
	ID string	`uri:"id" binding:"required"`
}

type GetDBServiceEntityByFqnParam struct {
	FQN string	`uri:"fqn" binding:"required"`
}

type CreateDBServiceEntityPayload struct {
	Name			string				`json:"name" binding:"required"`
	DisplayName		string				`json:"displayName"`
	Description		string				`json:"description"`

	ServiceType		string				`json:"serviceType" binding:"required"`
	Connection		*DatabaseConnection	`json:"connection" binding:"required"`
}

func ValidateCreateDBServiceEntityPayload(payload *CreateDBServiceEntityPayload) error {
	idx, serviceTypeErr := ValidateServiceType(payload.ServiceType)

	if serviceTypeErr != nil {
		return serviceTypeErr
	}

	if idx == 0 {
		return ValidatePostgresConnection(payload.Connection)
	} else if idx == 1 {
		return ValidateMysqlConnection(payload.Connection)
	} else {
		return errors.New("unsuported service type")
	}
}
