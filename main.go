package main

import (
	"order-svc/cmd/server"

	"github.com/spf13/cobra"
)

func main() {
	cmd := cobra.Command{}
	api := server.NewApi()
	cmd.AddCommand(api)
	cmd.Execute()
}
