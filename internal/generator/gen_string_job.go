package main

type GenStringJob struct {
	Fields map[string][]string
}

func (job *GenStringJob) Filenames() (string, string) {
	return "gen_string_job.tmpl", "gen_string.go"
}
