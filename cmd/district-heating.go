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

		/*
			Regnes ut basert på nasjonal fordeling oppgitt på https://www.fjernkontrollen.no/ for 2022,
			og med utslippsfaktorer fra https://www.fjernkontrollen.no/uploaded/files/2020_06_01_klimaregnskap_for_fjernvarme_2020.pdf
			og gjennomsnittsbergeningen for utslipp fra fjernvarmeinfrastruktur er hentet fra
			https://www.fjernkontrollen.no/uploaded/files/or.13.21_district_heating_infrastructure.pdf
		*/
		recycledHeatFactor := 0.0
		fossilOilFactor := 286.0
		fossilGasFactor := 242.0
		bioEnergyFactor := 11.4
		surroundingHeatFactor := float64(internal.ElectricityMixNorway2022.Factor) * 0.306 // 30.6% er bruk av elektrisitet
		factor := 
			0.479 * recycledHeatFactor + 
			0.012 * fossilOilFactor + 
			0.022 * fossilGasFactor + 
			0.072 * float64(internal.ElectricityMixNorway2022.Factor) +
			0.321 * bioEnergyFactor + // vet ikke hvilken kilde, så tar snittet av alle utslippsfaktorene (TODO: Bruk fordeling fra https://www.fjernkontrollen.no/)
			0.094 * surroundingHeatFactor


		districtHeating := internal.EmissionFactor{
			Type: internal.DistrictHeating,
			Year: year,
			Location: location,
			Factor: int(math.Ceil(factor)),
			Unit: "gram CO2e/kWh",
			Sources:
				[]string{
					"https://www.fjernkontrollen.no/", 
					"https://www.fjernkontrollen.no/uploaded/files/2020_06_01_klimaregnskap_for_fjernvarme_2020.pdf",
					"https://www.fjernkontrollen.no/uploaded/files/or.13.21_district_heating_infrastructure.pdf",
				},
			Description: "Beregnet CO2 faktor for fjernvarme i Norge i 2022.",
		}

		fmt.Println(internal.PrettyPrint(districtHeating));
	},
}
