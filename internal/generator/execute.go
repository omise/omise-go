package main

import (
	"os"
	"path/filepath"
	"text/template"
)

func Execute(j Job) error {
	tmplname, outname := j.Filenames()

	// prepare template
	tmpl, e := template.ParseFiles(filepath.Join("internal/generator", tmplname))
	if e != nil {
		return e
	}

	// open output file
	file, e := os.Create(outname)
	if e != nil {
		return e
	}

	ctx, e := NewContext(j)
	if e != nil {
		return e
	}

	defer file.Close()

	// writeout
	if e := tmpl.Execute(file, ctx); e != nil {
		return e
	}

	return nil
}
