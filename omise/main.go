package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	configure(OmiseCmd.PersistentFlags())
	if e := OmiseCmd.Execute(); e != nil {
		log.Fatalln(e)
	}
}

func checkArgs(args []string, argnames ...string) error {
	if len(args) == len(argnames) {
		return nil
	}

	return ErrMissingArg(argnames[len(args)])
}

func output(obj interface{}) error {
	// try stingers first
	if stringer, ok := obj.(ColorStringer); ok {
		_, e := os.Stdout.Write([]byte(stringer.ColorString() + "\n"))
		return e
	}
	if stringer, ok := obj.(fmt.Stringer); ok {
		_, e := os.Stdout.Write([]byte(stringer.String() + "\n"))
		return e
	}

	// fallback indented json
	bytes, e := json.MarshalIndent(obj, "", "  ")
	if e != nil {
		return e
	}

	bytes = append(bytes, []byte("\n")...)
	_, e = os.Stdout.Write(bytes)
	return e
}
