package cmd

import (
	"fmt"
	"log"

	"github.com/brudnyhenry/harrayp/config"
	"github.com/spf13/cobra"
)

var volumeType string

// volumesCmd represents the volumes command
var volumesCmd = &cobra.Command{
	Use:   "volumes",
	Short: "Get information about volumes",
	Long:  `Get information about volumes`,
	Run: func(cmd *cobra.Command, args []string) {

		a := HpArray{
			URL:      config.URL,
			user:     config.Login,
			password: config.Password,
			Client:   netClient,
		}

		volumes, err := a.ShowVolumes(volumeType)
		if err != nil {
			log.Fatal("Error while fetching volumes")
		}
		for _, volume := range volumes {
			fmt.Println(volume)
		}
	},
}

func init() {
	GetCmd.AddCommand(volumesCmd)
	volumesCmd.Flags().StringVarP(&volumeType, "type", "t", "all", "Volume type one of options: all|snapshot|volume")

}
