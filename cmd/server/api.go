package server

import (
	"net/http"
	"order-svc/internal/router"

	"github.com/spf13/cobra"
)

func NewApi() *cobra.Command {
	return &cobra.Command{
		Use: "api",
		Run: func(cmd *cobra.Command, args []string) {
			router := router.NewRouter()
			s := &http.Server{
				Addr:    "127.0.0.1:8080",
				Handler: router,
			}
			s.ListenAndServe()
		},
	}
}
