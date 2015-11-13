package main

import (
	"go/types"

	"bitbucket.org/pkg/inflect"
)

type GenListJob struct {
	Name string
}

func NewGenListJob(name string, struc *types.Struct) *GenListJob {
	return &GenListJob{name}
}

func (job *GenListJob) Filenames() (string, string) {
	return "gen_list_job.tmpl", inflect.Underscore(job.Name) + "_list.go"
}
