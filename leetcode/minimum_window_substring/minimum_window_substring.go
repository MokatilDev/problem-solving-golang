package minimum_window_substring

import "math"

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 || len(t) > len(s) {
		return ""
	}

	count1 := make(map[byte]int)
	for i := range t {
		count1[t[i]]++
	}

	count2 := make(map[byte]int)

	have, need := 0, len(count1)
	resStart, resLen := -1, math.MaxInt32

	left := 0
	for right := 0; right < len(s); right++ {
		rightChar := s[right]

		if count1[rightChar] > 0 {
			count2[rightChar]++

			if count2[rightChar] == count1[rightChar] {
				have++
			}
		}

		for have == need {
			if (right - left + 1) < resLen {
				resLen = right - left + 1
				resStart = left
			}

			leftChar := s[left]

			if count1[leftChar] > 0 {
				count2[leftChar]--

				if count2[leftChar] < count1[leftChar] {
					have--
				}
			}

			left++
		}
	}

	if resLen == math.MaxInt32 {
		return ""
	}

	return s[resStart : resStart+resLen]
}
