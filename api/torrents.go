package web

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dethi/torus"
)

func ListTorrentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	payload := struct {
		Data []torus.Record `json:"data"`
	}{
		Data: []torus.Record{
			torus.Record{
				ID: 1,
				Request: torus.Request{
					UserId:             1,
					URLs:               nil,
					State:              torus.Completed,
					RequestTime:        time.Now(),
					DownloadDuration:   5 * time.Minute,
					ProcessingDuration: 2 * time.Minute,
				},
				Torrents: []torus.Torrent{
					torus.Torrent{
						Name:     "Pirate 5",
						Size:     234567,
						InfoHash: "aqr3ew3221d3242",
						Files:    nil,
						Payload:  nil,
					},
				},
			},
		},
	}

	json.NewEncoder(w).Encode(payload)
}
