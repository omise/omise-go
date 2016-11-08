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
