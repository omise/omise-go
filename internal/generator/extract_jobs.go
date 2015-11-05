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
		if !isModelStruct(obj.Type()) {
			continue
		}

		// generates
		jobs = append(jobs, &GenListJob{name})
	}

	return jobs, nil
}

func isModelStruct(typ types.Type) (ok bool) {
	var named *types.Named
	if named, ok = typ.(*types.Named); !ok {
		return false
	}

	var struc *types.Struct
	if struc, ok = named.Underlying().(*types.Struct); !ok {
		return false
	}

	if strings.HasSuffix(named.Obj().Name(), "List") {
		return false
	}

	for i := 0; i < struc.NumFields(); i++ {
		if f := struc.Field(i); f.Anonymous() && f.Name() == "Base" {
			return true
		}
	}

	return false
}
