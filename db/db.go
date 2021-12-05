package db

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type People struct {
	Name   string
	Amount string
}

var Peoples []People

func init() {
	content, err := ioutil.ReadFile("./db/data.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(content, &Peoples)
	if err != nil {
		log.Fatal(err)
	}
}
