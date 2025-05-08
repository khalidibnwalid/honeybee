package server

import (
	database "khalidibnwalid/luma_server/internal/database"
)

type ServerContext struct {
	DB database.Database
}

func NewServerContext() (*ServerContext, error) {
	db, err := database.NewClient()
	if err != nil {
		return nil, err
	}

	return &ServerContext{
		DB: *db,
	}, nil
}