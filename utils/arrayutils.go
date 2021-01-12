package utils

func UniqueStrings(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
func ElementInStringArray(element string, array []string) bool {
	for _, entry := range array {
		if entry == element {
			return true
		}
	}
	return false
}
