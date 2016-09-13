package main

import (
	"net/http"
	"sort"
	"time"

	"github.com/dethi/goutil/fs"
	"github.com/hako/durafmt"
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
	fsStat, err := fs.GetFsStats(*dataPath)
	if err != nil {
		fsStat = &fs.FsStats{Available: 0}
	}

	type Result struct {
		Name     string
		InfoHash string
		Since    *durafmt.Durafmt
	}
	var res []Result

	records := db.ViewRecords()
	sort.Sort(sort.Reverse(byDate(records)))
	for _, r := range records {
		res = append(res, Result{
			Name:     r.Name,
			InfoHash: r.InfoHash,
			Since:    durafmt.Parse(time.Since(r.EndTime)),
		})
	}

	data := struct {
		AvailableDiskSpace uint64
		Records            []Result
	}{
		AvailableDiskSpace: fsStat.Available / GB,
		Records:            res,
	}

	listTmpl.Execute(w, data)
}
