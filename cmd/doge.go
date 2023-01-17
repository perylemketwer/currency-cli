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

// The function to doge command
func DogeCmd(cmd *cobra.Command, args []string) {
	resp, err := http.Get("https://economia.awesomeapi.com.br/json/last/DOGE")
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

	fmt.Printf("Cotação agora do Dogecoin: R$ %s", price.DogecoinToReal.Bid)
}

// dogeCmd represents the doge command
var dogeCmd = &cobra.Command{
	Use:   "doge",
	Short: "Cotação agora do Dogecoin.",
	Run:   DogeCmd,
}

func init() {
	rootCmd.AddCommand(dogeCmd)
}
