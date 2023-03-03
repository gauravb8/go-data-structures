package strings

// KMP algorithm for finding first instance of a substring in another string
func computeLPS(pat string) []int {
	lps := make([]int, len(pat))
	i, prevLPS := 1, 0

	for i < len(pat) {
		if pat[i] == pat[prevLPS] {
			lps[i] = prevLPS + 1
			prevLPS++
			i++
		} else if prevLPS == 0 {
			lps[i] = 0
			i++
		} else {
			prevLPS = lps[prevLPS-1]
		}
	}
	return lps
}

func KMP(s, pat string, lps []int) int {
	i, j := 0, 0

	for i < len(s) {
		if s[i] == pat[j] {
			i++
			j++
		} else if j == 0 {
			i++
		} else {
			j = lps[j-1]
		}
		if j == len(pat) {
			return i - len(pat)
		}
	}
	return -1
}
