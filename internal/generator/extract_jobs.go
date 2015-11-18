package main

import (
	"fmt"
	"go/importer"
	"go/types"
	"strings"
)

func ExtractJobs() ([]Job, error) {
	pkg, e := importer.Default().Import(MainImportPath)
	if e != nil {
		return nil, e
	}

	scope := pkg.Scope()
	listJob, stringJob := &GenListJob{}, &GenStringJob{map[string][]string{}}
	for _, name := range scope.Names() {
		obj := scope.Lookup(name)
		typ := obj.Type()

		switch findClass(typ) {
		case ModelStruct:
			fmt.Println("MODEL: ", name)
			struc := typ.(*types.Named).Underlying().(*types.Struct)
			listJob.Names = append(listJob.Names, name)
			stringJob.Fields[name] = collectFields(struc)
		}
	}

	fmt.Printf("%#v", stringJob)
	return []Job{listJob, stringJob}, nil
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
		fmt.Println(f.Name(), tag)

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
