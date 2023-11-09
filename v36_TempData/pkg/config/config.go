package config

import (
	"html/template"
	"log"
)

// AppConfig struct of configure
type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache bool
	InfoLog *log.Logger
}