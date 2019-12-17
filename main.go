package main

import (
	"github.com/go-martini/martini"
	"github.com/pedro823/maratona-runtime/handlers"
)

func main() {
	m := martini.Classic()
	handlers.RegisterRoutes(m)
	m.Run()
}