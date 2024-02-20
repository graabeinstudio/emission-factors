package cmd

import (
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
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		location, err := cmd.Flags().GetString("location")
		if err != nil {
			panic(err)
		}

		year, err := cmd.Flags().GetInt("year")
		if err != nil {
			panic(err)
		}

		electricityMix := internal.EmissionFactor{
			Type: internal.ElectricityMix,
			Year: year,
			Location: location,
			Factor: 19,
			Unit: "gram CO2e/kWh",
			Source:
				"https://www.nve.no/energi/energisystem/kraftproduksjon/hvor-kommer-stroemmen-fra/",
			Description: "Beregnet CO2 faktor for strømforbruk i Norge i 2022.",
		}

		fmt.Println(internal.PrettyPrint(electricityMix));
	},
}
