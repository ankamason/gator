package main

import (
	"github.com/ankamason/gator/internal/config"
	"github.com/ankamason/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}
