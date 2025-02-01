package mediumtasks

func lengthOfLongestSubstring(s string) int {

	max := 0
	chars := map[byte]int{}
	for i := 0; i < len(s); {
		old, present := chars[s[i]]
		chars[s[i]] = i

		if max < len(chars) {
			max = len(chars)
		}

		if present {
			clear(chars)
			i = old + 1
			continue
		}
		i++
	}

	return max
}
