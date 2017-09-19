/*
* @Author: wang
* @Date:   2017-09-18 17:29:40
* @Last Modified by:   wang
* @Last Modified time: 2017-09-19 18:09:10
 */

package goframe

import (
	"net/http"
)

type Request struct {
	*http.Request
}

func NewRequest(r *http.Request) *Request {
	return &Request{Request: r}
}
