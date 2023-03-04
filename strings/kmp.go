package strings

/* LPS - Longest Prefix Suffix
This denotes the longest proper(excluding complete string itself) prefix
which is also a suffix, for substring ending at index i.
This will help us avoid resetting the pointer to beginning in case of a mismatch,
and utilize information about the pattern to identify partial matches.
*/
func ComputeLPS(pat string) []int {
	lps := make([]int, len(pat))
	// lps[0] will always be 0 as there are no proper prefixes == suffix for a single character
	i, prevLPS := 1, 0

	for i < len(pat) {
		if pat[i] == pat[prevLPS] {
			// characters match, so lps value for this index has to be incremented
			lps[i] = prevLPS + 1
			prevLPS++
			i++
		} else if prevLPS == 0 {
			// We'll reach here after repeated mismatches. No partial characters can be matched,
			// so we set lps to 0 and proceed to the next index
			lps[i] = 0
			i++
		} else {
			// In case of mismatch, backtrack to the previous lps value to check
			//for potential partial match
			prevLPS = lps[prevLPS-1]
		}
	}
	return lps
}

// KMP algorithm for finding first instance of a substring in another string
func KMP(s, pat string, lps []int) int {
	i, j := 0, 0

	for i < len(s) {
		if s[i] == pat[j] {
			// characters match, continue with current substr
			i++
			j++
		} else if j == 0 {
			// we've backtracked as much as possible in the pattern,
			// no possibility of partial match in current substring.
			// Move the window further and start matching with pattern from beginning.
			i++
		} else {
			// mismatch found, check for previous lps value
			j = lps[j-1]
		}
		if j == len(pat) {
			// Complete substring match found
			return i - len(pat)
		}
	}
	return -1
}
