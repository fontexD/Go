package models

// Template data holds data sent from handlers to templates
type Templatedata struct {
	StringMap  map[string]string
	StringBool map[string]bool
	IntMap     map[string]int
	FloatMap   map[string]float32
	Data       map[string]interface{}
	CSRFToken  string
	Flash      string
	Warning    string
	Error      string
}
