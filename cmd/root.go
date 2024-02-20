package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "emission-factor",
	Short: "Run script to generate emission factor",
	Long: `Run script to generate emission factor. Currently only supports the electricity mix for Norway in 2022.`,
  }
  
  func Execute() {
	if err := rootCmd.Execute(); err != nil {
	  fmt.Fprintln(os.Stderr, err)
	  os.Exit(1)
	}
  }