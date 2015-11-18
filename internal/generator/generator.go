package main

import (
	"log"
)

const (
	MainImportPath       = "github.com/omise/omise-go"
	OperationsImportPath = "github.com/omise/omise-go/operations"
)

func main() {
	jobs, e := ExtractJobs()
	noError(e)

	log.Println(len(jobs), "job(s):")
	for _, job := range jobs {
		_, outname := job.Filenames()
		log.Printf("* %#v -> %s", job, outname)

		noError(Execute(job))
	}

	log.Println("done.")
}

func noError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
