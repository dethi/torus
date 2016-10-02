package web

import "net/http"

func init() {
	http.HandleFunc("/", WebHandler)
}

func WebHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		r.URL.Path = "/index.html"
	}
	ServeHTTP(w, r)
}
