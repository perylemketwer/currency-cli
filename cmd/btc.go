/*
Copyright © 2023 Pery Lemke <pery.lemke@thoughtworks.com>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/perylemketwer/currency-cli/models"
	"github.com/spf13/cobra"
)

// btcCmd represents the btc command
var btcCmd = &cobra.Command{
	Use:   "btc",
	Short: "Cotação agora do Bitcoin.",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("https://economia.awesomeapi.com.br/json/last/BTC")
		if err != nil {
			log.Fatal(err)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		var price models.CoinToReal
		json.Unmarshal([]byte(body), &price)

		fmt.Printf("Cotação agora do Bitcoin: R$ %s", price.BitcoinToReal.Bid)
	},
}

func init() {
	rootCmd.AddCommand(btcCmd)
}
