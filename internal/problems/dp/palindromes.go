package dp

/*
Longest Palindromic Substring

Given a string s, return the longest palindromic substring in s.

Example:

Input = "cbbd"
Output = "bb"
*/

func LongestPalindromicSubstring(s string) string {

	// This is an O(n^2) solution

	ret := ""
	retLen := 0

	for i, c := range s {

		// Odd-length palindromes
		odd := string(c)

		// Start with characters on the left and right, and expand outwards as far as you can
		l, r := i-1, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			odd = s[l : r+1]
			l--
			r++
		}
		if len(odd) > retLen {
			retLen = len(odd)
			ret = odd
		}

		// Even-length palindromes
		if i+1 < len(s) {

			// Start with two characters and start expanding outwards
			l, r = i, i+1
			even := s[l : r+1]
			for l >= 0 && r < len(s) && s[l] == s[r] {
				even = s[l : r+1]
				if len(even) > retLen {
					retLen = len(even)
					ret = even
				}
				l--
				r++
			}
		}
	}
	return ret
}

func PalindromicSubstrings(s string) (palindromes []string) {
	for i := range s {

		// Odd-length palindromes
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			palindromes = append(palindromes, s[l:r+1])
			l--
			r++
		}

		// Even-length palindromes
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			palindromes = append(palindromes, s[l:r+1])
			l--
			r++
		}
	}
	return palindromes
}
