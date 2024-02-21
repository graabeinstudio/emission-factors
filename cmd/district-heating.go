package cmd

import (
	"errors"
	"fmt"
	"math"

	"github.com/graabeinstudio/emission-factors/internal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(districtHeatingCommand)
	districtHeatingCommand.PersistentFlags().String("location", "norway", "location for district heating mission factor")
	districtHeatingCommand.PersistentFlags().Int("year", 2022, "year for district heating emission factor")
}

var districtHeatingCommand = &cobra.Command{
	Use:   "district-heating location year",
	Short: "Run script for generating the district heating emission factor",
	Long: `Run script for generating the district heating emission factor.
			Currently only supports the district heating for Norway in 2022.`,
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

		/*
			Regnes ut basert på nasjonal fordeling oppgitt på https://www.fjernkontrollen.no/ for 2022,
			og med utslippsfaktorer fra https://www.fjernkontrollen.no/uploaded/files/2020_06_01_klimaregnskap_for_fjernvarme_2020.pdf
		*/
		recycledHeatFactor := 0.0
		fossilOilFactor := 286.0
		fossilGasFactor := 242.0
		bioEnergyFactor := 11.4
		factor := 
			0.479 * recycledHeatFactor + 
			0.012 * fossilOilFactor + 
			0.022 * fossilGasFactor + 
			0.072 * float64(internal.ElectricityMixNorway2022.Factor) +
			0.321 * bioEnergyFactor + // vet ikke hvilken kilde, så tar snittet av alle utslippsfaktorene
			0.094 * float64(internal.ElectricityMixNorway2022.Factor) // vet ikke hvilken kilde, så må gå ut fra el-varmepumpe


		districtHeating := internal.EmissionFactor{
			Type: internal.DistrictHeating,
			Year: year,
			Location: location,
			Factor: int(math.Ceil(factor)),
			Unit: "gram CO2e/kWh",
			Sources:
				[]string{"https://www.fjernkontrollen.no/", "https://www.fjernkontrollen.no/uploaded/files/2020_06_01_klimaregnskap_for_fjernvarme_2020.pdf"},
			Description: "Beregnet CO2 faktor for fjernvarme i Norge i 2022.",
		}

		fmt.Println(internal.PrettyPrint(districtHeating));
	},
}
