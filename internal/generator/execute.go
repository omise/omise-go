package main

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func Execute(j Job) error {
	tmplname, outname := j.Filenames()

	tmpl := template.New(tmplname).Funcs(map[string]interface{}{
		"toLower": strings.ToLower,
	})

	tmpl, e := tmpl.ParseFiles(filepath.Join("internal/generator", tmplname))
	if e != nil {
		return e
	}

	file, e := os.Create(outname)
	if e != nil {
		return e
	}

	defer file.Close()

	ctx, e := NewContext(j)
	if e != nil {
		return e
	}

	if e := tmpl.Execute(file, ctx); e != nil {
		return e
	}

	return nil
}
