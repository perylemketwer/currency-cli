package tests

import (
	"testing"

	"github.com/matryer/is"
	"github.com/perylemketwer/currency-cli/cmd"
	"github.com/spf13/cobra"
)

func TestGbp(t *testing.T) {
	is := is.New(t)

	gbp := &cobra.Command{Use: "gbp", Run: cmd.GbpCmd}

	err := gbp.Execute()

	is.NoErr(err)
}
