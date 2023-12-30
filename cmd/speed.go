package cmd

import (
	"fmt"
	"log"

	"github.com/ddo/go-fast"
	"github.com/spf13/cobra"
)

// speedCmd represents the speed command
var speedCmd = &cobra.Command{
	Use:   "speed",
	Short: "Speed test",
	Long:  `Speed test is a tool to test your internet speed.`,
	Run:   runSpeedTest,
}

func runSpeedTest(cmd *cobra.Command, args []string) {
	fastCom := fast.New()

	err := fastCom.Init()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Download speed")
	urls, err := fastCom.GetUrls()
	if err != nil {
		log.Fatal(err)
	}

	if len(urls) == 0 {
		log.Fatal("No compatible servers found")
	}

	KbpsChan := make(chan float64)

	go func() {
		for Kbps := range KbpsChan {
			fmt.Printf("%.2f Kbps %.2f Mbps\n", Kbps, Kbps/1000)
		}

		fmt.Println("done")
	}()

	err = fastCom.Measure(urls, KbpsChan)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(speedCmd)
}
