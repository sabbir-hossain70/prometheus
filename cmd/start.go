package cmd

import (
	"fmt"
	"github.com/sabbir-hossain70/prometheus/api"
	"github.com/spf13/cobra"
)

var Port string

var (
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start the server",
		Long:  `Start the server`,
		Run: func(cmd *cobra.Command, args []string) {
			api.RunServer(Port)
		},
	}
)

func init() {
	fmt.Println("Inside start.go")
	startCmd.Flags().StringVarP(&Port, "port", "p", "8081", "Port to listen on")
	rootCmd.AddCommand(startCmd)
	println("Port: ", Port)
}
