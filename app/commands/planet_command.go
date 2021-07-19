package commands

import (
	"fmt"
	"star-wars-client/services"

	"github.com/urfave/cli/v2"
)

var PlanetFlags = []cli.Flag{
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
	Flags:   PlanetFlags,
	Action: func(c *cli.Context) error {
		all := c.Bool("all")

		search := c.String("search")

		if !all && search != "" {
			resp := services.SearchPlanet(search)

			for _, result := range resp.Planets {
				fmt.Printf("Nome: %s, Climate: %s,  \n", result.Name, result.Climate)
			}

		} else if all {
			resp := services.GetPlanets()

			for _, result := range resp {
				fmt.Printf("Nome: %s, Climate: %s,  \n", result.Name, result.Climate)
			}
		}

		return nil
	},
}
