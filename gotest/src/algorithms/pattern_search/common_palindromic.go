package pattern_search

// longest common substring
func lcsDp(a, b string) {

}

// longest palindromic substring
func lpsQuadratic(str string) (int, string) {
	maxlen := 0
	maxind := -1
	for i := 0; i < len(str); i++ {
		low, high := i, i
		for low-1 >= 0 && high+1 < len(str) && str[low-1] == str[high+1] {
			low--
			high++
		}
		if high-low+1 > maxlen {
			maxlen = high - low + 1
			maxind = low
		}
	}
	for i := 0; i < len(str)-1; i++ {
		l := 0
		for i-l >= 0 && i+1+l < len(str) && str[i-l] == str[i+1+l] {
			l++
		}
		if 2*l > maxlen {
			maxlen = 2 * l
			maxind = i - l + 1
		}
	}
	return maxind, str[maxind : maxind+maxlen]
}
