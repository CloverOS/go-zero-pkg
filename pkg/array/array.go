package array

func IsContainString(item any, items []string) bool {
	for _, v := range items {
		if v == item {
			return true
		}
	}
	return false
}

func IsContainStringMapKey(target string, maps map[string]string) bool {
	for key := range maps {
		if target == key {
			return true
		}
	}
	return false
}

func StringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	//b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func StringSliceEqualWithoutSort(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	for _, v := range a {
		if !IsContainString(v, b) {
			return false
		}
	}
	return true
}
