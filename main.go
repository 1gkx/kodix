package main

import (
	"log"
	"os"

	"github.com/urfave/cli"

	"kodix/internal/cmd"
)

func main() {
	app := cli.NewApp()
	app.Name = "Kodix"
	app.Usage = "Example project for interview Kodix Automotive"
	app.Version = "0.1.0"
	app.Commands = []cli.Command{cmd.Start}
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}
}
