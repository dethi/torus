package web

import (
	"encoding/json"
	"net/http"

	"github.com/dethi/torus/scraper"
	"github.com/gorilla/mux"
)

type SearchParam struct {
	Q string
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	vars := mux.Vars(r)
	if vars["tracker"] != "cpasbien" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var param SearchParam
	if err := decoder.Decode(&param, r.URL.Query()); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := scraper.ScrapeCpasbien(param.Q)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := struct {
		Hits []scraper.CpasbienResult `json:"hits"`
	}{res}
	json.NewEncoder(w).Encode(data)
}
