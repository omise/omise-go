package main

import (
	"fmt"
)

type Job interface {
	Filenames() (tmplname, outname string)
}

type GenListJob struct {
	Names []string
}

func (job *GenListJob) Filenames() (string, string) {
	return "gen_list_job.tmpl", "gen_list.go"
}

type GenSearchJob struct {
	Names []string
}

func (job *GenSearchJob) Filenames() (string, string) {
	return "gen_search_job.tmpl", "gen_search.go"
}

type GenStringJob struct {
	Fields map[string][]string
}

func (job *GenStringJob) Filenames() (string, string) {
	return "gen_string_job.tmpl", "gen_string.go"
}

type GenAPITreeJob struct {
	Names      []string
	Operations map[string][]string
}

func (job *GenAPITreeJob) Filenames() (string, string) {
	return "gen_api_tree.tmpl", "api/gen_api_tree.go"
}

func (job *GenAPITreeJob) OpSignature(model, op string) string {
	resultType, opType := job.typeNames(model, op)
	return "(op *" + opType + ") (*" + resultType + ", error)"
}

func (job *GenAPITreeJob) OpInvocation(model, op string) string {
	resultType, opType := job.typeNames(model, op)
	return fmt.Sprintf(`result := &%s{}
		if op == nil {
			op = &%s{}
		}
		if e := client.Do(result, op); e != nil {
			return nil, e
		}
		return result, nil`, resultType, opType)
}

func (job *GenAPITreeJob) typeNames(model, op string) (string, string) {
	switch op {
	case "List":
		return "omise." + model + "List", "operations.List" + model + "s"
	case "Destroy":
		return "omise.Deletion", "operations.Destroy" + model
	default:
		return "omise." + model, "operations." + op + model
	}
}
