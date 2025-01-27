package golang

func romanToInt(s string) int {

	const (
		M = 77 // 1000
		D = 68 // 500
		C = 67 // 100
		L = 76 // 50
		X = 88 // 10
		V = 86 // 5
		I = 73 // 1
	)

	sum, l := 0, len(s)
	for i := 0; i < l; i++ {
		switch s[i] {
		case M:
			sum += 1000
		case D:
			sum += 500
		case C:
			// in bound and not modifier, same for the others
			if i+1 < l && (s[i+1] == D || s[i+1] == M) {
				sum -= 100
			} else {
				sum += 100
			}
		case L:
			sum += 50
		case X:
			if i+1 < l && (s[i+1] == L || s[i+1] == C) {
				sum -= 10
			} else {
				sum += 10
			}
		case V:
			sum += 5
		case I:
			if i+1 < l && (s[i+1] == X || s[i+1] == V) {
				sum -= 1
			} else {
				sum += 1
			}
		default:
		}
	}

	return sum
}
