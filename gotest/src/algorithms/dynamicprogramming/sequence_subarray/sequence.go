package longest

func longestIncreasingSubsequence(seq []int) int {
	// for seq[i], check if seq[i] (0≤k<i) > seq[k], if seq[i] > seq[k], seq[i] = max(seq[i], seq[k]+1)
	if len(seq) < 0 {
		return 0
	}
	maxlis := 1
	lis := make([]int, len(seq))
	for i := 0; i < len(seq); i++ {
		lis[i] = 1
		for j := 0; j < i; j++ {
			if seq[j] < seq[i] && lis[i] < lis[j]+1 {
				lis[i] = lis[j] + 1
			}
			if maxlis < lis[i] {
				maxlis = lis[i]
			}
		}
	}
	return maxlis
}
