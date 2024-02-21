package internal

import (
	"fmt"
)

type Location = string

const (
	NORWAY Location = "Norway"
)

func ToLocation(location string) (Location, error) {
	if (location == "norway") {
		return NORWAY, nil;
	} else {
		return "", fmt.Errorf("%s is not av valid location", location)
	}
}