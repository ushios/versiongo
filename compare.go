package versiongo

import "fmt"

// Compare is compare version string.
func Compare(past Version, next Version) (Result, error) {
	pastSegments, err := past.IntSegments()

	if err != nil {
		return UnKnown, err
	}

	nextSegments, err := next.IntSegments()

	if err != nil {
		return UnKnown, err
	}

	numSegments := len(pastSegments)
	if numSegments != len(nextSegments) {
		return UnKnown, fmt.Errorf("Num of segments not matched between (%s) and (%s)", pastSegments, nextSegments)
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
