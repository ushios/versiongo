package versiongo

import (
	"errors"
	"fmt"
)

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

	return LessThan, err
}
