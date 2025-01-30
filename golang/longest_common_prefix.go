package golang

func longestCommonPrefix(strs []string) string {

	if len(strs) == 1 {
		return strs[0]
	}

	commonPrefix := strs[0]
	for i := 0; i < len(strs); i++ {

		newCommonPrefix := ""
		for j := 0; j < len(commonPrefix) && j < len(strs[i]) && commonPrefix[j] == strs[i][j]; j++ {
			newCommonPrefix += string(commonPrefix[j])
		}

		commonPrefix = newCommonPrefix
	}

	return commonPrefix
}
