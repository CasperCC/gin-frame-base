package cmd

import (
	"gin-frame-base/app/route"
	_ "gin-frame-base/internal/bootstrap"
	"gin-frame-base/internal/server"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start server",
	Long:  "start server",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func run() {
	http := server.New()
	http.GenRouter(route.New())
	http.Run()
}
