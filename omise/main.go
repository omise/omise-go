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

func output(obj interface{}) (e error) {
	bytes := []byte{}

	switch {
	case !config.JSON && !config.IndentedJSON:
		// try stingers first
		if stringer, ok := obj.(ColorStringer); ok {
			bytes = []byte(stringer.ColorString())
			break
		} else if stringer, ok := obj.(fmt.Stringer); ok {
			bytes = []byte(stringer.String())
			break
		}

		fallthrough

	case config.IndentedJSON:
		bytes, e = json.MarshalIndent(obj, "", "  ")
	default: // config.JSON:
		bytes, e = json.Marshal(obj)
	}

	bytes = append(bytes, []byte("\n")...)
	_, e = os.Stdout.Write(bytes)
	return e
}
