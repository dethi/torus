package web

import "net/http"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	URL, err := router.Get(DashboardRoute).URL()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, URL.String(), http.StatusSeeOther)
}
