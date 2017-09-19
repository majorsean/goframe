/*
* @Author: wang
* @Date:   2017-09-12 16:46:48
* @Last Modified by:   wang
* @Last Modified time: 2017-09-19 18:09:08
 */

package goframe

// import (
// 	"net/http"
// )

type Controller struct {
	Name string
	Req  *Request
	Resp *Response
}

func NewController(resp *Response, req *Request) *Controller {
	return &Controller{
		Req:  req,
		Resp: resp,
	}
}
