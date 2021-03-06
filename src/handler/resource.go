package handler

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"net/http"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Resource(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		resourceGet(w, r)

	default:
		http.NotFound(w, r)
	}

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func resourceGet(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	http.ServeFile(w, r, path)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
