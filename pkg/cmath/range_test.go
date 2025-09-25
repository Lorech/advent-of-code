package cmath_test

import (
	"lorech/advent-of-code/pkg/cmath"
	"testing"
)

func TestClosestInRange(t *testing.T) {
	tests := []struct {
		name                     string
		num                      int
		domainLower, domainUpper int
		targetLower, targetUpper int
		want                     int
	}{
		{
			name:        "inside range preserved",
			num:         5,
			domainLower: 1, domainUpper: 10,
			targetLower: 4, targetUpper: 7,
			want: 5,
		},
		{
			name:        "lower bound preserved",
			num:         4,
			domainLower: 1, domainUpper: 10,
			targetLower: 4, targetUpper: 7,
			want: 4,
		},
		{
			name:        "upper bound preserved",
			num:         7,
			domainLower: 1, domainUpper: 10,
			targetLower: 4, targetUpper: 7,
			want: 7,
		},
		{
			name:        "clamps onto lower bound",
			num:         3,
			domainLower: 1, domainUpper: 10,
			targetLower: 4, targetUpper: 7,
			want: 4,
		},
		{
			name:        "clamps onto upper bound",
			num:         8,
			domainLower: 1, domainUpper: 10,
			targetLower: 4, targetUpper: 7,
			want: 7,
		},
		{
			name:        "overflows when in range",
			num:         10,
			domainLower: 1, domainUpper: 10,
			targetLower: 4, targetUpper: 7,
			want: 7,
		},
		{
			name:        "underflows when in range",
			num:         1,
			domainLower: 1, domainUpper: 10,
			targetLower: 4, targetUpper: 7,
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cmath.ClosestInRange(tt.num,
				tt.domainLower, tt.domainUpper,
				tt.targetLower, tt.targetUpper)

			if got != tt.want {
				t.Errorf("ClosestInRange(%v, [%v..%v], [%v..%v]) = %v, want %v",
					tt.num, tt.domainLower, tt.domainUpper,
					tt.targetLower, tt.targetUpper, got, tt.want)
			}
		})
	}
}
