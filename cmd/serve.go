package cmd

import (
	"blog/pkg/bootstrap"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve app on dev server",
	Long:  `The application will be served on the dev server on the port specified in config.yml`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	bootstrap.Serve()
}
