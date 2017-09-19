/*
* @Author: wang
* @Date:   2017-09-18 17:27:41
* @Last Modified by:   wang
* @Last Modified time: 2017-09-19 18:09:12
 */
package goframe

import (
	"net/http"
)

type Response struct {
	Out http.ResponseWriter
}

func NewResponse(w http.ResponseWriter) *Response {
	return &Response{Out: w}
}
