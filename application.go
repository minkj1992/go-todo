package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

// Route routes url
type Route struct {
	method  string
	pattern *regexp.Regexp
	handler http.Handler
}

// Application grouping routes
type Application struct {
	routes []*Route
}

// Add is
func (a *Application) Add(method, path string, handler http.Handler) {
	a.routes = append(a.routes, &Route{
		method,
		regexp.MustCompile(path),
		handler,
	})
}

// AddFunc is ...
func (a *Application) AddFunc(method, path string, handler func(rw http.ResponseWriter, r *http.Request)) {
	a.Add(method, path, http.HandlerFunc(handler))
}

func (a *Application) Get(path string, handler func(rw http.ResponseWriter, r *http.Request)) {
	a.AddFunc(http.MethodGet, path, handler)
}

func (a *Application) Post(path string, handler func(rw http.ResponseWriter, r *http.Request)) {
	a.AddFunc(http.MethodPost, path, handler)
}

func (a *Application) Delete(path string, handler func(rw http.ResponseWriter, r *http.Request)) {
	a.AddFunc(http.MethodDelete, path, handler)
}

func (a *Application) Put(path string, handler func(rw http.ResponseWriter, r *http.Request)) {
	a.AddFunc(http.MethodPut, path, handler)
}

func (a *Application) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	for _, route := range a.routes {
		matched := route.pattern.MatchString(r.URL.Path) && route.method == r.Method
		if matched {
			route.handler.ServeHTTP(rw, r)
			return
		}
	}
	http.NotFound(rw, r)
}

func (a *Application) Static(root string) {
	fs := http.FileServer(http.Dir(root))
	a.Add(http.MethodGet, "/*", fs)
}

func (a *Application) Start(port string) {
	fmt.Println("server is running http://localhost" + port)
	http.ListenAndServe(port, a)
}

func NewApplication() *Application {
	return &Application{}
}

func Bind(r *http.Request, i interface{}) {
	json.NewDecoder(r.Body).Decode(i)
}

func QueryParam(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

func Json(rw http.ResponseWriter, i interface{}) {
	enc := json.NewEncoder(rw)
	enc.Encode(&i)
}
