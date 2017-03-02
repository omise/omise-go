// +build ignore

package main

//go:generate go run list_types_generator.go
//go:generate go fmt list_types.go

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

const (
	ModelsFile   = "models.txt"
	TemplateFile = "list_types.tmpl"
	OutputFile   = "list_types.go"
)

type Context struct {
	Models []string
}

func main() {
	bytes, e := ioutil.ReadFile(ModelsFile)
	if e != nil {
		log.Fatalln(e)
	}

	models := strings.Split(string(bytes), "\n")
	if models[len(models)-1] == "" {
		models = models[:len(models)-1]
	}

	tmpl := template.Must(template.ParseFiles(TemplateFile))

	outfile, e := os.Create(OutputFile)
	if e != nil {
		log.Fatalln(e)
	}

	if e := tmpl.Execute(outfile, &Context{Models: models}); e != nil {
		log.Fatalln(e)
	}
}
