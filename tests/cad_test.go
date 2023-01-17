package tests

import (
	"testing"

	"github.com/matryer/is"
	"github.com/perylemketwer/currency-cli/cmd"
	"github.com/spf13/cobra"
)

func TestCad(t *testing.T) {
	is := is.New(t)

	cad := &cobra.Command{Use: "cad", Run: cmd.CadCmd}

	err := cad.Execute()

	is.NoErr(err)
}
