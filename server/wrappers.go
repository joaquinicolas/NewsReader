package server

import (
	"net/http"
)

//A Handler represent a request controller
type Handler func(writter http.ResponseWriter, reader *http.Request)

//PostOnly wrappper
func PostOnly(h Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			h(w, r)
			return
		}
		http.Error(w, "post only", http.StatusMethodNotAllowed)
	}
}

//GetOnly wrapper
func GetOnly(h Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			h(w, r)
			return
		}
		http.Error(w, "get only", http.StatusMethodNotAllowed)
	}
}
