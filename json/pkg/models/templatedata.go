package models

// Template data holds data sent from handlers to templates
type Templatedata struct {
	Name   string `json:"Name"`
	Host   string
	Group  string `json:"Group"`
	Env    string `json:"Env"`
	Status string `json:"Status"`
}
