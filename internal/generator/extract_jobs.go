package main

import (
	"fmt"
	"go/importer"
	"go/types"
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
		scope      = pkg.Scope()
		listJob    = &GenListJob{}
		searchJob  = &GenSearchJob{}
		stringJob  = &GenStringJob{map[string][]string{}}
		apiTreeJob = &GenAPITreeJob{nil, map[string][]string{}}
	)

	for _, name := range scope.Names() {
		obj := scope.Lookup(name)
		typ := obj.Type()

		switch findClass(typ) {
		case ModelStruct:
			fmt.Println("MODEL:", name)
			struc := typ.(*types.Named).Underlying().(*types.Struct)
			listJob.Names = append(listJob.Names, name)
			stringJob.Fields[name] = collectFields(struc)
		}
	}

	// operations package types
	pkg, e = importer.Default().Import(OperationsImportPath)
	if e != nil {
		return nil, e
	}

	scope = pkg.Scope()
	for _, name := range scope.Names() {
		obj := scope.Lookup(name)
		typ := obj.Type()

		switch findClass(typ) {
		case OperationStruct:
			fmt.Println("OP:", name)

			modelName, opName := splitOpName(name)
			if strings.HasSuffix(modelName, "s") {
				modelName = modelName[:len(modelName)-1]
			}

			if ops, ok := apiTreeJob.Operations[modelName]; ok {
				apiTreeJob.Operations[modelName] = append(ops, opName)

			} else {
				apiTreeJob.Names = append(apiTreeJob.Names, modelName)
				apiTreeJob.Operations[modelName] = []string{opName}
			}
		}
	}

	// search results are fixed, for now.
	searchJob.Names = []string{
		"Charge",
		"Customer",
		"Dispute",
		"Recipient",
	}

	sort.Strings(listJob.Names)
	sort.Strings(searchJob.Names)
	sort.Strings(apiTreeJob.Names)
	return []Job{listJob, searchJob, stringJob, apiTreeJob}, nil
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

	for i := 0; i < named.NumMethods(); i++ {
		m := named.Method(i)
		if m.Name() == "Op" {
			return OperationStruct
		}
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
