package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/client9/csstool"
)

// formatCmd represents the format command
var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "Reformat a CSS file",
	Long: `Reformat a CSS file.
	Currently only supports stdin and stdout.
	--indent=0 will mostly remove whitespace and newlines
`,
	Run: func(cmd *cobra.Command, args []string) {
		if flagTab {
			flagIndent = 1
		}
		cssformat := csstool.NewCSSFormat(flagIndent, flagTab, nil)
		cssformat.AlwaysSemicolon = flagSemicolon
		err := cssformat.Format(os.Stdin, os.Stdout)
		if err != nil {
			log.Printf("FAIL: %s", err)
		}
	},
}

var (
	flagTab       bool
	flagIndent    int
	flagSemicolon bool
)

func init() {
	rootCmd.AddCommand(formatCmd)
	formatCmd.Flags().BoolVarP(&flagTab, "tab", "t", false, "use tabs for indentation")
	formatCmd.Flags().IntVarP(&flagIndent, "indent", "i", 2, "spaces for indentation")
	formatCmd.Flags().BoolVarP(&flagSemicolon, "semicolon", "", true, "always end rule with semicolon, even if not needed")
}
