package set

func Intersect[K comparable](set1, set2 map[K]struct{}) map[K]struct{} {
	intersection := make(map[K]struct{})
	for k := range set1 {
		if _, ok := set2[k]; ok {
			intersection[k] = struct{}{}
		}
	}
	return intersection
}

func Difference[K comparable](set1, set2 map[K]struct{}) map[K]struct{} {
	diff := make(map[K]struct{})
	for k := range set1 {
		if _, ok := set2[k]; !ok {
			diff[k] = struct{}{}
		}
	}
	return diff
}

func Union[K comparable](set1, set2 map[K]struct{}) map[K]struct{} {
	union := make(map[K]struct{})
	for k := range set1 {
		union[k] = struct{}{}
	}
	for k := range set2 {
		union[k] = struct{}{}
	}
	return union
}

func SymmetricDifference[K comparable](set1, set2 map[K]struct{}) map[K]struct{} {
	diff1 := Difference(set1, set2)
	diff2 := Difference(set2, set1)
	return Union(diff1, diff2)
}

func IsSubset[K comparable](set1, set2 map[K]struct{}) bool {
	for k := range set1 {
		if _, ok := set2[k]; !ok {
			return false
		}
	}
	return true
}

func FromSlice[K comparable](slice []K) map[K]struct{} {
	set := make(map[K]struct{})
	for _, v := range slice {
		set[v] = struct{}{}
	}
	return set
}

func ToSlice[K comparable](set map[K]struct{}) []K {
	slice := make([]K, 0, len(set))
	for k := range set {
		slice = append(slice, k)
	}
	return slice
}
