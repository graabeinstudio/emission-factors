package cmd

import (
	"errors"
	"fmt"

	"github.com/graabeinstudio/emission-factors/internal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(electricityMixCommand)
	electricityMixCommand.PersistentFlags().String("location", "norway", "location for electricity mix emission factor")
	electricityMixCommand.PersistentFlags().Int("year", 2022, "year for electricity mix emission factor")
}

var electricityMixCommand = &cobra.Command{
	Use:   "electricity-mix location year",
	Short: "Run script for generating the electricity mix emission factor",
	Long: `Run script for generating the electricity mix emission factor.
			Currently only supports the electricity mix for Norway in 2022.`,
	Run: func(cmd *cobra.Command, args []string) {
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

		if (year != 2022 || location != internal.NORWAY) {
			panic(errors.New("only supported year and location is '2022' and 'norway'"));
		} 
		
		fmt.Println(internal.PrettyPrint(internal.ElectricityMixNorway2022));
	},
}
