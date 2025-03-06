package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strings"
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
	ID					string				`json:"id"`
	Name				string				`json:"name"`
	FullyQualifiedName	string				`json:"fullyQualifiedName"`
	
	DisplayName			string				`json:"displayName"`
	Description			string				`json:"description"`

	ServiceType			string				`json:"serviceType"`			
	Connection			*DatabaseConnection	`json:"connection"`

	Deleted				bool				`json:"deleted"`
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

func (s *DBService) ToEntityReference() *EntityReference {
	entityRef := &EntityReference{
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
	Type				string						`json:"type"`
	Scheme				string						`json:"scheme"`
	Username			string						`json:"username"`
	AuthType			map[string]interface{}		`json:"authType"`
	HostPort			string						`json:"hostPort"`
	Database			string						`json:"database"`
	IngestAllDatabases	bool						`json:"ingestAllDatabases"`
}

var PostgresType = map[string]int {"Postgres": 0}
var PostgresScheme = map[string]int {"postgresql+psycopg2": 0, "pgspider+psycopg2": 1}

func (c *PostgresConnection) SelfValidate() error {
	if _, ok := PostgresType[c.Type]; !ok {
		return errors.New("invalid postgres type")
	}

	if _, ok := PostgresScheme[c.Scheme]; !ok {
		return errors.New("invalid postgres scheme")
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

	return nil
}

func ValidatePostgresConnection(conn interface{}) error {
	bytes, err := json.Marshal(conn)

	if err != nil {
		return errors.New("failed to marshal postgres connection config")
	}

	var c *PostgresConnection = &PostgresConnection{}

	if err := json.Unmarshal(bytes, c); err != nil {
		return errors.New("failed to unmarshal to PostgresConnection")
	}

	return c.SelfValidate()
}

// Mysql
type MysqlConnection struct {
	Type				string					`json:"type"`
	Scheme				string					`json:"scheme"`
	Username			string					`json:"username"`
	AuthType			map[string]interface{}	`json:"authType"`
	HostPort			string					`json:"hostPort"`
	DatabaseName		string					`json:"databaseName"`
	DatabaseSchema		string					`json:"databaseSchema"`
}

var MysqlType = map[string]int {"Mysql": 0}
var MysqlScheme = map[string]int {"mysql+pymysql": 0}

func (c *MysqlConnection) SelfValidate() error {
	if _, ok := MysqlType[c.Type]; !ok {
		return errors.New("invalid mysql type")
	}

	if _, ok := MysqlScheme[c.Scheme]; !ok {
		return errors.New("invalid mysql scheme")
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

	return nil
}

func ValidateMysqlConnection(conn interface{}) error {
	bytes, err := json.Marshal(conn)
	
	if err != nil {
		return errors.New("failed to marshal mysql connection config")
	}

	var c *MysqlConnection = &MysqlConnection{}

	if err := json.Unmarshal(bytes, c); err != nil {
		return errors.New("failed to unmarshal to MysqlConnection")
	}

	return c.SelfValidate()
}

// APIs
type GetDBServiceEntitiesQuery struct {
	Limit int	`form:"limit"`
	Offset int	`form:"offset"`
}

type GetDBServiceEntityParam struct {
	ID string	`uri:"id" binding:"required"`
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
		return ValidatePostgresConnection(payload.Connection.Config)
	} else if idx == 1 {
		return ValidateMysqlConnection(payload.Connection.Config)
	} else {
		return errors.New("unsuported service type")
	}
}
