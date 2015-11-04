package main

import (
	"go/importer"
	"go/types"
	"strings"
)

func ExtractJobs(importpath string) ([]Job, error) {
	pkg, e := importer.Default().Import(importpath)
	if e != nil {
		return nil, e
	}

	scope, jobs := pkg.Scope(), []Job{}
	for _, name := range scope.Names() {
		obj := scope.Lookup(name)
		typ := obj.Type().Underlying()

		// skips
		if strings.HasSuffix(name, "List") {
			continue
		}

		// generates
		if _, isStruct := typ.(*types.Struct); isStruct {
			jobs = append(jobs, &GenListJob{name})
		}
	}

	return jobs, nil
}
