package main

import "fmt"

// https://leetcode.com/problems/number-of-music-playlists/

// Your music player contains N different songs and she wants to listen to L (not
// necessarily different) songs during your trip. You create a playlist so that: 
//   Every song is played at least once 
//   A song can only be played again only if K other songs have been played 
// Return the number of possible playlists. As the answer can be very large, 
// return it modulo 10^9 + 7. 
// Example 1: 
//   Input: N = 3, L = 3, K = 1
//   Output: 6
//   Explanation: There are 6 possible playlists. [1, 2, 3], [1, 3, 2], [2, 1, 3], [2, 3, 1], [3, 1, 2], [3, 2, 1].
// Example 2: 
//   Input: N = 2, L = 3, K = 0
//   Output: 6
//   Explanation: There are 6 possible playlists. [1, 1, 2], [1, 2, 1], [2, 1, 1], [2, 2, 1], [2, 1, 2], [1, 2, 2].
// Example 3: 
//   Input: N = 2, L = 3, K = 1
//   Output: 2
//   Explanation: There are 2 possible playlists. [1, 2, 1], [2, 1, 2]
// Note: 
//   0 <= K < N <= L <= 100 

// https://leetcode.com/problems/number-of-music-playlists/discuss/178415/C%2B%2BJavaPython-DP-Solution
// transition function: F(N,L,K) = F(N-1, L-1, K) * N + F(N, L-1, K) * (N - K)
// F(N-1, L-1, K): 
//   If only N - 1 in the L - 1 first songs.
//   We need to put the rest one at the end of music list.
//   Any song can be this last song, so there are N possible combinations.
// F(N, L-1, K):
//   If already N in the L - 1 first songs.
//   We can put any song at the end of music list,
//   but it should be different from K last song.
//   We have N - K choices.
func numMusicPlaylists(N int, L int, K int) int {
	mod := 1000000007
	old, new := make([]int, N+1), make([]int, N+1)   // old[i] means the number of playlists with i distinct songs  
	x := N
	for i := N-1; i>=N-K; i-- {
		x = (x*i)%mod
	}
	// when playlist length = K+1, the array is old[1]=old[2]...=old[K]=0, old[K+1]=N*(N-1)*...(N-K)
	old[K+1] = x

	// for playlist length from K+2 to L
	for leng:=K+2; leng<=L; leng++ {
		new[K+1] = old[K+1]
		for i:=K+2; i<=N; i++ {
			new[i] = old[i-1]*(N+1-i) + old[i]*(i-K)  // got this from pattern, but there is a transition function got from others
			new[i] = new[i]%mod
		}
		old, new = new, old
	}
	return old[N]
}

// pattern for N=4, L=6, K=1
// songs:  leng:  1  2   3   4   5   6          
//  1                0   0   0
//  2               12  12  12                   new[2]=old[2]
//  3                   24  72                   new[3]=old[2]*2 + old[3]*2
//  4                       24                   new[4]=old[3]*1 + old[4]*3

func main() {
	fmt.Println(numMusicPlaylists(3,3,1), 6)
	fmt.Println(numMusicPlaylists(2,3,0), 6)
	fmt.Println(numMusicPlaylists(2,3,1), 2)
	fmt.Println(numMusicPlaylists(4,6,0), 1560)
	fmt.Println(numMusicPlaylists(4,6,1), 600)
	fmt.Println(numMusicPlaylists(4,6,2), 168)
	fmt.Println(numMusicPlaylists(4,6,3), 24)
	fmt.Println(numMusicPlaylists(4,7,1), 2160)
	fmt.Println(numMusicPlaylists(4,7,2), 360)
	fmt.Println(numMusicPlaylists(50,70,30), 295368477)
	fmt.Println(numMusicPlaylists(1,4,0), 1)
	fmt.Println(numMusicPlaylists(2,2,1), 2)
	fmt.Println(numMusicPlaylists(2,2,0), 2)
}