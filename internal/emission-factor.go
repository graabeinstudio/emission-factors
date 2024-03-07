package internal

import "encoding/json"

type EmissionFactorType struct {
	Name         string `json:"name"`
	FriendlyName string `json:"friendlyName"`
	Description  string `json:"description"`
}

var (
	ElectricityMix  EmissionFactorType = EmissionFactorType{"ELECTRICITY_MIX", "Electricity mix", "Greenhouse gas emission per kwh consumed from the electricity grid in a specific region"}
	DistrictHeating EmissionFactorType = EmissionFactorType{"DISTRICT_HEATING", "District heating", "Greenhouse gas emission per kwh of district heating"}
	DistrictCooling EmissionFactorType = EmissionFactorType{"DISTRICT_COOLING", "District cooling", "Greenhouse gas emission per kwh of district cooling"}
)

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

type EmissionFactor struct {
	Type        EmissionFactorType `json:"type"`
	Location    Location           `json:"location"`
	Year        int                `json:"year"`
	Factor      int                `json:"factor"`
	Unit        string             `json:"unit"`
	Sources     []string           `json:"sources"`
	Description string             `json:"description"`
}

var EmissionFactorTypes = []EmissionFactorType{ElectricityMix, DistrictHeating, DistrictCooling}
var EmissionFactorsNorway2022 = []EmissionFactor{ElectricityMixNorway2022, DistrictHeatingNorway2022, DistrictCoolingNorway2022}
