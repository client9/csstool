package cmd

import (
	"fmt"
	"os"

	// homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

// var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "css",
	Short: "tools to simplify, minify and reformat CSS files",
	Long: `Tools to simplify, minify and reformat CSS files.

The main command is 'css cut' which will remove unused CSS rules based on existing HTML files. For frameworks like bootstrap, the savings can be over 90%.

'css format' will unminify CSS files and indent them nicely.

'css minify' is a CSS whitespace minifier.  It does not do more advanced transformations altering the rules themselves. 

'css count' is mostly for debugging and identifying rarely used css rules.
`,
}

// Execute is the entry point for the CLI command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var (
	flagDebug bool
)

func init() {
	//cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVarP(&flagDebug, "debug", "d", false, "debug logging")

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.css.yaml)")
}

/*
// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".css" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".css")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
*/
