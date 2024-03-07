package internal

var ElectricityMixNorway2022 = EmissionFactor{
	Type: ElectricityMix,
	Year: 2022,
	Location: NORWAY,
	Factor: 19,
	Unit: "gram CO2e/kWh",
	Sources:
		[]string{"https://www.nve.no/energi/energisystem/kraftproduksjon/hvor-kommer-stroemmen-fra/"},
	Name: "Elektrisitetsmiks for Norge i 2022",
	Description: "Beregnet CO2 faktor for strømforbruk i Norge i 2022.",
}