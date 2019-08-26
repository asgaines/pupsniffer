package pupservice

import (
	"html/template"
	"io"
	"path"

	"github.com/asgaines/pupsniffr/pup"
)

func (p pupsvc) PupReport(total int, pups []*pup.Pup, wr io.Writer) error {
	file := path.Join(p.staticPath, "report.html")
	tmpl := template.Must(template.ParseFiles(file))

	data := struct {
		Pups []*pup.Pup
		Art  string
	}{
		Pups: pups,
		Art:  pup.WoofASCII,
	}

	return tmpl.Execute(wr, data)
}
