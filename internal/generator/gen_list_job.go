package main

type GenListJob struct {
	Names []string
}

func (job *GenListJob) Filenames() (string, string) {
	return "gen_list_job.tmpl", "gen_list.go"
}
