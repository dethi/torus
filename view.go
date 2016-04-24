package main

import "net/http"

var listTmpl = MustTemplate("tmpl/list.html.tmpl")

func listView(w http.ResponseWriter, r *http.Request) {
	data := struct {
		AppId     string
		SearchKey string
		IndexName string
	}{
		AppId:     *algoliaAppId,
		SearchKey: *algoliaSearchKey,
		IndexName: *algoliaIndex,
	}

	listTmpl.Execute(w, data)
}
