package services

import (
	"encoding/json"
	"fmt"
	"star-wars-client/exceptions"
	"star-wars-client/models"
)

func GetPersons() ([]models.Person, error)  {
	results := []models.Person{}

	var page uint = 1

	for  {
		personsResult := models.PersonResult{}
		url := fmt.Sprintf("%s/%s?page=%d&format=json", URL, "people", page)

		resp, err := MakeGetRequest(url)

		if err == exceptions.NotFoundException {
			break
		}


		if err != nil && err != exceptions.NotFoundException {
			return nil, err
		}
	



		if errUnMarshall := json.Unmarshal(resp, &personsResult); errUnMarshall != nil {
			return nil, errUnMarshall
		}


		results = append(results, personsResult.Persons...)
		page++
	}



	return results, nil
}
