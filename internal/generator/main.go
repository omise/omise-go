package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var Context = &struct {
	Models           []string
	SearchableModels []string
}{
	Models: []string{
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
		"Occurrence",
		"Recipient",
		"Refund",
		"Schedule",
		"Token",
		"Transaction",
		"Transfer",
	},
	SearchableModels: []string{
		"Charge",
		"Dispute",
		"Recipient",
		"Customer",
		"Refund",
		"Transfer",
		"Link",
	},
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

	if e := tmpl.ExecuteTemplate(outfile, templateFile, Context); e != nil {
		log.Fatalln(e)
	}
}
