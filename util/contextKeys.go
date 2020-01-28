package util

import (
	"github.com/go-martini/martini"
)

type ContextKey string

const (
	CompilerContextKey ContextKey = "compiler"
	UserContextKey     ContextKey = "user"
)

type ContextMap map[ContextKey]interface{}

func UseContextMap() martini.Handler {
	return func(c martini.Context) {
		c.Map(ContextMap{})
	}
}
