package services

import (
	"encoding/json"
	"fmt"
	"log"
	"star-wars-client/exceptions"
	"star-wars-client/models"
)



func GetPlanets() ([]models.Planet, error) {
	responses := []models.Planet{}

	page := 1

	for {
		planetResult := models.PlanetResult{}
		url := fmt.Sprintf("%s/%s?format=json&page=%d",URL, "planets",page)


		fmt.Println(url)

		resp, err := MakeGetRequest(url)


		if err == exceptions.NotFoundException{
			break
		}else if  err !=  nil{
			return nil, err
		}


		if errBody := json.Unmarshal(resp, &planetResult); errBody != nil {
			return nil, errBody
		}


		responses = append(responses, planetResult.Planets...)
		page++
	}


	return responses, nil


}

func SearchPlanet(name string) models.PlanetResult {
	url := fmt.Sprintf("%s/%s?format=json&search=%s", URL, "planets", name)

	var results models.PlanetResult

	resp, err := MakeGetRequest(url)

	if err != nil {
		log.Fatalln("err: ", err)
	}



	json.Unmarshal(resp, &results)

	return results

}

