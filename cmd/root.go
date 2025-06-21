package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "eg-lm-auto",
	Short: "Endgate Listmonk Automation Layer",
	Long:  `CLI tool for managing and automating Listmonk workflows.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
