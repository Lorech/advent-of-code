package cslices

// Return the elements from `a` that are not in `b`.
func Difference[T comparable](a, b []T) []T {
	mb := make(map[T]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	out := make([]T, 0)
	for _, x := range a {
		if _, found := mb[x]; !found {
			out = append(out, x)
		}
	}
	return out
}

// Return the elements that are in both `a` and `b`.
func Intersection[T comparable](a, b []T) []T {
	mb := make(map[T]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	out := make([]T, 0)
	for _, x := range a {
		if _, ok := mb[x]; ok {
			out = append(out, x)
		}
	}
	return out
}

// Return the elements that are in either `a` or `b`.
func Union[T comparable](a, b []T) []T {
	seen := make(map[T]struct{}, len(a)+len(b))
	out := make([]T, 0)

	for _, x := range a {
		if _, ok := seen[x]; !ok {
			seen[x] = struct{}{}
			out = append(out, x)
		}
	}
	for _, x := range b {
		if _, ok := seen[x]; !ok {
			seen[x] = struct{}{}
			out = append(out, x)
		}
	}
	return out
}
