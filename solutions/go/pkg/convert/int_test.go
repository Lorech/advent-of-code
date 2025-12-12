package convert

import (
	"slices"
	"testing"
)

// Tests converting an int 0 to a false boolean.
func TestItobFalse(t *testing.T) {
	if r := Itob(0); r != false {
		t.Errorf("Itob() = %v, expected %v", r, false)
	}
}

// Tests converting a positive int to a true boolean.
func TestItobPositive(t *testing.T) {
	if r := Itob(1); r != true {
		t.Errorf("Itob() = %v, expected %v", r, true)
	}
}

// Tests converting a negative int to a true boolean.
func TestItobNegative(t *testing.T) {
	if r := Itob(-1); r != true {
		t.Errorf("Itob() = %v, expected %v", r, true)
	}
}

func TestBinToIntIndex(t *testing.T) {
	tests := []struct {
		name   string
		num    int
		digits int
		want   []int
	}{
		{
			name:   "consecutive bits",
			num:    0b11111,
			digits: 5,
			want:   []int{0, 1, 2, 3, 4},
		},
		{
			name:   "separated bits",
			num:    0b10001,
			digits: 5,
			want:   []int{0, 4},
		},
		{
			name:   "leading zeroes consecutive bits",
			num:    0b00111,
			digits: 5,
			want:   []int{2, 3, 4},
		},
		{
			name:   "leading zeroes separated bits",
			num:    0b00101,
			digits: 5,
			want:   []int{2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinToIntIndex(tt.num, tt.digits)
			if !slices.Equal(got, tt.want) {
				t.Errorf("BinToIntIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
