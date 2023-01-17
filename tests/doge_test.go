package tests

import (
	"testing"

	"github.com/matryer/is"
	"github.com/perylemketwer/currency-cli/cmd"
	"github.com/spf13/cobra"
)

func TestDoge(t *testing.T) {
	is := is.New(t)

	doge := &cobra.Command{Use: "doge", Run: cmd.DogeCmd}

	err := doge.Execute()

	is.NoErr(err)
}
