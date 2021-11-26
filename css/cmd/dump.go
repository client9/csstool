package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/lemondevxyz/csstool"
)

// formatCmd represents the format command
var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump grammar/token",
	Long: `Print CSS grammar and tokens.  Only useful for debugging.
	Currently only supports stdin and stdout.
`,
	Run: func(cmd *cobra.Command, args []string) {
		err := csstool.Dump(os.Stdin, os.Stdout)
		if err != nil {
			log.Printf("FAIL: %s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(dumpCmd)
}
