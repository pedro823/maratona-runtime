package util

import (
	"encoding/json"
	"net/http"

	"github.com/go-martini/martini"
)

type JSONRenderer struct {
	writer  http.ResponseWriter
	options *JSONOptions
	logger *TimeLogger
}

type JSONOptions struct {
	Indent bool
}

var defaultOptions = JSONOptions{Indent: false}

func UseJSONRenderer(options *JSONOptions) martini.Handler {
	return func(res http.ResponseWriter, c martini.Context, logger *TimeLogger) {
		c.Map(&JSONRenderer{writer: res, options: options, logger: logger})
	}
}

func (r *JSONRenderer) JSON(status int, v interface{}) {
	var result []byte
	var err error
	var options JSONOptions
	if r.options != nil {
		options = *r.options
	} else {
		options = defaultOptions
	}

	if options.Indent {
		result, err = json.MarshalIndent(v, "", "    ")
	} else {
		result, err = json.Marshal(v)
	}
	if err != nil {
		r.logger.TimePrintf("Could not marshal JSON response: %v", err)
		http.Error(r.writer, "Internal Server Error", 500)
		return
	}

	// json rendered
	r.writer.Header().Set("Content-Type", "application/json")
	r.writer.WriteHeader(status)
	_, _ = r.writer.Write(result)
}
