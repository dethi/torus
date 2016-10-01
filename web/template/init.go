package template

import (
	"bytes"
	"html/template"
	"io"
)

const (
	DashboardTmpl = "web.dashboard"
)

// Global template register
var tmpl = template.New("root")

func Render(wr io.Writer, name string, data interface{}) error {
	return tmpl.ExecuteTemplate(wr, name, data)
}

func mustRead(pathname string) string {
	r, err := Open(pathname)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	return buf.String()
}

func mustTmpl(name string, pathname string) {
	t, err := tmpl.New(name).Parse(mustRead(pathname))
	if err != nil {
		panic(err)
	}
	tmpl = t
}

func init() {
	mustTmpl(DashboardTmpl, "dashboard.html")
}
