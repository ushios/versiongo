package versiongo

import "testing"

func TestSplit(t *testing.T) {
	segment, err := FuzzySplit("10.0.05", []string{})

	if err != nil {
		t.Error("FuzzySplit got error.")
	}

	if segment[0] != "10" {
		t.Errorf("segment 0 want (%s)", segment[0])
	}

	if segment[1] != "0" {
		t.Errorf("segment 0 want (%s)", segment[1])
	}

	if segment[2] != "5" {
		t.Errorf("segment 0 want (%s)", segment[3])
	}
}

func TestFuzzySplitWithString(t *testing.T) {
	segment, err := FuzzySplit("v12", []string{})

	if err != nil {
		t.Error("FuzzySplit got error.")
	}

	if segment[0] != "12" {
		t.Errorf("segment 0 want (%s)", segment[0])
	}
}

func TestFuzzySplitWithDoubleString(t *testing.T) {
	segment, err := FuzzySplit("vv15", []string{})

	if err != nil {
		t.Error("FuzzySplit got error.")
	}

	if segment[0] != "15" {
		t.Errorf("segment 0 want (%s)", segment[0])
	}
}

func TestFuzzySplitWithAfterString(t *testing.T) {
	segment, err := FuzzySplit("vv20version", []string{})

	if err != nil {
		t.Error("FuzzySplit got error.")
	}

	if segment[0] != "20" {
		t.Errorf("segment 0 want (%s)", segment[0])
	}
}

func TestFuzzySplitWithZeroLast(t *testing.T) {
	segment, err := FuzzySplit("1.1.0", []string{})

	if err != nil {
		t.Error("FuzzySplit got error.")
	}

	if segment[2] != "0" {
		t.Errorf("segment 0 want (%s)", segment[0])
	}
}

func TestFuzzySplitFail(t *testing.T) {
	_, err := FuzzySplit("version test", []string{})

	if err == nil {
		t.Errorf("FuzzySplit result want error.")
	}
}
