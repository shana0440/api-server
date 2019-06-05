package models

import (
	"encoding/json"
	"net/http"

	"github.com/dannypsnl/rocket/response"
	log "github.com/sirupsen/logrus"
)

// Response is a wrapper of rocket response, use to provider more usable function
type Response struct {
	resp *response.Response
}

// NewResponse is a factory function of Response
func NewResponse() *Response {
	return &Response{}
}

// JSON is use to marshal json response
func (r *Response) JSON(status int, body interface{}) *response.Response {
	str, err := json.Marshal(body)
	if err != nil {
		log.Error("cannot parse body", err)
		r.resp = NewResponse().JSON(http.StatusInternalServerError, ErrorResponse{
			Message: "Oops, server internal error",
		})
	} else {
		r.resp = response.New(response.Json(str)).Status(status)
	}
	return r.resp
}
