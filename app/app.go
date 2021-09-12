package app

import (
	"fmt"
	"star-wars-client/app/commands"

	"github.com/urfave/cli/v2"
)




func CriarCli() *cli.App {
	app := &cli.App{
		Name: "star-wars-cli",
		Usage: "Star-Wars Cli",
		Action: func(c *cli.Context) error{
			fmt.Println("May The Force Be With You")

			return nil
		},

		Commands: []*cli.Command{
			&commands.PlanetCommand,
			&commands.PeopleCommand,
			&commands.StarshipCommand,
		},
	}


	return app
}


