package internal

import "math"

func districtHeatingNorway2022() EmissionFactor {
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
	surroundingHeatFactor := float64(ElectricityMixNorway2022.Factor) * 0.306 // 30.6% er bruk av elektrisitet
	factor := 
		0.479 * recycledHeatFactor + 
		0.012 * fossilOilFactor + 
		0.022 * fossilGasFactor + 
		0.072 * float64(ElectricityMixNorway2022.Factor) +
		0.321 * bioEnergyFactor + // vet ikke hvilken kilde, så tar snittet av alle utslippsfaktorene (TODO: Bruk fordeling fra https://www.fjernkontrollen.no/)
		0.094 * surroundingHeatFactor


	return EmissionFactor{
		Type: DistrictHeating,
		Year: 2022,
		Location: NORWAY,
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
}

var DistrictHeatingNorway2022 = districtHeatingNorway2022();