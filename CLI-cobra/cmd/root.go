/*
Copyright © 2025 Guilli

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var dataFile string
var ignoreConfig bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli-cobra",
	Short: "This app is a simple CLI application built with Cobra",
	Long: `cli-cobra uses Cobra to build a simple and powerful
command line interface application. You can extend this application
by adding more commands and features as needed. Great for learning
how to build CLI apps in Go!`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	home, err := homedir.Dir()
	if err != nil {
		log.Fatalf("Could not get current working directory: %v", err)
	}

	dataFile = filepath.Join(home, ".todo.json")

	rootCmd.PersistentFlags().StringVar(&dataFile, "datafile", dataFile, "data file to store todos")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli-cobra.yaml)")
	rootCmd.PersistentFlags().BoolVar(&ignoreConfig, "ignore-config", false, "ignore configuration file and use default settings")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if ignoreConfig || os.Getenv("IGNORE_CONFIG") == "1" {
		fmt.Fprintln(os.Stderr, "Ignoring config file and environment variables, using default settings")
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		dataFile = filepath.Join(home, ".todo.json")
		return
	}

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cli-cobra")
		viper.SetDefault("datafile", filepath.Join(home, ".todo.json"))
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("new")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		dataFile = viper.GetString("datafile")
		fmt.Fprintln(os.Stderr, "Using data file:", dataFile)
	} else {
		fmt.Fprintln(os.Stderr, "No config file found, using default data file.")
	}
}
