package slices

func ContainString(arr []string, x string) bool {
	for _, n := range arr {
		if x == n {
			return true
		}
	}
	return false
}

func ContainInt(arr []int, x int) bool {
	for _, n := range arr {
		if x == n {
			return true
		}
	}
	return false
}


