package factory

import "github.com/urfave/cli/v2"

type ActionService interface {
	Command() *cli.Command
	Do(c *cli.Context) error
}
