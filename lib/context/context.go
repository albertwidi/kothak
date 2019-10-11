package context

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/albertwidi/kothak/lib/http/response"
)

// RequestContext struct
type RequestContext struct {
	httpResponseWriter http.ResponseWriter
	httpRequest        *http.Request
	path               string
	method             string
}

// Constructor of context
type Constructor struct {
	HTTPResponseWriter http.ResponseWriter
	HTTPRequest        *http.Request
	Path               string
	Method             string
}

// New context
func New(constructor Constructor) *RequestContext {
	rc := RequestContext{
		httpResponseWriter: constructor.HTTPResponseWriter,
		httpRequest:        constructor.HTTPRequest,
		path:               constructor.Path,
		method:             constructor.Method,
	}
	return &rc
}

// Request return http request from request context
func (rc *RequestContext) Request() *http.Request {
	return rc.httpRequest
}

// ResponseWriter return http response writer from request context
func (rc *RequestContext) ResponseWriter() http.ResponseWriter {
	return rc.httpResponseWriter
}

// JSON to create a json response
func (rc *RequestContext) JSON() *response.JSONResponse {
	j := response.JSON(rc.httpResponseWriter)
	return j
}

// DecodeJSON from request body
func (rc *RequestContext) DecodeJSON(out interface{}) error {
	in, err := ioutil.ReadAll(rc.httpRequest.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(in, out); err != nil {
		return err
	}
	return nil
}
