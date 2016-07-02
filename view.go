package main

import (
	"net/http"

	"github.com/dethi/goutil/fs"
)

const GB = 1024 * 1024 * 1024

func listView(w http.ResponseWriter, r *http.Request) {
	fsStat, err := fs.GetFsStats(*dataPath)
	if err != nil {
		fsStat = &fs.FsStats{Available: 0}
	}

	data := struct {
		AppId              string
		SearchKey          string
		IndexName          string
		AvailableDiskSpace uint64
	}{
		AppId:              *algoliaAppId,
		SearchKey:          *algoliaSearchKey,
		IndexName:          *algoliaIndex,
		AvailableDiskSpace: fsStat.Available / GB,
	}

	listTmpl.Execute(w, data)
}
