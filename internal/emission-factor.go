package internal

import "encoding/json"

type EmissionFactorType = string

const (
	ElectricityMix EmissionFactorType = "Electricity mix"
	DistrictHeating EmissionFactorType = "District heating"
	DistrictCooling EmissionFactorType = "District cooling"
)

func PrettyPrint(i EmissionFactor) string {
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
