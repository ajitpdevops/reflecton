package config

import (
	"html/template"
	"log"
)

// A new type to hold application wide (Global) variables
type AppConfig struct {
	UseCache 		bool
	TemplateCache 	map[string]*template.Template
	InfoLog			*log.Logger			
}