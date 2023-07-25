package contains

func Contains(arr []string, item string) bool {
	for _, v := range arr {
		if item == v {
			return true
		}
	}
	return false
}
