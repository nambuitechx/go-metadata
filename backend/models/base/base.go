package models

type JsonPatchOperation struct {
	Op   	string    		`json:"op"`
	Path 	string      	`json:"path"`
	Value 	interface{} 	`json:"value"`
}
