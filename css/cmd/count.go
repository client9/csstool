package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/client9/csstool"
	"github.com/mattn/go-zglob"
	"github.com/spf13/cobra"
)

// countCmd represents the count command
var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Print selector counts",
	Long: `Print element, class and id counts from HTML files

This is mostly useful for debugging "css cut" and/or finding rarely
used selectors.
`,
	Run: func(cmd *cobra.Command, args []string) {
		c := csstool.NewCSSCount()
		if flagHTML == "" {
			if flagDebug {
				log.Printf("reading from stdin")
			}
			err := c.Add(os.Stdin)
			if err != nil {
				log.Fatalf("FAIL: %s", err)
			}
		} else {
			if flagDebug {
				log.Printf("Using pattern %q", flagHTML)
			}
			files, err := zglob.Glob(flagHTML)
			if err != nil {
				log.Fatalf("FAIL: %s", err)
			}
			for _, f := range files {
				log.Printf("reading %s", f)
				r, err := os.Open(f)
				if err != nil {
					log.Fatalf("FAIL: %s", err)
				}
				err = c.Add(r)
				if err != nil {
					log.Fatalf("FAIL: %s", err)
				}
				r.Close()
			}
		}

		switch flagFormat {
		case "list":
			for _, val := range c.List() {
				fmt.Printf("%s\n", val)
			}
		case "csv":
			fmt.Printf("%s\n", strings.Join(c.List(), ","))
		case "count", "counts":
		default:
			log.Fatalf("Unknown format %q", flagFormat)
		}
	},
}
var (
	flagFormat string
)

func init() {
	rootCmd.AddCommand(countCmd)
	countCmd.Flags().StringVarP(&flagFormat, "format", "f", "list", "output format entation")
	countCmd.Flags().StringVarP(&flagHTML, "html", "", "", "glob pattern to find HTML files")
}
