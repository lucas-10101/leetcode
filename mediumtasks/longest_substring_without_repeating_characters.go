package mediumtasks

func lengthOfLongestSubstring(s string) int {
	max := 0
	for i := 0; i < len(s); i++ {

		j := i + 1
		chars := map[byte]*int{s[i]: nil}

		for ; j < len(s); j++ {
			_, present := chars[s[j]]

			if present {
				if max < len(s[i:j]) {
					max = len(s[i:j])
				}
				break
			} else {
				chars[s[j]] = nil
			}
		}

		if max < len(chars) {
			max = len(chars)
		}
	}

	if max == 0 {
		return len(s)
	}

	return max
}
