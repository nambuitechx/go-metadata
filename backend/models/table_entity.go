package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
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
	ID					string				`json:"id"`
	Name				string				`json:"name"`
	FullyQualifiedName	string				`json:"fullyqualifiedName"`
	
	DisplayName			string				`json:"displayName"`
	Description			string				`json:"description"`

	ServiceType			string				`json:"serviceType"`
	Service				EntityReference		`json:"service"`
	Database			EntityReference		`json:"database"`
	DatabaseSchema		EntityReference		`json:"databaseSchema"`

	TableType			string				`json:"tableType"`
	TableConstraints	[]TableConstraint	`json:"tableConstraints"`

	Deleted				bool				`json:"deleted"`
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

// Column
type Column struct {
	Name				string		`json:"name"`
	FullyQualifiedName	string		`json:"fullyQualifiedName"`

	DisplayName			string		`json:"displayName"`
	Description			string		`json:"description"`

	DataType			string		`json:"dataType"`
	ArrayDataType		string		`json:"arrayDataType"`
	DataLength			int32		`json:"dataLength"`
	DataTypeDisplay		string		`json:"dataTypeDisplay"`

	Precision			int32		`json:"precision"`
	Scale				int32		`json:"scale"`

	Constraint			string		`json:"constraint"`
	OrdinalPosition		int32		`json:"ordinalPosition"`
	JsonSchema			string		`json:"jsonSchema"`
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
	ConstraintType		string		`json:"constraintType"`
	Columns				[]string	`json:"columns"`
	ReferredColumns		[]string	`json:"referredColumns"`
	RelationshipType	string		`json:"relationshipType"`
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
