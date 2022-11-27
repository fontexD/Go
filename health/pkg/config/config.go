package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// appconfig for the whole application
type Appconfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
