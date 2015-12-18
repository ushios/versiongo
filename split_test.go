package versiongo

import "testing"

func TestSplit(t *testing.T) {
	segment, err := Split("10.0.05")

	if err != nil {
		t.Error("Split got error.")
	}

	if segment[0] != 10 {
		t.Errorf("segment 0 want (%d)", segment[0])
	}

	if segment[1] != 0 {
		t.Errorf("segment 0 want (%d)", segment[1])
	}

	if segment[2] != 5 {
		t.Errorf("segment 0 want (%d)", segment[3])
	}
}

func TestSplitWithString(t *testing.T) {
	segment, err := Split("v12")

	if err != nil {
		t.Error("Split got error.")
	}

	if segment[0] != 12 {
		t.Errorf("segment 0 want (%d)", segment[0])
	}
}

func TestSplitWithDoubleString(t *testing.T) {
	segment, err := Split("vv15")

	if err != nil {
		t.Error("Split got error.")
	}

	if segment[0] != 15 {
		t.Errorf("segment 0 want (%d)", segment[0])
	}
}

func TestSplitWithAfterString(t *testing.T) {
	segment, err := Split("vv20version")

	if err != nil {
		t.Error("Split got error.")
	}

	if segment[0] != 20 {
		t.Errorf("segment 0 want (%d)", segment[0])
	}
}

func TestSplitWithZeroLast(t *testing.T) {
	segment, err := Split("1.1.0")

	if err != nil {
		t.Error("Split got error.")
	}

	if segment[2] != 0 {
		t.Errorf("segment 0 want (%d)", segment[0])
	}
}

func TestSplitFail(t *testing.T) {
	_, err := Split("version test")

	if err == nil {
		t.Errorf("Split result want error.")
	}
}
