package main

import (
	"log"
	"os"

	"github.com/go-martini/martini"
	"github.com/subosito/gotenv"

	"github.com/pedro823/maratona-runtime/database"
	"github.com/pedro823/maratona-runtime/handlers"
	"github.com/pedro823/maratona-runtime/util"
)

func main() {
	logger := log.New(os.Stdout, "[maratona-runtime] ", 0)

	err := gotenv.Load()
	if err != nil {
		logger.Printf("[WARNING] error while loading env file: %s", err.Error())
	}

	db := database.NewDatabase(logger)

	m := martini.Classic()
	m.Map(db)
	m.Logger(logger)
	m.Map(util.NewTimeLogger(logger))
	m.Use(util.UseJSONRenderer(nil))

	handlers.RegisterRoutes(m)
	m.Run()
}
