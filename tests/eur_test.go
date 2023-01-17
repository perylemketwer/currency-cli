package tests

import (
	"testing"

	"github.com/matryer/is"
	"github.com/perylemketwer/currency-cli/cmd"
	"github.com/spf13/cobra"
)

func TestEur(t *testing.T) {
	is := is.New(t)

	eur := &cobra.Command{Use: "eur", Run: cmd.EurCmd}

	err := eur.Execute()

	is.NoErr(err)
}
