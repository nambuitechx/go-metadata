package models

type EntityTotal struct {
	Total	int				`json:"total" db:"total"`
}

type JsonPatchOperation struct {
	Op   	string    		`json:"op"`
	Path 	string      	`json:"path"`
	Value 	interface{} 	`json:"value"`
}
