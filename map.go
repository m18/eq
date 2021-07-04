package eq

// StringMaps checks if two maps contain the same data.
func StringMaps(x, y map[string]string) bool {
	if len(x) != len(y) {
		return false
	}
	for k, v := range x {
		if y[k] != v {
			return false
		}
	}
	return true
}

// StringToSimpleTypeMaps checks if two maps contain the same data.
func StringToSimpleTypeMaps(x, y map[string]interface{}) bool {
	if len(x) != len(y) {
		return false
	}
	for k, v := range x {
		if y[k] != v {
			return false
		}
	}
	return true
}

// StringSets checks if two sets containf the same items.
func StringSets(x map[string]struct{}, y map[string]struct{}) bool {
	if len(x) != len(y) {
		return false
	}
	for k := range x {
		if _, ok := y[k]; !ok {
			return false
		}
	}
	return true
}
