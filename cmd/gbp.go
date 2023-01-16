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

// gbpCmd represents the gbp command
var gbpCmd = &cobra.Command{
	Use:   "gbp",
	Short: "Cotação agora da Libra Esterlina",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("https://economia.awesomeapi.com.br/json/last/GBP")
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

		fmt.Printf("Cotação agora da Libra Esterlina: R$ %s", price.PoundToReal.Bid)
	},
}

func init() {
	rootCmd.AddCommand(gbpCmd)
}
