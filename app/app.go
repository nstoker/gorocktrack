package app

import (
	"database/sql"
	"text/template"
)

// Application Version
const Version = "0.0.1"

// DB is the main system database
var DB *sql.DB

// TPL is application templates
var TPL *template.Template
