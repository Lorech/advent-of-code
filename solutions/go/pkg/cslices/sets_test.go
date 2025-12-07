package cslices_test

import (
	"lorech/advent-of-code/pkg/cslices"
	"testing"
)

func equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestDifference(t *testing.T) {
	t.Run("strings", func(t *testing.T) {
		got := cslices.Difference([]string{"a", "b", "c"}, []string{"b"})
		want := []string{"a", "c"}
		if !equal(got, want) {
			t.Fatalf("Difference(strings) = %v; want %v", got, want)
		}
	})

	t.Run("ints", func(t *testing.T) {
		got := cslices.Difference([]int{1, 2, 3, 4}, []int{2, 4})
		want := []int{1, 3}
		if !equal(got, want) {
			t.Fatalf("Difference(ints) = %v; want %v", got, want)
		}
	})
}

func TestIntersection(t *testing.T) {
	t.Run("runes", func(t *testing.T) {
		got := cslices.Intersection([]rune{'a', 'b', 'c'}, []rune{'b', 'd'})
		want := []rune{'b'}
		if !equal(got, want) {
			t.Fatalf("Intersection(runes) = %v; want %v", got, want)
		}
	})

	t.Run("ints", func(t *testing.T) {
		got := cslices.Intersection([]int{1, 2, 3}, []int{3, 4, 1})
		want := []int{1, 3}
		if !equal(got, want) {
			t.Fatalf("Intersection(ints) = %v; want %v", got, want)
		}
	})
}

func TestUnion(t *testing.T) {
	t.Run("strings", func(t *testing.T) {
		got := cslices.Union([]string{"a", "b"}, []string{"b", "c"})
		want := []string{"a", "b", "c"}
		if !equal(got, want) {
			t.Fatalf("Union(strings) = %v; want %v", got, want)
		}
	})

	t.Run("ints", func(t *testing.T) {
		got := cslices.Union([]int{1, 2, 3}, []int{3, 4, 5})
		want := []int{1, 2, 3, 4, 5}
		if !equal(got, want) {
			t.Fatalf("Union(ints) = %v; want %v", got, want)
		}
	})
}
