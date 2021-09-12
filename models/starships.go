package models

import "fmt"

type Starship struct {
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	Length       string `json:"length"`
	Crew         string `json:"crew"`
	Category     string `json:"starship_class"`
}

func (s Starship) Stringify() string {
	return fmt.Sprintf(
		"Nome: %s, Fabricante: %s, Category: %s, Crew: %s",
		s.Name,
		s.Manufacturer,
		s.Category,
		s.Crew,
	)
}
