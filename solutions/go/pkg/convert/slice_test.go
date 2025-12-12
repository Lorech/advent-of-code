package convert

import "testing"

// Tests slice conversion to integer with default configuration.
func TestSliceToInt(t *testing.T) {
	s := []int{1, 2, 3}
	e := 123
	if r, err := Stoi(s); err != nil || r != e {
		t.Errorf("Stoi() = %v, expected %v", r, e)
	}
}

// Tests slice conversion to integer with a different base.
func TestSliceToIntDifferentBase(t *testing.T) {
	s := []int{1, 2, 3}
	e := 83
	if r, err := Stoi(s, 8); err != nil || r != e {
		t.Errorf("Stoi() = %v, expected %v", r, e)
	}
}

// Tests slice conversion to integer with leading zeroes.
func TestSliceToIntWithZeroes(t *testing.T) {
	s := []int{0, 0, 0, 1, 2, 3}
	e := 123
	if r, err := Stoi(s); err != nil || r != e {
		t.Errorf("Stoi() = %v, expected %v", r, e)
	}
}

// Tests slice conversion to integer with an invalid slice.
func TestInvalidSliceToInt(t *testing.T) {
	s := []int{1, 2, 0}
	e := -1
	if r, err := Stoi(s, 1); err == nil || r != e {
		t.Errorf("Stoi() = %v, expected %v", r, e)
	}
}

func TestIntToBinRepresentation(t *testing.T) {
	tests := []struct {
		name string
		num  []int
		want int
	}{
		{
			name: "consecutive bits",
			num:  []int{0, 1, 2, 3, 4},
			want: 0b11111,
		},
		{
			name: "separated bits",
			num:  []int{0, 4},
			want: 0b10001,
		},
		{
			name: "leading zeroes consecutive bits",
			num:  []int{2, 3, 4},
			want: 0b00111,
		},
		{
			name: "leading zeroes separated bits",
			num:  []int{2, 4},
			want: 0b00101,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntIndexToBinary(tt.num)
			if got != tt.want {
				t.Errorf("IntToBinRepresentation() = %05b, want %05b", got, tt.want)
			}
		})
	}
}
