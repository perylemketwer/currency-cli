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

// The function to eur command
func EurCmd(cmd *cobra.Command, args []string) {
	resp, err := http.Get("https://economia.awesomeapi.com.br/json/last/EUR")
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

	fmt.Printf("Cotação agora do Euro: R$ %s", price.EuroToReal.Bid)
}

// eurCmd represents the eur command
var eurCmd = &cobra.Command{
	Use:   "eur",
	Short: "Cotação agora do Euro.",
	Run:   EurCmd,
}

func init() {
	rootCmd.AddCommand(eurCmd)
}
