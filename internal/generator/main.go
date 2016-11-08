package main

import (
	"log"
)

const MainImportPath = "github.com/omise/omise-go"

func main() {
	jobs, e := ExtractJobs()
	noError(e)

	log.Println(len(jobs), "job(s):")
	for _, job := range jobs {
		_, outname := job.Filenames()
		log.Printf("* %s\n\033[90m%#v\033[0m", outname, job)

		noError(Execute(job))
	}

	log.Println("done.")
}

func noError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
