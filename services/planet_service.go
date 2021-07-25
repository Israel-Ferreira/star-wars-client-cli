package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"star-wars-client/models"
)



func GetPlanets() ([]models.Planet) {
	responses := []models.Planet{}

	totalPages := 10

	for page := 1; page <= totalPages; page++ {

		planetResult := models.PlanetResult{}

		url := fmt.Sprintf("%s&page=%d",URL, page)
		resp, errorResp := http.Get(url)

		if errorResp != nil {
			log.Fatalln("Err: ", errorResp)
		}

		body := deserializeJson(*resp)

		json.Unmarshal(body, &planetResult)

		responses = append(responses, planetResult.Planets...)
	}

	return responses

}

func SearchPlanet(name string) models.PlanetResult {
	url := fmt.Sprintf("%s&search=%s", URL, name)

	var results models.PlanetResult

	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln("err: ", err)
	}

	body := deserializeJson(*resp)
	

	json.Unmarshal(body, &results)

	return results

}

func deserializeJson(resp http.Response) []byte {

	defer resp.Body.Close()

	body, errBody := ioutil.ReadAll(resp.Body)

	if errBody != nil {
		log.Fatalln("Erro ao parsear o json: ", errBody)
	}


	return body
}
