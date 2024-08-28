package src

import (
	"encoding/json"
	"log"
)

func CustomLog(d interface{}) {
	result, err := json.Marshal(d)

	if err != nil {
		log.Println(err)
	}

	log.Println(string(result))
}
