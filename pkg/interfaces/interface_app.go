package interfaces

import "database/sql"

type AppInterface interface {
	GetDB() *sql.DB
}
