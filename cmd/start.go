package cmd

import (
	"client-go/deployment"
	"client-go/service"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Create book server deployment and service",
	Long:  `Create book server deployment and service`,
	Run: func(cmd *cobra.Command, args []string) {
		deployment.Deploy()
		service.Serve()
	},
}
