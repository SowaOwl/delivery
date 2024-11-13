package debug

import (
	"encoding/json"
	"fmt"
	"log"
)

func JsonPrint(object interface{}) {
	jsonResponse, err := json.Marshal(object)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonResponse))
}
