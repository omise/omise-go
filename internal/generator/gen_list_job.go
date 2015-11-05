package main

import (
	"bitbucket.org/pkg/inflect"
)

type GenListJob struct {
	Name string
}

func (job *GenListJob) Filenames() (string, string) {
	return "gen_list_job.tmpl", inflect.Underscore(job.Name) + "_list.go"
}
