package controllers

import "github.com/richardbahrii-emodo/go-rest/database"

type HandlerWithDb struct {
	DB database.Database
}

func NewHandlerWithDb(db database.Database) *HandlerWithDb {
	return &HandlerWithDb{
		DB: db,
	}
}
