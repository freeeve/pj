package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	m := map[string]interface{}{}
	err := json.NewDecoder(os.Stdin).Decode(&m)
	if err != nil {
		log.Fatal(err)
	}
	buf, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(buf)
	os.Stdout.Write([]byte("\n"))
	os.Stdout.Close()
}
