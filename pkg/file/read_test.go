package file

import "testing"

// Tests the reading and newline stripping of a file with a custom path.
func TestReadFile(t *testing.T) {
	r, err := ReadFile("infiles/1.txt")
	e := `1
2
3
4
5`

	if err != nil {
		t.Errorf("ReadFile() failed: %v", err)
		return
	}

	if r != e {
		t.Errorf("ReadFile() = %v, expected %v", r, e)
	}
}

// Tests the reading and newline stripping of a file with a custom path from
// the root of the module.
func TestReadFileFromRoot(t *testing.T) {
	r, err := readFromRoot("infiles/3005.txt")
	e := `1
2
3
4
5`

	if err != nil {
		t.Errorf("readFileFromRoot() failed: %v", err)
		return
	}

	if r != e {
		t.Errorf("readFileFromRoot() = %v, expected %v", r, e)
	}
}

// Tests the abstraction of reading an input file for a specific day.
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

// Tests the abstraction of reading an input file with a variation for a
// specific day.
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

// Tests the abstraction of reading an input file for a specific day relative
// to the package directory.
func TestReadInfileFromPackage(t *testing.T) {
	r, err := ReadInfile(1, "", "true")
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

// Tests the abstraction of reading an input file with a variation for a
// specific day relative to the package directory.
func TestReadInfileWithVariationFromPackage(t *testing.T) {
	r, err := ReadInfile(1, "1", "true")
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

// Tests the abstraction of reading a test input file for a specific day.
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

// Tests the abstraction of reading a test input file for a specific day.
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

// Tests the abstraction of reading a test input file for a specific day
// relative to the package directory.
func TestReadTestFileFromPackage(t *testing.T) {
	r, err := ReadTestFile(1, "", "true")
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

// Tests the abstraction of reading a test input file with a variation for a
// specific day relative to the package directory.
func TestReadTestFileWithVariationFromPackage(t *testing.T) {
	r, err := ReadTestFile(1, "1", "true")
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
