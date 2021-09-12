package services

import (
	"encoding/json"
	"fmt"
	"star-wars-client/exceptions"
	"star-wars-client/models"
)

func GetAllStarships() ([]models.Starship, error) {
	responses := []models.Starship{}

	page := 1

	for {
		planetResult := models.StarshipResult{}
		url := fmt.Sprintf("%s/%s?format=json&page=%d", URL, "starships", page)


		resp, err := MakeGetRequest(url)

		if err == exceptions.NotFoundException {
			break
		} else if err != nil {
			return nil, err
		}

		if errBody := json.Unmarshal(resp, &planetResult); errBody != nil {
			return nil, errBody
		}

		responses = append(responses, planetResult.Starships...)
		page++
	}

	return responses, nil
}
