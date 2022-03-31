package rest

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type APIResponse struct {
	IsOk     bool
	response *http.Response
	Error    APIError
}

func (ar *APIResponse) Body() io.ReadCloser {
	return ar.response.Body
}

func (ar *APIResponse) ReadBody() ([]byte, error) {
	defer ar.response.Body.Close()
	return ioutil.ReadAll(ar.response.Body)
}

func (ar *APIResponse) StatusCode() int {
	return ar.response.StatusCode
}

func (ar *APIResponse) Location() (*url.URL, error) {
	return ar.response.Location()
}

type APIError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (ae *APIError) ToString() string {
	return fmt.Sprintf("Err Code: %d, Err: %s", ae.Code, ae.Message)
}
