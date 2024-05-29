package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "prometheus-demo",
	Short: "prometheus-server",
	Long:  `prometheus-server`,
}

func Execute() {
	fmt.Println("Inside root.go")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
