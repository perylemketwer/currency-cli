package tests

import (
	"testing"

	"github.com/matryer/is"
	"github.com/perylemketwer/currency-cli/cmd"
	"github.com/spf13/cobra"
)

func TestUsd(t *testing.T) {
	is := is.New(t)

	usd := &cobra.Command{Use: "usd", Run: cmd.UsdCmd}

	err := usd.Execute()

	is.NoErr(err)
}
