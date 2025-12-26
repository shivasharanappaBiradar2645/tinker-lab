package main

import (
	"net/http"

	"github.com/shivasharanappaBiradar2645/tinker-lab/golang/rss-aggregator/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)
