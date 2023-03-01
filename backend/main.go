package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("./backend/artists.txt")
	if err != nil {
		log.Fatalln(err)
	}

	app := application{}

	err = json.Unmarshal(data, &app.artists)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(app.artists[0])
}
