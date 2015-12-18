package versiongo

import (
	"errors"
	"fmt"
)

// Compare is compare version string.
func Compare(past string, next string) (Result, error) {
	pastSegments, err := Split(past)

	if err != nil {
		return UnKnown, err
	}

	nextSegments, err := Split(next)

	if err != nil {
		return UnKnown, err
	}

	numSegments := len(pastSegments)
	if numSegments != len(nextSegments) {
		message := fmt.Sprintf("Num of segments not matched between (%s) and (%s)", pastSegments, nextSegments)
		return UnKnown, errors.New(message)
	}

	for i := 0; i < numSegments; i++ {
		if pastSegments[i] == nextSegments[i] {
			continue
		}

		if pastSegments[i] > nextSegments[i] {
			return GreaterThan, err
		}

		if pastSegments[i] < nextSegments[i] {
			return LessThan, err
		}
	}

	return Equals, err
}
