package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"star-wars-client/app"
	"star-wars-client/models"
)

const URL = "https://swapi.dev/api/people?format=json"

func getPersonsRequest(page int, done chan string) {
	novaUrl := fmt.Sprintf("%s&page=%d", URL, page)

	res, err1 := http.Get(novaUrl)

	if err1 != nil {
		log.Fatalln(err1)
	}

	body, err2 := ioutil.ReadAll(res.Body)

	if err2 != nil {
		log.Fatalln(err2)
	}

	httpResp := models.PersonResult{}

	json.Unmarshal(body, &httpResp)

	defer fmt.Println(httpResp.Persons)

	done <- "Teste Channels"
}

func main() {
	app := app.CriarCli()

	err := app.Run(os.Args)

	if err !=  nil {
		log.Fatalln("Error: ", err)
	}
}
