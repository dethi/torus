package main

import (
	"net/http"
	"sort"

	"github.com/dethi/goutil/fs"
)

const GB = 1024 * 1024 * 1024

type byDate []Record

func (p byDate) Len() int {
	return len(p)
}

func (p byDate) Less(i, j int) bool {
	return p[i].EndTime.Before(p[j].EndTime)
}

func (p byDate) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func listView(w http.ResponseWriter, r *http.Request) {
	fsStat, err := fs.GetFsStats(cfg.DataPath)
	if err != nil {
		fsStat = &fs.FsStats{Available: 0}
	}

	type Result struct {
		Name     string
		InfoHash string
		Date     int64
	}
	var res []Result

	records := db.ViewRecords()
	sort.Sort(sort.Reverse(byDate(records)))
	for _, r := range records {
		res = append(res, Result{
			Name:     r.Name,
			InfoHash: r.InfoHash,
			Date:     r.EndTime.UnixNano() / 1e6,
		})
	}

	data := struct {
		Version            string
		AvailableDiskSpace uint64
		Records            []Result
	}{
		Version:            version,
		AvailableDiskSpace: fsStat.Available / GB,
		Records:            res,
	}

	listTmpl.Execute(w, data)
}
