package commands

import (
	"fmt"
	"star-wars-client/services"

	"github.com/urfave/cli/v2"
)


var flags = []cli.Flag{
	&cli.StringFlag{
		Name: "search",
		Value: "",
	},
}

var StarshipCommand = cli.Command{
	Name: "starship",
	Flags: flags,
	Aliases: []string{"ships", "starships", "spacheships", "s"},
	Action: func(c *cli.Context) error {
		starships, err := services.GetAllStarships()


		if err != nil {
			return err
		}

		for _, sts := range starships {
			fmt.Println(sts.Stringify())
		}

		return nil
	},
}
