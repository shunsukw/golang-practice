package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/shunsukw/golang-practice/dino/dinowebportal"
)

type configuration struct {
	ServerAddress      string `json:"webserver"`
	DatabaseType       uint8  `json:"databasetype"`
	DatabaseConnection string `json:"dbconnection"`
	FrontEnd           string `json:"frontend"`
}

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	config := new(configuration)
	json.NewDecoder(file).Decode(config)
	log.Println("Starting on web server on address", config.ServerAddress)
	dinowebportal.RunWebPortal(config.DatabaseType, config.ServerAddress, config.DatabaseConnection, config.FrontEnd)
}
