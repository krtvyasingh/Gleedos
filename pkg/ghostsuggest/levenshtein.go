package ghostsuggest

func Levenshtein(a, b string) int {
	f := make([]int, len(b)+1)
	for j := range f {
		f[j] = j
	}
	for i := 1; i <= len(a); i++ {
		prev := i
		for j := 1; j <= len(b); j++ {
			cost := 0
			if a[i-1] != b[j-1] {
				cost = 1
			}
			cur := min(f[j]+1, min(prev+1, f[j-1]+cost))
			f[j-1] = prev
			prev = cur
		}
		f[len(b)] = prev
	}
	return f[len(b)]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
