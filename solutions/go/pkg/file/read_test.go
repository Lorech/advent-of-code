package file

import "testing"

// Tests reading a puzzle input file.
func TestReadInfile(t *testing.T) {
	r, err := ReadInfile(3005)
	e := `1
2
3
4
5`

	if err != nil {
		t.Errorf("ReadInfile() failed: %v", err)
		return
	}

	if r != e {
		t.Errorf("ReadInfile() = %v, expected %v", r, e)
	}
}

// Tests reading a puzzle input variation file.
func TestReadInfileVariation(t *testing.T) {
	r, err := ReadInfile(3005, "var")
	e := `6
7
8
9
10`

	if err != nil {
		t.Errorf("ReadInfile() failed: %v", err)
		return
	}

	if r != e {
		t.Errorf("ReadInfile() = %v, expected %v", r, e)
	}
}

// Tests reading a puzzle test input file.
func TestReadTestFile(t *testing.T) {
	r, err := ReadTestFile(3005)
	e := `1
2
3
4
5`

	if err != nil {
		t.Errorf("ReadTestFile() failed: %v", err)
		return
	}

	if r != e {
		t.Errorf("ReadTestFile() = %v, expected %v", r, e)
	}
}

// Tests reading a puzzle test input variaton file.
func TestReadTestFileWithVariation(t *testing.T) {
	r, err := ReadTestFile(3005, "var")
	e := `6
7
8
9
10`

	if err != nil {
		t.Errorf("ReadTestFile() failed: %v", err)
		return
	}

	if r != e {
		t.Errorf("ReadTestFile() = %v, expected %v", r, e)
	}
}
