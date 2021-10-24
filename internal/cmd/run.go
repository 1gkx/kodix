package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/urfave/cli"

	"kodix/internal/router"
	"kodix/internal/store"
)

var Start = cli.Command{
	Name:        "start",
	Usage:       "Start web server",
	Description: `Description`,
	Action:      runWeb,
}

func runWeb(c *cli.Context) {

	db, err := store.InitStore()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := router.NewRouter(db)

	fmt.Printf("Server start at %s port\n", ":8000")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
}
