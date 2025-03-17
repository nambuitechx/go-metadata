package models

type EntityReference struct {
	ID					string		`json:"id"`
	Type 				string		`json:"type"`
	Name 				string		`json:"name"`
	FullyQualifiedName	string		`json:"fullyQualifiedName"`

	DisplayName			string		`json:"displayName"`
	Description			string		`json:"description"`

	Deleted				bool		`json:"deleted"`
}
