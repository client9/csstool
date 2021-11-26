package cmd

import (
	"log"
	"os"

	"github.com/mattn/go-zglob"
	"github.com/spf13/cobra"

	"github.com/lemondevxyz/csstool"
)

// cutCmd represents the cut command
var cutCmd = &cobra.Command{
	Use:   "cut",
	Short: "Remove unneeded CSS rules based on HTML usage",
	Long: `Removes unneeded CSS rules based on HTML usage.
	
For use with Hugo:
    css cut --html 'public/**/*.html' < bootstrap.min.css \
         > static/bootstrap-csscut.min.css
`,

	Run: func(cmd *cobra.Command, args []string) {
		c := csstool.NewCSSCount()
		if flagDebug {
			log.Printf("using pattern %q", flagHTML)
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

		// now get CSS file
		m := csstool.NewTagMatcher(c.List())
		for _, key := range flagKeep {
			m.AddSelector(key)
		}
		for _, key := range flagRemove {
			m.RemoveSelector(key)
		}
		cf := csstool.NewCSSFormat(0, false, m)
		cf.RemoveAtRule = flagRemoveAtRule
		cf.RemoveSourceMap = flagRemoveSourceMap
		cf.Debug = flagDebug
		err = cf.Format(os.Stdin, os.Stdout)
		if err != nil {
			log.Printf("FAIL: %s", err)
		}
	},
}

var (
	flagHTML            string
	flagKeep            []string
	flagRemove          []string
	flagRemoveAtRule    []string
	flagRemoveSourceMap bool
)

func init() {
	rootCmd.AddCommand(cutCmd)
	cutCmd.Flags().StringVarP(&flagHTML, "html", "", "", "glob pattern to find HTML files")
	cutCmd.Flags().StringSliceVarP(&flagKeep, "keep", "", nil, "csv list of selectors to keep")
	cutCmd.Flags().StringSliceVarP(&flagRemove, "remove", "", nil, "csv list of selectors to remove")
	cutCmd.Flags().StringSliceVarP(&flagRemoveAtRule, "remove-at-rule", "", nil, "csv list of at-rules to remove. for example '@media print'")
	cutCmd.Flags().BoolVarP(&flagRemoveSourceMap, "remove-source-map", "", false, "remove comments which contain a source mapping url")
	cutCmd.MarkFlagRequired("html")
}
