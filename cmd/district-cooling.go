package cmd

import (
	"errors"
	"fmt"
	"math"

	"github.com/graabeinstudio/emission-factors/internal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(districtCoolingCommand)
}

var districtCoolingCommand = &cobra.Command{
	Use:   "district-cooling",
	Short: "Run script for generating the district cooling emission factor",
	Long: `Run script for generating the district cooling emission factor.
			Currently only supports the district cooling for Norway in 2022.`,
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
		heatDrivenCoolingFactor := 0.0
		freeCoolingFactor := float64(internal.ElectricityMixNorway2022.Factor) * 0.063 // 6.3% er hjelpeelektrisitet 
		heatPumpFactor := float64(internal.ElectricityMixNorway2022.Factor) * 0.25 // 25% er bruk av elektrisitet
		compressorCoolingFactor := float64(internal.ElectricityMixNorway2022.Factor) * 0.25 // 25% er bruk av elektrisitet
		 
		factor := 
			0.408 * float64(internal.ElectricityMixNorway2022.Factor) +  // vet ikke hvilken kilde, så bruker elektritetsmiksen
			0.087 * freeCoolingFactor +
			0.022 * heatDrivenCoolingFactor + 
			0.44 * heatPumpFactor+
			0.043 * compressorCoolingFactor


		districtCooling := internal.EmissionFactor{
			Type: internal.DistrictCooling,
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
			Description: "Beregnet CO2 faktor for fjernkjøling i Norge i 2022.",
		}

		fmt.Println(internal.PrettyPrint(districtCooling));
	},
}
