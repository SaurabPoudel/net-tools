package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

type Ip struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Timezone string `json:"timezone"`
	Postal   string `json:"postal"`
}

var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace the IP",
	Long:  `Trace the IP`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				showData(ip)
			}
		} else {
			fmt.Println("Please provide IP to trace")
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)
}

func showData(ip string) {
	url := "http://ipinfo.io/" + ip + "/geo"
	resByte := getData(url)
	data := Ip{}
	err := json.Unmarshal(resByte, &data)
	if err != nil {
		log.Println("Unable to unmarshall the response")
	}
	fmt.Println("DATA FOUND : ")
	fmt.Printf("IP : %s\nCITY:%s\nREGION:%s\nCOUNTRY:%s\nLOCATION:%s\nTIMEZONE:%s\nPOSTAL%s\n", data.IP, data.City, data.Region, data.Country, data.Loc, data.Timezone, data.Postal)
}

func getData(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		log.Println("404 not found")
	}

	resByte, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("502 Bad Gateway unable to read the response")
	}

	return resByte
}
