package versiongo

// Result have Compare's result.
type Result int

const (
	// UnKnown is almost error.
	UnKnown Result = iota
	// Equals is compared same version.
	Equals
	// LessThan is past less than next.
	LessThan
	// GreaterThan is past greater than next.
	GreaterThan
)
