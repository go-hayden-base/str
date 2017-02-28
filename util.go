package str

// Index return index of s in slice src
func Index(src []string, s string) int {
	for idx, itme := range src {
		if itme == s {
			return idx
		}
	}
	return -1
}

// Contains return slice src contains s or not
func Contains(src []string, s string) bool {
	return Index(src, s) > -1
}

// Distinct remove repetition
func Distinct(src []string) []string {
	l := len(src)
	if l == 0 {
		return nil
	}
	m := make(map[string]bool)
	for _, item := range src {
		m[item] = true
	}

	l = len(m)
	res := make([]string, 0, l)
	for key, _ := range m {
		res = append(res, key)
	}
	return res
}
