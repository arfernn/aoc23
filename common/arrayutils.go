package common

func Intersect(stringsA []string, stringsB []string) []string {

	fakeSet := make(map[string]bool)

	for _, str := range stringsA {
		if str != "" {
			fakeSet[str] = false
		}
	}

	result := make([]string, 0)
	for _, str := range stringsB {
		_, ok := fakeSet[str]
		if ok {
			result = append(result, str)
		}
	}

	return result

}

func SafeSlice(arr []uint64, i, j int) []uint64 {
	if i > len(arr) {
		return nil // Return nil if i is out of bounds
	}

	if j > len(arr) || j < 0 {
		j = len(arr) // Adjust j to the length of the array if it's out of bounds
	}

	return arr[i:j]
}
