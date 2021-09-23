package config

import (
	"html/template"
	"log"
)

type AppConfig struct {
	UseCache bool
	TempCache map[string]*template.Template
	InfoLog *log.Logger
}