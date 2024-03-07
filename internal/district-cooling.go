package internal

/*
	Regnes ut basert på nasjonal fordeling oppgitt på https://www.fjernkontrollen.no/ for 2022,
	og med utslippsfaktorer fra https://www.fjernkontrollen.no/uploaded/files/2020_06_01_klimaregnskap_for_fjernvarme_2020.pdf
	og gjennomsnittsbergeningen for utslipp fra fjernvarmeinfrastruktur er hentet fra
	https://www.fjernkontrollen.no/uploaded/files/or.13.21_district_heating_infrastructure.pdf
*/
func districtCoolingNorway2022() EmissionFactor {
	/*
		Regnes ut basert på nasjonal fordeling oppgitt på https://www.fjernkontrollen.no/ for 2022,
		og med utslippsfaktorer fra https://www.fjernkontrollen.no/uploaded/files/2020_06_01_klimaregnskap_for_fjernvarme_2020.pdf
		og gjennomsnittsbergeningen for utslipp fra fjernvarmeinfrastruktur er hentet fra
		https://www.fjernkontrollen.no/uploaded/files/or.13.21_district_heating_infrastructure.pdf
	*/
	heatDrivenCoolingFactor := 0.0
	freeCoolingFactor := float64(ElectricityMixNorway2022.Factor) * 0.063 // 6.3% er hjelpeelektrisitet 
	heatPumpFactor := float64(ElectricityMixNorway2022.Factor) * 0.25 // 25% er bruk av elektrisitet
	compressorCoolingFactor := float64(ElectricityMixNorway2022.Factor) * 0.25 // 25% er bruk av elektrisitet
	
	factor := 
		0.408 * float64(ElectricityMixNorway2022.Factor) +  // vet ikke hvilken kilde, så bruker elektritetsmiksen
		0.087 * freeCoolingFactor +
		0.022 * heatDrivenCoolingFactor + 
		0.44 * heatPumpFactor+
		0.043 * compressorCoolingFactor


	return EmissionFactor{
		Type: DistrictCooling,
		Year: 2022,
		Location: NORWAY,
		Factor: factor,
		Unit: "gram CO2e/kWh",
		Sources:
			[]string{
				"https://www.fjernkontrollen.no/", 
				"https://www.fjernkontrollen.no/uploaded/files/2020_06_01_klimaregnskap_for_fjernvarme_2020.pdf",
				"https://www.fjernkontrollen.no/uploaded/files/or.13.21_district_heating_infrastructure.pdf",
			},
		Name: "Fjernkjøling i Norge 2022",
		Description: "Beregnet CO2 faktor for fjernkjøling i Norge i 2022.",
	}
}

var DistrictCoolingNorway2022 = districtCoolingNorway2022();
