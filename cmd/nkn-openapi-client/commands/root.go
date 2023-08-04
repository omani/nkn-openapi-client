package commands

import (
	"net/url"

	client "github.com/omani/nkn-openapi-client"
	"github.com/spf13/cobra"
)

// Globals
var (
	apiurl string
	c      client.Client
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
	cobra.OnInitialize(initialize)

	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().StringVar(&apiurl, "url", "https://openapi.nkn.org/api/v1/", "Use a different NKN OpenAPI URL")
}

func initialize() {
	c = client.New()
	uri, err := url.Parse(apiurl)
	if err != nil {
		panic(err)
	}
	c.SetAddress(uri.String())
}
