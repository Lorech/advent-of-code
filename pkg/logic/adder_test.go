package logic

import "testing"

// Tests a half adder using 0 and 0.
func TestHalfAdderFalseFalse(t *testing.T) {
	esum, ecarry := false, false
	if rsum, rcarry := HalfAdder(false, false); rsum != esum || rcarry != ecarry {
		t.Errorf("HalfAdder() = %v, %v, expected %v, %v", rsum, rcarry, esum, ecarry)
	}
}

// Tests a half adder using 1 and 0.
func TestHalfAdderTrueFalse(t *testing.T) {
	esum, ecarry := true, false
	if rsum, rcarry := HalfAdder(true, false); rsum != esum || rcarry != ecarry {
		t.Errorf("HalfAdder() = %v, %v, expected %v, %v", rsum, rcarry, esum, ecarry)
	}
}

// Tests a half adder using 0 and 1.
func TestHalfAdderFalseTrue(t *testing.T) {
	esum, ecarry := true, false
	if rsum, rcarry := HalfAdder(false, true); rsum != esum || rcarry != ecarry {
		t.Errorf("HalfAdder() = %v, %v, expected %v, %v", rsum, rcarry, esum, ecarry)
	}
}

// Tests a half adder using 1 and 1.
func TestHalfAdderTrueTrue(t *testing.T) {
	esum, ecarry := false, true
	if rsum, rcarry := HalfAdder(true, true); rsum != esum || rcarry != ecarry {
		t.Errorf("HalfAdder() = %v, %v, expected %v, %v", rsum, rcarry, esum, ecarry)
	}
}

// Tests a full adder using 0 and 0 with no carry bit.
func TestFullAdderFalseFalseFalse(t *testing.T) {
	esum, ecarry := false, false
	if rsum, rcarry := FullAdder(false, false, false); rsum != esum || rcarry != ecarry {
		t.Errorf("FullAdder() = %v, %v, expected %v, %v", rsum, rcarry, esum, ecarry)
	}
}

// Tests a full adder using 0 and 0 with a carry bit.
func TestFullAdderFalseFalseTrue(t *testing.T) {
	esum, ecarry := true, false
	if rsum, rcarry := FullAdder(false, false, true); rsum != esum || rcarry != ecarry {
		t.Errorf("FullAdder() = %v, %v, expected %v, %v", rsum, rcarry, esum, ecarry)
	}
}

// Tests a full adder using 0 and 1 with no carry bit.
func TestFullAdderFalseTrueFalse(t *testing.T) {
	esum, ecarry := true, false
	if rsum, rcarry := FullAdder(false, true, false); rsum != esum || rcarry != ecarry {
		t.Errorf("FullAdder() = %v, %v, expected %v, %v", rsum, rcarry, esum, ecarry)
	}
}

// Tests a full adder using 0 and 1 with a carry bit.
func TestFullAdderFalseTrueTrue(t *testing.T) {
	esum, ecarry := false, true
	if rsum, rcarry := FullAdder(false, true, true); rsum != esum || rcarry != ecarry {
		t.Errorf("FullAdder() = %v, %v, expected %v, %v", rsum, rcarry, esum, ecarry)
	}
}

// Tests a full adder using 1 and 0 with no carry bit.
func TestFullAdderTrueFalseFalse(t *testing.T) {
	esum, ecarry := true, false
	if rsum, rcarry := FullAdder(true, false, false); rsum != esum || rcarry != ecarry {
		t.Errorf("FullAdder() = %v, %v, expected %v, %v", rsum, rcarry, esum, ecarry)
	}
}

// Tests a full adder using 1 and 0 with a carry bit.
func TestFullAdderTrueFalseTrue(t *testing.T) {
	esum, ecarry := false, true
	if rsum, rcarry := FullAdder(true, false, true); rsum != esum || rcarry != ecarry {
		t.Errorf("FullAdder() = %v, %v, expected %v, %v", rsum, rcarry, esum, ecarry)
	}
}

// Tests a full adder using 1 and 1 with no carry bit.
func TestFullAdderTrueTrueFalse(t *testing.T) {
	esum, ecarry := false, true
	if rsum, rcarry := FullAdder(true, true, false); rsum != esum || rcarry != ecarry {
		t.Errorf("FullAdder() = %v, %v, expected %v, %v", rsum, rcarry, esum, ecarry)
	}
}

// Tests a full adder using 1 and 1 with a carry bit.
func TestFullAdderTrueTrueTrue(t *testing.T) {
	esum, ecarry := true, true
	if rsum, rcarry := FullAdder(true, true, true); rsum != esum || rcarry != ecarry {
		t.Errorf("FullAdder() = %v, %v, expected %v, %v", rsum, rcarry, esum, ecarry)
	}
}
