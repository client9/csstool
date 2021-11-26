package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/lemondevxyz/csstool"
)

// minifyCmd represents the minify command
var minifyCmd = &cobra.Command{
	Use:   "minify",
	Short: "Minify a CSS file",
	Long: `Minify a CSS file
Currently only supports stdin and stdout.

This is shorthand for 'css format --indent=0 -semicolon=false ...'
`,
	Run: func(cmd *cobra.Command, args []string) {
		cssformat := csstool.NewCSSFormat(0, false, nil)
		cssformat.AlwaysSemicolon = false
		err := cssformat.Format(os.Stdin, os.Stdout)
		if err != nil {
			log.Printf("FAIL: %s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(minifyCmd)
}
