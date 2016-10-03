package web

import (
	"encoding/json"
	"net/http"

	"github.com/dethi/torus"
)

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	data := struct {
		Version   string `json:"version"`
		Build     string `json:"build"`
		BuildTime string `json:"build_time"`
		Revision  string `json:"revision"`
	}{
		Version:   torus.Version,
		Build:     torus.Build,
		BuildTime: torus.BuildTime,
		Revision:  torus.Revision(),
	}
	json.NewEncoder(w).Encode(data)
}
