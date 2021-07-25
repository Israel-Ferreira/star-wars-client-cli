package commands

import (
	"fmt"
	"star-wars-client/models"
	"star-wars-client/services"

	"github.com/urfave/cli/v2"
)

var PeopleCommand = cli.Command{
	Name:    "people",
	Usage:   "Get Star Wars Characters",
	Aliases: []string{"plp", "persons", "characters"},
	Flags:   Flags,

	Action: func(c *cli.Context) error {
		all := c.Bool("all")

		search := c.String("search")

		var results []models.Person

		if search != "" {
			fmt.Println("Function Not implemented")
		} else if all {
			resp, err := services.GetPersons()

			if err != nil {
				return err
			}

			results = append(results, resp...)

		}

		if len(results) > 0 {
			for _, person := range results {
				fmt.Println("Nome: ", person.Name)
			}
		}

		return nil
	},
}
