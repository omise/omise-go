package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	patterns := os.Args[1:]
	for _, pattern := range patterns {
		filenames, err := filepath.Glob(pattern)
		if err != nil {
			log.Fatal(err)
		}

		for _, filename := range filenames {
			log.Println("processing:", filename)
			if err := refactorFile(filename); err != nil {
				log.Fatalln(err)
			}
		}
	}
}

func printHelp() {
	fmt.Fprintf(os.Stderr, "Usage: %s (filename)\n", os.Args[0])
}

func refactorFile(filename string) error {
	fset, v := token.NewFileSet(), &visitor{}
	filenode, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return err
	} else if ast.Walk(v, filenode); v.err != nil {
		return v.err
	}

	if file, err := os.Create(filename); err != nil {
		return err
	} else {
		defer file.Close()
		return printer.Fprint(file, fset, filenode)
	}
}
