package leetcode

import "fmt"

func isRectangleCover(rectangles [][]int) bool {
	minx, miny := rectangles[0][0], rectangles[0][1]
	maxx, maxy := rectangles[0][2], rectangles[0][3]
	allarea := 0
	set := make(map[int]struct{}, 8)
	for i := range rectangles {
		x1, y1, x2, y2 := rectangles[i][0], rectangles[i][1], rectangles[i][2], rectangles[i][3]
		if x1 < minx {
			minx = x1
		}
		if y1 < miny {
			miny = y1
		}
		if x2 > maxx {
			maxx = x2
		}
		if y2 > maxy {
			maxy = y2
		}
		allarea += (x2 - x1) * (y2 - y1)
		if _, ok := set[x1<<16+y1]; ok {
			delete(set, x1<<16+y1)
		} else {
			set[x1<<16+y1] = struct{}{}
		}
		if _, ok := set[x1<<16+y2]; ok {
			delete(set, x1<<16+y2)
		} else {
			set[x1<<16+y2] = struct{}{}
		}
		if _, ok := set[x2<<16+y1]; ok {
			delete(set, x2<<16+y1)
		} else {
			set[x2<<16+y1] = struct{}{}
		}
		if _, ok := set[x2<<16+y2]; ok {
			delete(set, x2<<16+y2)
		} else {
			set[x2<<16+y2] = struct{}{}
		}
	}
	_, ok1 := set[minx<<16+miny]
	_, ok2 := set[minx<<16+maxy]
	_, ok3 := set[maxx<<16+miny]
	_, ok4 := set[maxx<<16+maxy]
	if !ok1 || !ok2 || !ok3 || !ok4 || len(set) != 4 {
		return false
	}
	return allarea == (maxy-miny)*(maxx-minx)
}

func main() {
	fmt.Println(isRectangleCover([][]int{
		{1, 1, 3, 3},
		{3, 1, 4, 2},
		{3, 2, 4, 4},
		{1, 3, 2, 4},
		{2, 3, 3, 4}}))
	fmt.Println(isRectangleCover([][]int{
		{1, 1, 3, 3},
		{3, 1, 4, 2},
		{1, 3, 2, 4},
		{3, 2, 4, 4}}))
	fmt.Println(isRectangleCover([][]int{
		{1, 1, 3, 3},
		{3, 1, 4, 2},
		{1, 3, 2, 4},
		{2, 2, 4, 4}}))
	fmt.Println(isRectangleCover([][]int{{0, 0, 1, 1}, {0, 1, 3, 2}, {1, 0, 2, 2}}))
}
