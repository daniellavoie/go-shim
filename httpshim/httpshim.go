// Code generated by counterfeiter. DO NOT EDIT.
// with command: counterfeiter -p net/http
package httpshim

import (
	"net/http"
)

type HttpShim struct{}

func (sh *HttpShim) Get(url string) (resp *http.Response, err error) {
	return http.Get(url)
}
