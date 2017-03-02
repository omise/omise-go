package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type Context struct {
	Models []string
}

var AllModels = []string{
	"Account",
	"Balance",
	"BankAccount",
	"Card",
	"Charge",
	"Customer",
	"Deletion",
	"Dispute",
	"Document",
	"Event",
	"Link",
	"Recipient",
	"Refund",
	"Token",
	"Transaction",
	"Transfer",
}

func main() {
	name := strings.TrimSpace(os.Args[1])
	if name == "" {
		log.Fatalln("template name must be given as first argument")
	}

	var (
		templateFile = name + ".tmpl"
		outputFile   = name + ".go"
	)

	tmpl := template.Must(template.New("root").
		Funcs(map[string]interface{}{"toLower": strings.ToLower}).
		ParseFiles(templateFile))

	outfile, e := os.Create(outputFile)
	if e != nil {
		log.Fatalln(e)
	}

	if e := tmpl.ExecuteTemplate(outfile, templateFile, &Context{Models: AllModels}); e != nil {
		log.Fatalln(e)
	}
}
