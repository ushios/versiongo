package versiongo

import "testing"

func TestCompare(t *testing.T) {
	past, _ := NewVersion("1.0.2", FuzzySplit)
	next, _ := NewVersion("1.1.0", FuzzySplit)
	result, err := Compare(past, next)

	if err != nil {
		t.Errorf("result got err (%s)", err)
	}

	if result != LessThan {
		t.Errorf("result want (%s)", result)
	}

	past, _ = NewVersion("20.0.0", FuzzySplit)
	next, _ = NewVersion("v10.0.0", FuzzySplit)
	result, err = Compare(past, next)

	if err != nil {
		t.Errorf("result got err (%s)", err)
	}

	if result != GreaterThan {
		t.Errorf("result want (%s)", result)
	}

	past, _ = NewVersion("10.0.0", FuzzySplit)
	next, _ = NewVersion("v10.0.0", FuzzySplit)
	result, err = Compare(past, next)

	if err != nil {
		t.Errorf("result got err (%s)", err)
	}

	if result != Equals {
		t.Errorf("result want (%s)", result)
	}

	past, _ = NewVersion("0.9", FuzzySplit)
	next, _ = NewVersion("1.0", FuzzySplit)
	result, err = Compare(past, next)

	if err != nil {
		t.Errorf("result got err (%s)", err)
	}

	if result != LessThan {
		t.Errorf("result want (%s)", result)
	}

	past, _ = NewVersion("1.19", FuzzySplit)
	next, _ = NewVersion("1.19", FuzzySplit)
	result, err = Compare(past, next)

	if err != nil {
		t.Errorf("result got err (%s)", err)
	}

	if result != LessThan {
		t.Errorf("result want (%s)", result)
	}
}

func TestCompareWithError(t *testing.T) {
	past, _ := NewVersion("1.0.0", FuzzySplit)
	next, _ := NewVersion("1.0", FuzzySplit)
	result, err := Compare(past, next)

	if err == nil {
		t.Errorf("result want error but nil")
	}

	if result != UnKnown {
		t.Errorf("result want (%s)", result)
	}

	past, _ = NewVersion("aaa", FuzzySplit)
	next, _ = NewVersion("1.0.0", FuzzySplit)
	result, err = Compare(past, next)

	if err == nil {
		t.Errorf("result want error but nil")
	}

	if result != UnKnown {
		t.Errorf("result want (%s)", result)
	}
}
