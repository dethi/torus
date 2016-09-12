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
		AvailableDiskSpace uint64
	}{
		AvailableDiskSpace: fsStat.Available / GB,
	}

	listTmpl.Execute(w, data)
}
