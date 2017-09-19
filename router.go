/*
* @Author: wang
* @Date:   2017-09-12 18:25:53
* @Last Modified by:   wang
* @Last Modified time: 2017-09-19 18:09:14
 */
package goframe

import (
	// "net/http"
	"reflect"
	// "regexp"
	// "log"
	"strings"
)

// type Route struct {
// 	r string
// 	// cr          *regexp.Regexp
// 	method      string
// 	c           interface{}
// 	httpHandler http.Handler
// }

type Route map[string]reflect.Value

// func AddRoute(s Service, r string, c *Controller) {
// cr, err := regexp.Compile(r)
// if err != nil {
// 	log.Printf("regex error:%s->%s", r, err.Error())
// }
// fv := reflect.ValueOf(handler)
// s.routes = append(s.routes, Route{r: r, c: c})
// }

func routeHandler(s *Service, c *Controller) reflect.Value {
	requestPath := c.Req.URL.Path
	// for _, route := range s.routes {
	// if req.Method != route.method {
	// 	continue
	// }

	var action string
	var controller string

	uriSplit := strings.Split(requestPath, "/")
	if len(uriSplit) >= 3 && uriSplit[2] != "" {
		action = uriSplit[2]
	} else {
		action = "index"
	}
	if uriSplit[1] != "" {
		controller = uriSplit[1]
	} else {
		controller = "index"
	}

	route := s.routes[controller+"."+action]
	return route
}
