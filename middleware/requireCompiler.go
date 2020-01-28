package middleware

import (
	"github.com/pedro823/maratona-runtime/errors"
	"github.com/pedro823/maratona-runtime/runtime"
	"github.com/pedro823/maratona-runtime/util"
	"net/http"
)

var (
	compilerTranslation = map[string]runtime.AvailableCompiler{
		"python3": runtime.Python3{},
		"c++11":   runtime.CPlusPlus11{},
		"c":       runtime.C{},
		"java8":   runtime.Java8{},
		"go":      runtime.Go{},
	}
)

func RequireCompiler(req *http.Request, res *util.JSONRenderer, context util.ContextMap) {
	compilerHeader := req.Header.Get("Compiler")

	compiler, ok := compilerTranslation[compilerHeader]
	if !ok {
		errors.NewHTTPError(400, "compiler parameter is empty or compiler was not found").WriteJSON(res)
	}

	context[util.CompilerContextKey] = compiler
}
