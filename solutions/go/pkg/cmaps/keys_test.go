package cmaps

import (
	"slices"
	"testing"
)

// Tests the ability to abtain a slice of keys from a map.
func TestKeysSlice(t *testing.T) {
	v := map[string]int{
		"foo": 1,
		"bar": 2,
	}

	e := []string{"bar", "foo"}
	r := KeysSlice(v)
	slices.Sort(r) // Sort the slice to ensure deterministic order.

	if len(r) != len(e) || r[0] != e[0] || r[1] != e[1] {
		t.Errorf("KeysSlice() = %v, expected %v", r, e)
	}
}
