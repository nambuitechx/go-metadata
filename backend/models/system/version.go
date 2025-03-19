package models

type SystemVersion struct {
	Version			string		`json:"version"`
	Revision		string		`json:"revision"`
	Timestamp		int			`json:"timestamp"`
}
