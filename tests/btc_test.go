package tests

import (
	"testing"

	"github.com/matryer/is"
	"github.com/perylemketwer/currency-cli/cmd"
	"github.com/spf13/cobra"
)

func TestBtc(t *testing.T) {
	is := is.New(t)

	btc := &cobra.Command{Use: "btc", Run: cmd.BtcCmd}

	err := btc.Execute()

	is.NoErr(err)
}
