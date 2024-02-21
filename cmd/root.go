package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/graabeinstudio/emission-factors/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "emission-factor",
	Short: "Run script to generate emission factor",
	Long:  `Run script to generate emission factor. Currently only supports the electricity mix for Norway in 2022.`,
	Run: func(cmd *cobra.Command, args []string) {
		year, location := getYearAndLocation(cmd);

		if (year != 2022 || location != internal.NORWAY) {
			panic(errors.New("only supported year and location is '2022' and 'norway'"));
		} 
		
		fmt.Println(internal.PrettyPrint(internal.EmissionFactorsNorway2022));
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("location", "norway", "location for emission factor")
	rootCmd.PersistentFlags().Int("year", 2022, "year for emission factor")
}

func getYearAndLocation(cmd *cobra.Command) (int, internal.Location) {
	locationAsString, err := cmd.Flags().GetString("location")
	if err != nil {
		panic(err)
	}

	location, err := internal.ToLocation(locationAsString)
	if err != nil {
		panic(err)
	}

	year, err := cmd.Flags().GetInt("year")
	if err != nil {
		panic(err)
	}

	return year, location
}
