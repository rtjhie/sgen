package sgen

import (
	"text/template"

	"github.com/rtjhie/sgen/internal/bintmpl"
)

//go:generate go run github.com/go-bindata/go-bindata/go-bindata -o=internal/bintmpl/bintmpl.go -pkg=bintmpl -modtime=1 ./schema.go ./template/...

var (
	templates = template.New("templates")
	funcs     = template.FuncMap{
		"snake":  snake,
		"camel":  camel,
		"pascal": pascal,
	}
)

func initTemplates() {
	templates.Funcs(funcs)
	for _, asset := range bintmpl.AssetNames() {
		templates = template.Must(templates.Parse(string(bintmpl.MustAsset(asset))))
	}
}
