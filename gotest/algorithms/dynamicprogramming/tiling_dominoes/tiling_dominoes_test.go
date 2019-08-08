package tiling_dominoes

import "testing"

func TestBoard2xnDomino2x1(t *testing.T) {
	testcases := [][]int{{1, 1}, {2, 2}, {4, 5}, {5, 8}}
	for i := range testcases {
		ans := board2xnDomino2x1(testcases[i][0])
		if ans != testcases[i][1] {
			t.Errorf("index %d, expect %d, got %d", i, testcases[i][1], ans)
		}
	}
}

// 4 upper-right corner miss cases for n = 3, also 4 lower-right corner cases
// so there are 8 solutions for fill 3×3 board.
//    AA    AA    BC    BB
//    BBD   BCD   BCD   ACC
//    CCD   BCD   AAD   ADD
func TestBoard3xnDomino2x1(t *testing.T) {
	testcases := [][]int{{2, 3}, {3, 8}, {8, 153}, {12, 2131}}
	for i := range testcases {
		ans := board3xnDomino2x1(testcases[i][0])
		if ans != testcases[i][1] {
			t.Errorf("index %d, expect %d, got %d", i, testcases[i][1], ans)
		}
	}
}

func TestPaintFence(t *testing.T) {
	if paintFence(3, 2) != 6 {
		t.Fail()
	}
	if paintFence(5, 4) != 864 {
		t.Fail()
	}
}
