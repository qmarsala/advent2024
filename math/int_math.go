package math

func GetDiff(a int64, b int64) int64 {
	if a < b {
		return b - a
	} else {
		return a - b
	}
}
