package commands

import (
	"fmt"
	"star-wars-client/models"
	"star-wars-client/services"
	"star-wars-client/utils"

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

			utils.HandleErrors(err, "Erro ao Obter os Personagens na API")
	
			results = append(results, resp...)

			for _, person := range results {
				fmt.Println("Nome: ", person.Name)
			}

		}


		return nil
	},
}
