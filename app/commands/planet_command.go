package commands

import (
	"fmt"
	"star-wars-client/services"
	"star-wars-client/utils"

	"github.com/urfave/cli/v2"
)

var Flags = []cli.Flag{
	&cli.BoolFlag{
		Name:  "all",
		Value: true,
	},
	&cli.StringFlag{
		Name: "search",
	},
}

var PlanetCommand = cli.Command{
	Name:    "planets",
	Aliases: []string{"pln", "p"},
	Flags:   Flags,
	Action: func(c *cli.Context) error {
		all := c.Bool("all")

		search := c.String("search")

		if !all && search != "" {
			resp := services.SearchPlanet(search)

			for _, result := range resp.Planets {
				fmt.Printf("Nome: %s, Climate: %s,  \n", result.Name, result.Climate)
			}

		} else if all {
			resp, err := services.GetPlanets()

			utils.HandleErrors(err, "Erro ao obter todos os personagens")

			for _, result := range resp {
				fmt.Printf("Nome: %s, Climate: %s,  \n", result.Name, result.Climate)
			}
		}

		return nil
	},
}
