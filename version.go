package versiongo

import "strconv"

// Version have version info.
type Version struct {
	Segments  []string
	Separator string
}

// NewVersion create version instance
func NewVersion(
	versionString string,
	spliter func(ver string, sep []string) ([]string, error),
) (Version, error) {
	version := Version{}
	segments, err := spliter(versionString, []string{""})

	if err != nil {
		return version, err
	}

	version.Segments = segments
	return version, err
}

// IntSegments return segments as int.
func (v Version) IntSegments() ([]int, error) {
	var err error
	segments := []int{}

	for _, segment := range v.Segments {
		val, err := strconv.Atoi(segment)
		if err != nil {
			return []int{}, err
		}
		segments = append(segments, val)
	}

	return segments, err
}
