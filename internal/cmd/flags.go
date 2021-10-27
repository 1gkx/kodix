package cmd

import (
	"github.com/urfave/cli/v2"
	"kodix/internal/config"
)

var FlagPort = &cli.IntFlag{
	Name: "port",
	Aliases: []string{"p"},
	Required: true,
	Value: 3000,
	Usage: "use custom port",
	Destination: &config.Port,
}
