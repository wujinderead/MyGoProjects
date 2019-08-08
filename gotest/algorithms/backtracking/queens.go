package backtracking

import "fmt"

type board struct {
	size  int
	piece [][2]int
	rows  []int
	adds  []int
	subs  []int
}

func newBoard(size int) *board {
	b := new(board)
	b.size = size
	b.piece = make([][2]int, size)
	b.adds = make([]int, 2*size-1)
	b.subs = make([]int, 2*size-1)
	b.rows = make([]int, size)
	return b
}

func (b *board) set(i, j int) {
	b.piece[j][0] = i
	b.piece[j][1] = j
	b.rows[i] = 1
	b.adds[i+j] = 1
	b.subs[i-j+b.size-1] = 1
}

func (b *board) unset(i, j int) {
	b.rows[i] = 0
	b.adds[i+j] = 0
	b.subs[i-j+b.size-1] = 0
}

func (b *board) canSet(i, j int) bool {
	if b.rows[i] == 0 && b.adds[i+j] == 0 && b.subs[i-j+b.size-1] == 0 {
		return true
	}
	return false
}

func (b *board) print() {
	fmt.Println(b.piece)
}

func (b *board) setCol(j int) {
	for i := 0; i < b.size; i++ {
		if b.canSet(i, j) {
			b.set(i, j)
			if j < b.size-1 {
				b.setCol(j + 1)
			} else {
				b.print()
			}
			b.unset(i, j)
		}
	}
}

func (b *board) setColTrack(j int) {
	for i := 0; i < b.size; i++ {
		if b.canSet(i, j) {
			for k := 0; k < j; k++ {
				fmt.Print("     ")
			}
			fmt.Printf("[%d,%d]\n", i, j)
			b.set(i, j)
			if j < b.size-1 {
				b.setColTrack(j + 1)
			}
			b.unset(i, j)
		}
	}
}

func queens(size int) {
	b := newBoard(size)
	b.setCol(0)
}

func queensNoStruct(size int) {
	piece := make([][2]int, size)
	adds := make([]int, 2*size-1)
	subs := make([]int, 2*size-1)
	rows := make([]int, size)
	setColNoStruct(size, 0, piece, adds, subs, rows)
}

func setColNoStruct(size, j int, piece [][2]int, adds, subs, rows []int) {
	for i := 0; i < size; i++ {
		if rows[i] == 0 && adds[i+j] == 0 && subs[i-j+size-1] == 0 {
			piece[j][0] = i
			piece[j][1] = j
			rows[i] = 1
			adds[i+j] = 1
			subs[i-j+size-1] = 1
			if j < size-1 {
				setColNoStruct(size, j+1, piece, adds, subs, rows)
			} else {
				fmt.Println(piece)
			}
			rows[i] = 0
			adds[i+j] = 0
			subs[i-j+size-1] = 0
		}
	}
}
