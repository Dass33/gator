package main

import (
	"github.com/Dass33/gator/internal/config"
	"github.com/Dass33/gator/internal/database"
)

type state struct {
	config *config.Config
	db     *database.Queries
}
