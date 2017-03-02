package main

import (
	"go/importer"
	"go/types"
	"log"
	"sort"
	"strings"
	"unicode"
)

func ExtractJobs() ([]Job, error) {
	pkg, e := importer.Default().Import(MainImportPath)
	if e != nil {
		return nil, e
	}

	// main package types
	var (
		scope     = pkg.Scope()
		searchJob = &GenSearchJob{}
		stringJob = &GenStringJob{map[string][]string{}}
	)

	for _, name := range scope.Names() {
		obj := scope.Lookup(name)
		typ := obj.Type()

		switch findClass(typ) {
		case ModelStruct:
			log.Println("discovered:", name)
			if name == "SearchResult" {
				continue
			}

			struc := typ.(*types.Named).Underlying().(*types.Struct)
			stringJob.Fields[name] = collectFields(struc)
		}
	}

	// search results are fixed, for now.
	searchJob.Names = []string{
		"Charge",
		"Customer",
		"Dispute",
		"Recipient",
	}

	sort.Strings(searchJob.Names)
	return []Job{searchJob, stringJob}, nil
}

func findClass(typ types.Type) (class Class) {
	named, ok := typ.(*types.Named)
	if !ok {
		return Uninterested
	}

	name := named.Obj().Name()
	switch name {
	case "List", "Base":
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

func collectFields(struc *types.Struct) []string {
	result := []string{}

	for i := 0; i < struc.NumFields(); i++ {
		f, tag := struc.Field(i), struc.Tag(i)

		switch {
		case f.Anonymous():
			if named, ok := f.Type().(*types.Named); ok {
				if struc_, ok := named.Underlying().(*types.Struct); ok {
					result = append(result, collectFields(struc_)...)
				}
			}

		case strings.Contains(tag, `pretty:""`):
			result = append(result, f.Name())
		}
	}

	return result
}

func splitOpName(name string) (string, string) {
	for i, r := range name {
		if i != 0 && unicode.IsUpper(r) {
			return name[i:], name[:i]
		}
	}

	return name, ""
}
