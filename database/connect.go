package database

import (
	"database/sql"
)

func Connect() {
	db, err := sql.Open("mysql", "")
}
