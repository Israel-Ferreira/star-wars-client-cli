package models

type Result struct {
	Count int `json:"count"`
	Next string `json:"next"`
}


type PersonResult struct{
	Result
	Persons []Person `json:"results"`
}