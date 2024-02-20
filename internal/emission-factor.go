package internal

import "encoding/json"

type EmissionFactorType = string

const (
	ElectricityMix EmissionFactorType = "ElectricityMix"
)

func PrettyPrint(i EmissionFactor) string {
    s, _ := json.MarshalIndent(i, "", "\t")
    return string(s)
}

type EmissionFactor struct {
	Type        EmissionFactorType `json:"type"`
	Location    string             `json:"location"`
	Year        int                `json:"year"`
	Factor      int                `json:"factor"`
	Unit        string             `json:"unit"`
	Source      string             `json:"source"`
	Description string             `json:"description"`
}
