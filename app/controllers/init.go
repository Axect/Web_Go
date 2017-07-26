package controllers

import (
	"time"

	"github.com/revel/modules/db/app"
	"github.com/revel/revel"
)

func InitDB() {
	db.Init()
	schema := `
CREATE TABLE IF NOT EXISTS posts (
	id INTEGER PRIMARY KEY AUTOINCREMENT
	title TEXT NOT NULL,
	body TEXT NOT NULL,
	created_at DATETIME NOT NULL,
	updated_at DATETIME NOT NULL,
);
`
	db.Db.Exec(schema)
}

func init() {
	revel.OnAppStart(InitDB)

	revel.TemplateFuncs["formatDate"] = func(date time.Time) string {
		return date.Format("2006/01/02 03:04")
	}
}
