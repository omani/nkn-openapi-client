package commands

import (
	"github.com/spf13/cobra"
)

// Globals
var (
	url string
)

var rootCmd = &cobra.Command{
	Use:     "nkn-openapi-client",
	Version: "1.0",
	Short:   "nkn-openapi-client - A cli tool for the NKN OpenAPI",
}

// RootCmd function  
func RootCmd() *cobra.Command {
	return rootCmd
}

// Execute function  
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

// init function  
func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().StringVar(&url, "url", "", "Use a different NKN OpenAPI URL")
}
