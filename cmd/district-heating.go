package cmd

import (
	"errors"
	"fmt"

	"github.com/graabeinstudio/emission-factors/internal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(districtHeatingCommand)
}

var districtHeatingCommand = &cobra.Command{
	Use:   "district-heating",
	Short: "Run script for generating the district heating emission factor",
	Long: `Run script for generating the district heating emission factor.
			Currently only supports the district heating for Norway in 2022.`,
	Run: func(cmd *cobra.Command, args []string) {
		year, location := getYearAndLocation(cmd);

		if (year != 2022 || location != internal.NORWAY) {
			panic(errors.New("only supported year and location is '2022' and 'norway'"));
		} 

		fmt.Println(internal.PrettyPrint(internal.DistrictHeatingNorway2022));
	},
}
