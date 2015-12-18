package versiongo

import (
	"errors"
	"strconv"
)

// Split version string to int slice.
func Split(version string) ([]int, error) {
	var segments []int
	var tempSegment int
	var err error
	var added bool

	segments = []int{}
	tempSegment = 0
	added = false

	for _, r := range version {
		char := string(r)
		if num, err := checkChar(char); err == nil {
			tempSegment = tempSegment*10 + num
			added = true
		} else {
			if added {
				segments = append(segments, tempSegment)
				tempSegment = 0
				added = false
			}
		}
	}

	if added {
		segments = append(segments, tempSegment)
	}

	if len(segments) < 1 {
		return segments, errors.New("Version number not found.")
	}

	return segments, err
}

// checkChar
func checkChar(char string) (num int, err error) {
	if num, err = strconv.Atoi(char); err == nil {
		return num, err
	}

	return 0, err
}
