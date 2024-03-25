package dp

/*
A message containing letters from A-Z can be encoded into numbers using the following mapping:
'A' -> "1"
'B' -> "2"
...
'Z' -> "26"
To decode an encoded message, all the digits must be grouped then mapped back into
letters using the reverse of the mapping above (there may be multiple ways).

For example, "11106" can be mapped into:
"AAJF" with the grouping (1 1 10 6)
"KJF" with the grouping (11 10 6)
Note that the grouping (1 11 06) is invalid because "06" cannot be mapped into 'F' since "6" is different from "06".
Given a string s containing only digits, return the number of ways to decode it.

Example 1:

Input: s = "12"
Output: 2
Explanation: "12" could be decoded as "AB" (1 2) or "L" (12).

Example 2:

Input: s = "226"
Output: 3
Explanation: "226" could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).

Example 3:

Input: s = "06"
Output: 0
Explanation: "06" cannot be mapped to "F" because of the leading zero ("6" is different from "06").

*/

var alphabet map[string]string = map[string]string{
	"1":  "A",
	"2":  "B",
	"3":  "C",
	"4":  "D",
	"5":  "E",
	"6":  "F",
	"7":  "G",
	"8":  "H",
	"9":  "I",
	"10": "J",
	"11": "K",
	"12": "L",
	"13": "M",
	"14": "N",
	"15": "O",
	"16": "P",
	"17": "Q",
	"18": "R",
	"19": "S",
	"20": "T",
	"21": "U",
	"22": "V",
	"23": "W",
	"24": "X",
	"25": "Y",
	"26": "Z",
}

func DecodeWaysRecursive(s string) int {
	// Base cases
	if len(s) == 0 {
		// Only 1 way to decode an empty string
		return 1
	}

	if string(s[0]) == "0" {
		// 0 starting a string is invalid, not a way to decode, so fail fast.
		return 0
	}

	if len(s) == 1 {
		// Only 1 way to decode a single char, and we already know it's not '0'.
		return 1
	}

	if len(s) == 2 {
		// 2 chars: could be a valid double-digit mapping, or not.

		if alphabet[s] != "" {
			// Double-digit mapping exists, that's 1 way, then add on an additional way
			// if the last character by itself exists.
			// Examples:
			//   18 -> both double-digit and single-digit works
			//   20 -> only double-digit works, because '0' by itself doesn't exist.
			return 1 + DecodeWaysRecursive(s[1:])
		} else {
			// Double-digit mapping doesn't exist. So try to treat the characters individually,
			// which boils down to seeing if the last character exists (is not 0).
			// Examples:
			//   35 -> both 3, 5 exist
			//   40 -> 4 exists, but '0' doesn't, so this whole thing cannot be decoded.
			return DecodeWaysRecursive(s[1:])
		}
	} else {
		// String is >2 chars in length, and doesn't start with '0', so we have to keep breaking it down.
		// Examples: 343

		if alphabet[s[:2]] != "" {
			// If next two chars are valid, try that decision tree and the rest of the string
			return DecodeWaysRecursive(s[2:]) + DecodeWaysRecursive(s[1:])
		}
		// Next two chars are not valid, so just try consuming one char
		return DecodeWaysRecursive(s[1:])
	}
}

func DecodeWaysMemoized(s string) int {

	if len(s) == 0 {
		return 1
	}

	if string(s[0]) == "0" {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	// Memo will store intermediate results
	memo := make([]int, len(s))

	// Figure out last two characters first
	double := s[len(s)-2:]
	single := s[len(s)-1:]

	if single != "0" {
		memo[len(memo)-1] = 1
		if alphabet[double] != "" {
			// This was something you could parse in two ways, i.e. "...23"
			memo[len(memo)-2] = 2
		} else {
			// This was something like "...35", or "...02"
			memo[len(memo)-2] = memo[len(memo)-1]
		}
	} else {
		// Last character is "0"
		if alphabet[double] != "" {
			// This is something valid, like "...20"
			memo[len(memo)-1] = 0
			memo[len(memo)-2] = 1
		} else {
			// This is something invalid, like "...70". No matter what way you parse this it will be incorrect.
			return 0
		}
	}

	// Start from the third character and work backwards
	for i := len(s) - 3; i >= 0; i-- {
		if s[i:i+1] == "0" {
			if s[i+1:i+2] == "0" {
				// We've found 2 "0" in a row, illegal
				return 0
			}
			// Otherwise just drag whatever we've found back 1 and move on to the next character
			memo[i] = memo[i+1]
		} else {
			double = s[i : i+2]
			if alphabet[double] != "" {
				if s[i+1:i+2] == "0" {
					// We've found a double, but no singles, only one way to evaluate
					memo[i] = memo[i+1]
					memo[i+1] = 0
				} else {
					// There's two ways to evaluate these characters, as double or as singles
					memo[i] = memo[i+1] + memo[i+2]
				}
			} else {
				if s[i+1:i+2] == "0" {
					// We've come across something like "70"
					return 0
				}
				memo[i] = memo[i+1]
			}
		}
	}

	return memo[0]
}
