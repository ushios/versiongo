package versiongo

import "testing"

func TestCompare(t *testing.T) {
	result, err := Compare("1.0.2", "1.1.0")

	if err != nil {
		t.Errorf("result got err (%s)", err)
	}

	if result != LessThan {
		t.Errorf("result want (%s)", result)
	}

	result, err = Compare("20.0.0", "v10.0.0")

	if err != nil {
		t.Errorf("result got err (%s)", err)
	}

	if result != GreaterThan {
		t.Errorf("result want (%s)", result)
	}
}
