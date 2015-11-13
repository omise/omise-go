package main

import (
	"go/importer"
	"go/types"
)

func ExtractJobs(importpath string) ([]Job, error) {
	pkg, e := importer.Default().Import(importpath)
	if e != nil {
		return nil, e
	}

	scope, jobs := pkg.Scope(), []Job{}
	for _, name := range scope.Names() {
		obj := scope.Lookup(name)
		typ := obj.Type()

		switch findClass(typ) {
		case ModelStruct:
			struc := typ.Underlying().(*types.Struct)
			jobs = append(jobs, NewGenListJob(name, struc))
			jobs = append(jobs, NewGenStringJob(name, struc))
		}
	}

	return jobs, nil
}

func findClass(typ types.Type) (class Class) {
	named, ok := typ.(*types.Named)
	if !ok {
		return Uninterested
	}

	var struc *types.Struct
	if struc, ok = named.Underlying().(*types.Struct); !ok {
		return Uninterested
	}

	for i := 0; i < struc.NumFields(); i++ {
		f := struc.Field(i)
		switch {
		case !f.Anonymous():
			continue
		case f.Name() == "Base":
			return ModelStruct
		case f.Name() == "List":
			return ListStruct
		}
	}

	return Uninterested
}
