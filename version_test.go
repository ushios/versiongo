package versiongo

import "testing"

func TestNewVersion(t *testing.T) {
	version, err := NewVersion("1.0.5", FuzzySplit)
	if err != nil {
		t.Errorf("NewVersion go error (%s).", err)
	}

	segments := version.Segments

	if segments[0] != "1" {
		t.Errorf("segments 0 want (%s)", segments[0])
	}

	if segments[1] != "0" {
		t.Errorf("segments 0 want (%s)", segments[1])
	}

	if segments[2] != "5" {
		t.Errorf("segments 0 want (%s)", segments[3])
	}
}

func TestVersoinCompare(t *testing.T) {
	version, _ := NewVersion("10.0.0", FuzzySplit)

	compVer, _ := NewVersion("10.0.01", FuzzySplit)

	if res, _ := version.Compare(compVer); res != LessThan {
		t.Errorf("%s and %s result %s", version, compVer, res)
	}
}
