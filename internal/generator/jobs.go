package main

type Job interface {
	Filenames() (tmplname, outname string)
}

type GenListJob struct {
	Names []string
}

func (job *GenListJob) Filenames() (string, string) {
	return "gen_list_job.tmpl", "gen_list.go"
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
	return "gen_api_tree.tmpl", "gen_api_tree.go"
}
