package cmd

import (
	"fmt"
	"log"

	"github.com/brudnyhenry/harrayp/config"

	"github.com/spf13/cobra"
)

// hostCmd represents the host command
var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "Prints out list of hosts",
	Long:  `Prints to the stdout list of hosts defined in the array configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		a := HpArray{
			URL:      config.URL,
			user:     config.Login,
			password: config.Password,
			Client:   netClient,
		}
		hosts, err := a.ShowHosts()
		if err != nil {
			log.Fatal("Error while fetching hosts")
		}
		for _, host := range hosts {
			fmt.Println(host)
		}
	},
}

func init() {
	GetCmd.AddCommand(hostCmd)
}
