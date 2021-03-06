package cmd

import (
	"fmt"
	"time"

	magichome "github.com/apoclyps/magic-home/pkg"
	"github.com/spf13/cobra"
)

// discoverCmd represents the discover command
var discoverCmd = &cobra.Command{
	Use:   "discover",
	Short: "Discover Magic Home Devices",
	Long: `Discover Magic Home Devices on the local area network

	Defaults to searching on '255.255.255.255' but can be provided 
	with a specific broadcast address.
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		broadcastAddr, err := cmd.Flags().GetString("broadcast")
		if err != nil {
			fmt.Println(err)
			return err
		}
		if broadcastAddr == "" {
			broadcastAddr = magichome.DEFAULT_BROADCAST_ADDR
		}
		return discover(broadcastAddr)
	},
}

func discover(broadcastAddr string) error {
	fmt.Print("Discovering ")

	go func() {
		for {
			fmt.Print(".")
			time.Sleep(100 * time.Millisecond)
		}
	}()

	devices, err := magichome.Discover(magichome.DiscoverOptions{
		BroadcastAddr: broadcastAddr,
		Timeout:       1,
	})
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	fmt.Print("\n\nDiscovered the following devices:\n\n")

	fmt.Println("IP         \t| ID         \t| Model")
	fmt.Println("-----------------------------------")
	for _, device := range *devices {
		fmt.Printf("%s\t| %s\t| %s\n", device.IP, device.ID, device.Model)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(discoverCmd)
	discoverCmd.Flags().StringP("broadcast", "b", "255.255.255.255", "Specify a broadcast address to use for discovering devices e.g. 255.255.255.255")
}
