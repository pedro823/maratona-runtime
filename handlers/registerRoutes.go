package handlers

import (
	"github.com/go-martini/martini"
	"github.com/pedro823/maratona-runtime/middleware"
)

func RegisterRoutes(m *martini.ClassicMartini) {
	m.Get("/_healthcheck", HealthCheck)

	m.Group("/challenge", func(r martini.Router) {
		r.Get("/all", GetAllChallenges)
		r.Post("/upload", UploadChallenge)
	}, middleware.RequireToken)

	m.Group("/attempt", func(r martini.Router) {
		m.Post("/challenge/:challenge", ExtractAndExecute)
	}, middleware.RequireCTFID)
}
