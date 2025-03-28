package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	typeModels "github.com/nambuitechx/go-metadata/models/type"
	// baseUtils "github.com/nambuitechx/go-metadata/utils"
)

// Table entity
type TableEntity struct {
	ID					string				`db:"id" json:"id"`
	Name				string				`db:"name" json:"name"`
	Json				*Table				`db:"json" json:"json"`
	UpdatedAt			int64				`db:"updatedat" json:"updatedAt"`
	UpdatedBy			string				`db:"updatedby" json:"updatedBy"`
	Deleted				bool				`db:"deleted" json:"deleted"`
	FqnHash				string				`db:"fqnhash" json:"fqnHash"`
}

// Table
type Table struct {
	ID					string						`json:"id"`
	Name				string						`json:"name"`
	FullyQualifiedName	string						`json:"fullyQualifiedName"`
	
	DisplayName			string						`json:"displayName"`
	Description			string						`json:"description"`

	ServiceType			string						`json:"serviceType"`
	Service				*typeModels.EntityReference	`json:"service"`
	Database			*typeModels.EntityReference	`json:"database"`
	DatabaseSchema		*typeModels.EntityReference	`json:"databaseSchema"`

	TableType			string						`json:"tableType"`
	TableConstraints	[]TableConstraint			`json:"tableConstraints"`

	Columns				[]Column					`json:"columns"`

	Deleted				bool						`json:"deleted"`
}

func (s Table) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *Table) Scan(value interface{}) error {
	val, ok := value.([]byte)

	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(val, &s)
}

func (s *Table) ToEntityReference() *typeModels.EntityReference {
	entityRef := &typeModels.EntityReference{
		ID: s.ID,
		Type: "table",
		Name: s.Name,
		FullyQualifiedName: s.FullyQualifiedName,
		DisplayName: s.DisplayName,
		Description: s.Description,
		Deleted: s.Deleted,
	}

	return entityRef
}

// Table type
var TableType = map[string]int {
	"Regular": 0,
	"External": 1,
	"Dynamic": 2,
	"View": 3,
	"SecureView": 4,
	"MaterializedView": 5,
	"Iceberg": 6,
	"Local": 7,
	"Partitioned": 8,
	"Foreign": 9,
	"Transient": 10,
}

func ValidateTableType(tableType string) (int, error) {
	idx, ok := TableType[tableType]

	if !ok {
		return -1, errors.New("invalid table type")
	}

	return idx, nil
}

// Column
type Column struct {
	Name				*string		`json:"name"`
	FullyQualifiedName	*string		`json:"fullyQualifiedName"`

	DisplayName			string		`json:"displayName"`
	Description			string		`json:"description"`

	DataType			*string		`json:"dataType"`
	ArrayDataType		*string		`json:"arrayDataType"`
	DataLength			int32		`json:"dataLength"`
	DataTypeDisplay		string		`json:"dataTypeDisplay"`

	Precision			int32		`json:"precision"`
	Scale				int32		`json:"scale"`

	Constraint			*string		`json:"constraint"`
	OrdinalPosition		*int32		`json:"ordinalPosition"`
	JsonSchema			*string		`json:"jsonSchema"`
}

func ValidateColumn(column *Column, tableFqn string) error {
	if column.Name == nil {
		return errors.New("invalid column name")
	} else if *column.Name == "" {
		return errors.New("column name cannot be empty")
	}

	if column.FullyQualifiedName == nil {
		v := fmt.Sprintf("%v.%v", tableFqn, column.Name)
		column.FullyQualifiedName = &v
	} else if *column.FullyQualifiedName == "" {
		return errors.New("column fqn cannot be empty")
	}

	if column.DataType != nil {
		_, ok := DataType[*column.DataType]

		if !ok {
			return errors.New("invalid column data type")
		}
	}

	if column.ArrayDataType != nil {
		_, ok := DataType[*column.ArrayDataType]

		if !ok {
			return errors.New("invalid column data array type")
		}
	}

	if column.Constraint != nil {
		_, ok := Constraint[*column.Constraint]

		if !ok {
			return errors.New("invalid column constraint")
		}
	}

	return  nil
}

