package internal

import "encoding/json"

type EmissionFactorType = string

const (
	ElectricityMix EmissionFactorType = "ELECTRICITY_MIX"
	DistrictHeating EmissionFactorType = "DISTRICT_HEATING"
	DistrictCooling EmissionFactorType = "DISTRICT_COOLING"
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