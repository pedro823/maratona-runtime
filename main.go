package main

import (
	"github.com/go-martini/martini"
	"github.com/subosito/gotenv"

	"github.com/pedro823/maratona-runtime/database"
	"github.com/pedro823/maratona-runtime/handlers"
	"github.com/pedro823/maratona-runtime/util"
)

func main() {
	gotenv.Load()

	db := database.NewDatabase()

	m := martini.Classic()
	m.Map(db)
	m.Use(util.UseJSONRenderer(nil))

	handlers.RegisterRoutes(m)
	m.Run()
}
