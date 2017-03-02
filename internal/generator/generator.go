package generator

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

func Main(name string) {
	var (
		templateFile = name + ".tmpl"
		outputFile   = name + ".go"
	)

	tmpl := template.New("root").
		Funcs(map[string]interface{}{
			"toLower": strings.ToLower,
		})

	tmpl = template.Must(tmpl.ParseFiles(templateFile))

	outfile, e := os.Create(outputFile)
	if e != nil {
		log.Fatalln(e)
	}

	if e := tmpl.ExecuteTemplate(outfile, templateFile, &Context{Models: AllModels}); e != nil {
		log.Fatalln(e)
	}
}
