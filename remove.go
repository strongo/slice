package slice

func RemoveStringInPlace(v string, slice []string) []string {
	shift := 0
	for i, s := range slice {
		if s == v {
			shift++
			continue
		}
		slice[i-shift] = s
	}
	return slice[:len(slice)-shift]
}
