package bytep

// Equal 比较.
// Deprecated: 使用 bytes.Equal.
func Equal(bytesA, bytesB []byte) bool {
	var (
		lengthA = len(bytesA)
		lengthB = len(bytesB)
	)

	switch {
	case lengthA != lengthB:
		return false
	case lengthA == 0:
		return true
	default:
		for idx, item := range bytesA {
			if bytesB[idx] != item {
				return false
			}
		}

		return true
	}
}
