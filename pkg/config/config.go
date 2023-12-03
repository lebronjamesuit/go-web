package config

import (
	"text/template"
)

type AppConfig struct {
	MyCaches map[string]*template.Template
}