// Data type
var DataType = map[string]int {
	"NUMBER": 0,
	"TINYINT": 1,
	"SMALLINT": 2,
	"INT": 3,
	"BIGINT": 4,
	"BYTEINT": 5,
	"BYTES": 6,
	"FLOAT": 7,
	"DOUBLE": 8,
	"DECIMAL": 9,
	"NUMERIC": 10,
	"TIMESTAMP": 11,
	"TIMESTAMPZ": 12,
	"TIME": 13,
	"DATE": 14,
	"DATETIME": 15,
	"INTERVAL": 16,
	"STRING": 17,
	"MEDIUMTEXT": 18,
	"TEXT": 19,
	"CHAR": 20,
	"LONG": 21,
	"VARCHAR": 22,
	"BOOLEAN": 23,
	"BINARY": 24,
	"VARBINARY": 25,
	"ARRAY": 26,
	"BLOB": 27,
	"LONGBLOB": 28,
	"MEDIUMBLOB": 29,
	"MAP": 30,
	"STRUCT": 31,
	"UNION": 32,
	"SET": 33,
	"GEOGRAPHY": 34,
	"ENUM": 35,
	"JSON": 36,
	"UUID": 37,
	"VARIANT": 38,
	"GEOMETRY": 39,
	"BYTEA": 40,
	"AGGREGATEFUNCTION": 41,
	"ERROR": 42,
	"FIXED": 43,
	"RECORD": 44,
	"NULL": 45,
	"SUPER": 46,
	"HLLSKETCH": 47,
	"PG_LSN": 48,
	"PG_SNAPSHOT": 49,
	"TSQUERY": 50,
	"TXID_SNAPSHOT": 51,
	"XML": 52,
	"MACADDR": 53,
	"TSVECTOR": 54,
	"UNKNOWN": 55,
	"CIDR": 56,
	"INET": 57,
	"CLOB": 58,
	"ROWID": 59,
	"LOWCARDINALITY": 60,
	"YEAR": 61,
	"POINT": 62,
	"POLYGON": 63,
	"TUPLE": 64,
	"SPATIAL": 65,
	"TABLE": 66,
	"NTEXT": 67,
	"IMAGE": 68,
	"IPV4": 69,
	"IPV6": 70,
	"DATETIMERANGE": 71,
	"HLL": 72,
	"LARGEINT": 73,
	"QUANTILE_STATE": 74,
	"AGG_STATE": 75,
	"BITMAP": 76,
	"UINT": 77,
	"BIT": 78,
	"MONEY": 79,
}

// Constraint
var Constraint = map[string]int {
	"NULL": 0,
	"NOT_NULL": 1,
	"UNIQUE": 2,
	"PRIMARY_KEY": 3,
}

// Table constraint
type TableConstraint struct {
	ConstraintType		*string		`json:"constraintType"`
	Columns				[]string	`json:"columns"`
	ReferredColumns		[]string	`json:"referredColumns"`
	RelationshipType	*string		`json:"relationshipType"`
}

// Constraint type
var ConstraintType = map[string]int {
	"UNIQUE": 0,
	"PRIMARY_KEY": 1,
	"FOREIGN_KEY": 2,
	"SORT_KEY": 3,
	"DIST_KEY": 4,
}

// Relationship type
var RelationshipType = map[string]int {
	"ONE_TO_ONE": 0,
	"ONE_TO_MANY": 1,
	"MANY_TO_ONE": 2,
	"MANY_TO_MANY": 3,
}

// APIs
type GetTableEntitiesQuery struct {
	DatabaseSchema		string	`form:"databaseSchema"`
	Limit 				int		`form:"limit"`
	Offset 				int		`form:"offset"`
}

type GetTableEntityByIdParam struct {
	ID string	`uri:"id" binding:"required"`
}

type GetTableEntityByFqnParam struct {
	FQN string	`uri:"fqn" binding:"required"`
}

type CreateTableEntityPayload struct {
	Name				string				`json:"name" binding:"required"`
	DisplayName			string				`json:"displayName"`
	Description			string				`json:"description"`

	DatabaseSchema		string				`json:"databaseSchema" binding:"required"`

	TableType			string				`json:"tableType" binding:"required"`
	TableConstraints	[]TableConstraint	`json:"tableConstraints"`

	Columns				[]Column			`json:"columns"`
}

func ValidateCreateTableEntityPayload(payload *CreateTableEntityPayload) error {
	_, tableTypeErr := ValidateTableType(payload.TableType)

	if tableTypeErr != nil {
		return tableTypeErr
	}

	tableFqn := fmt.Sprintf("%v.%v", payload.DatabaseSchema, payload.Name)

	for _, column := range payload.Columns {
		err := ValidateColumn(&column, tableFqn)

		if err != nil {
			return err
		}
	}

	return nil
}
