package services

import (
	"io/ioutil"
	"net/http"
	"star-wars-client/exceptions"
)


func MakeGetRequest(url string) ([]byte, error) {

	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}


	defer res.Body.Close()


	if res.StatusCode == 400  {
		return nil, exceptions.BadRequestException
	}else if res.StatusCode == 500 {
		return nil, exceptions.InternalServerError
	}else if res.StatusCode == 404 {
		return nil, exceptions.NotFoundException
	}




	body, errBody := ioutil.ReadAll(res.Body)

	if errBody != nil {
		return nil, errBody
	}


	return body, nil
}