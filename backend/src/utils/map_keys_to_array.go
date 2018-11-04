package utils

func MapKeyToArray(target map[string]interface{}) []string {
	keys := make([]string, 0, len(target))
	for k := range target {
		keys = append(keys, k)
	}
	return keys
}
