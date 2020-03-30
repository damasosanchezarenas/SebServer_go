package main

import (
	"net/http"
)

//Route to handler
type Router struct { //methodRest-Path-Handler
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc), //empty map
	}
}

func (r *Router) findHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, exist := r.rules[path] //Verify if exists URL
	handler, methodExists := r.rules[path][method]
	return handler, methodExists, exist
}

//Router for all Requests
func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, exist, methodExists := r.findHandler(request.URL.Path, request.Method)

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !methodExists {
		w.WriteHeader(http.StatusMethodNotAllowed) // URL exists but not with this method
		return
	}

	handler(w, request)
}
