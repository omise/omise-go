package main

import (
	"fmt"
)

const TargetImport = "github.com/omise/omise-go"

func main() {
	jobs, e := ExtractJobs(TargetImport)
	noError(e)

	for _, job := range jobs {
		noError(Execute(job))
	}
	fmt.Println("done.")
}

func noError(e error) {
	if e != nil {
		panic(e)
	}
}
