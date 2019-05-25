package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var netClient = &http.Client{
	Timeout: time.Second * 20,
}

// GetCmd represents the get command
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Query array resources",
	Long:  `Get information about specific array resources`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
	},
}
