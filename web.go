/*
* @Author: wang
* @Date:   2017-09-12 17:02:15
* @Last Modified by:   wang
* @Last Modified time: 2017-09-19 18:09:19
 */

package goframe

import (
	// "fmt"
	"log"
	"net"
	"net/http"
	"reflect"
	// "regexp"
	"strings"
)

var (
	s *Service
)

func init() {
	s = NewService()
}

func Run() {
	mux := http.NewServeMux()
	mux.Handle("/", s)

	l, err := net.Listen("tcp", ":10198")
	if err != nil {
		log.Fatalf("Listen err %v", err)
	}
	defer l.Close()

	log.Printf("Listening on %s", l.Addr())
	http.Serve(l, mux)
}

// func handler(rw http.ResponseWriter, req *http.Request) {
// 	controller := NewController(rw, req)
// }

func AddRoute(r string, c interface{}) {
	cVal := reflect.ValueOf(c)
	if !cVal.IsValid() {
		return
	}
	cType := reflect.TypeOf(c)
	elem := cType.Elem()
	elemName := strings.ToLower(strings.TrimSuffix(elem.Name(), "Controller"))

	for i := 0; i < elem.NumMethod(); i++ {
		m := elem.Method(i)
		if strings.HasSuffix(m.Name, "Action") {
			method := cVal.MethodByName(m.Name)
			if method.IsValid() {
				mName := strings.ToLower(strings.TrimSuffix(m.Name, "Action"))
				s.routes[elemName+"."+mName] = method
			}
		}
	}
	// s.routes = append(s.routes, Route{r: r, c: c})
}
