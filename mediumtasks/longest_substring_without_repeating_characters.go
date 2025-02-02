package mediumtasks

func lengthOfLongestSubstring(s string) int {

	var max, leftWindow int
	lastPositions := map[byte]int{}

	// " "
	for i := 0; i < len(s); i++ {

		lastPos, found := lastPositions[s[i]]

		if found {
			if max < (i - leftWindow) {
				max = i - leftWindow
			}

			if lastPos >= leftWindow {
				leftWindow = lastPos + 1
			}
		}

		lastPositions[s[i]] = i
	}

	if max < (len(s) - leftWindow) {
		max = len(s) - leftWindow
	}

	return max
}
