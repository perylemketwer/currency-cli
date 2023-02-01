/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
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

func EthCmd(cmd *cobra.Command, args []string) {
	resp, err := http.Get("https://economia.awesomeapi.com.br/json/last/ETH")
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

	fmt.Printf("Cotação agora do Ethereum: R$ %s", price.EthereumToReal.Bid)
}

// ethCmd represents the eth command
var ethCmd = &cobra.Command{
	Use:   "eth",
	Short: "Cotação agora do Ethereum.",
	Run:   EthCmd,
}

func init() {
	rootCmd.AddCommand(ethCmd)
}
