package web

import (
	"net/http"

	"github.com/dethi/goutil/fs"
	"github.com/dethi/torus"
	"github.com/dethi/torus/web/template"
)

const GB = 1 << 30

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	var diskSpace uint64
	if stats, err := fs.GetStats(""); err == nil {
		diskSpace = stats.Available / GB
	}

	data := struct {
		Version            string
		AvailableDiskSpace uint64
		Records            []struct {
			Name     string
			InfoHash string
			Date     int64
		}
	}{
		Version:            torus.Revision(),
		AvailableDiskSpace: diskSpace,
		Records:            nil,
	}

	template.Render(w, template.DashboardTmpl, data)
}
