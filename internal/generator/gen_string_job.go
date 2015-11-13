package main

import (
	"go/types"
	"strings"

	"bitbucket.org/pkg/inflect"
)

type GenStringJob struct {
	Name   string
	Fields map[string]string
}

func NewGenStringJob(name string, struc *types.Struct) *GenStringJob {
	fields := map[string]string{}
	collectFields(fields, struc)
	return &GenStringJob{name, fields}
}

func collectFields(fields map[string]string, struc *types.Struct) {
	// TODO:
	for i := 0; i < struc.NumFields(); i++ {
		f, tag := struc.Field(i), struc.Tag(i)

		switch {
		case f.Anonymous():
			if named, ok := f.Type().(*types.Named); ok {
				if struc_, ok := named.Underlying().(*types.Struct); ok {
					collectFields(fields, struc_)
				}
			}

		case strings.Contains(tag, `pretty:""`):
			fields[f.Name()] = inflect.Underscore(f.Name())
		}
	}
}

func (job *GenStringJob) Filenames() (string, string) {
	return "gen_string_job.tmpl", inflect.Underscore(job.Name) + "_string.go"
}
