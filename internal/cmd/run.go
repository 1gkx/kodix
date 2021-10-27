package cmd

import (
	"fmt"
	"net/http"

	"github.com/urfave/cli/v2"

	"kodix/internal/config"
	"kodix/internal/router"
	"kodix/internal/store"
)

var Start = cli.Command{
	Name:        "start",
	Usage:       "Start web server",
	Description: `Description`,
	Flags:       []cli.Flag{FlagPort},
	Action:      runWeb,
}

func runWeb(c *cli.Context) error {

	db, err := store.InitStore()
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}
	defer db.Close()

	fmt.Printf("Server start at %d port\n", config.Port)
	if err := http.ListenAndServe(
		fmt.Sprintf(":%d", config.Port),
		router.NewRouter(db),
	); err != nil {
		return cli.Exit(err.Error(), 1)
	}
	return nil
}
