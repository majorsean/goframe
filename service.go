/*
* @Author: wang
* @Date:   2017-09-13 11:11:13
* @Last Modified by:   wang
* @Last Modified time: 2017-09-19 18:09:17
 */
package goframe

import (
	"net/http"
	"reflect"
	// "regexp"
	// "log"
	// "os"
)

type Service struct {
	Config Config
	routes map[string]reflect.Value
}

func NewService() *Service {
	return &Service{
		Config: Config{
			HttpAddr: "0.0.0.0",
			HttpPort: "18080",
		},
		routes: make(map[string]reflect.Value),
	}
}

func (s *Service) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	s.Process(rw, req)
}

func (s *Service) Process(rw http.ResponseWriter, req *http.Request) {
	response := NewResponse(rw)
	request := NewRequest(req)
	controller := NewController(response, request)
	handler := routeHandler(s, controller)
	if !handler.IsValid() {
		http.NotFound(rw, req)
	} else {
		var args []reflect.Value
		ret := handler.Call(args)
		content := []byte(ret[0].String())
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		rw.Write(content)
	}
}
