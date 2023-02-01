package tests

import (
	"testing"

	"github.com/matryer/is"
	"github.com/perylemketwer/currency-cli/cmd"
	"github.com/spf13/cobra"
)

func TestEth(t *testing.T) {
	is := is.New(t)

	eth := &cobra.Command{Use: "eth", Run: cmd.EthCmd}

	err := eth.Execute()

	is.NoErr(err)
}
