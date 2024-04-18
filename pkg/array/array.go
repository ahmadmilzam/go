package array

// Contains is a function to check if a string exists in a slice of string
func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
