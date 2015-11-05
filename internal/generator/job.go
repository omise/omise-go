package main

import (
	"os"
	"path/filepath"
	"text/template"
)

type Job interface {
	Filenames() (tmplname, outname string)
}

func Execute(j Job) error {
	tmplname, outname := j.Filenames()

	tmpl, e := template.ParseFiles(filepath.Join("internal/generator", tmplname))
	if e != nil {
		return e
	}

	file, e := os.Create(outname)
	if e != nil {
		return e
	}

	ctx, e := NewContext(j)
	if e != nil {
		return e
	}

	defer file.Close()
	if e := tmpl.Execute(file, ctx); e != nil {
		return e
	}

	return nil
}
