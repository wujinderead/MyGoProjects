package stdlib

import (
	"fmt"
	"sort"
	"testing"

	"golang.org/x/exp/constraints"
)

func Sort[V constraints.Ordered](arr []V) {
	for i := len(arr); i > 0; i-- {
		for j := 0; j < i-1; j++ {
			if arr[j+1] < arr[j] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}

func SortByLessFun[V any](arr []V, less func(a, b V) bool) {
	for i := len(arr); i > 0; i-- {
		for j := 0; j < i-1; j++ {
			if less(arr[j+1], arr[j]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func SortByKey[V any, K constraints.Ordered](arr []V, key func(a V) K) {
	for i := len(arr); i > 0; i-- {
		for j := 0; j < i-1; j++ {
			if key(arr[j+1]) < key(arr[j]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func TestGenericSort(t *testing.T) {
	t.Run("sort ordered", func(t *testing.T) {
		t.Run("[]int", func(t *testing.T) {
			arr := []int{4, 3, 1, 6, 9, 3, 6, 7}
			Sort[int](arr)
			fmt.Println(arr)
		})
		t.Run("[]string", func(t *testing.T) {
			arr := []string{"ba", "cca", "a", "bba", "ca", "acc"}
			Sort[string](arr)
			fmt.Println(arr)
		})
		t.Run("can't compile", func(t *testing.T) {
			type unodered struct {
				a int
				b string
			}
			arr := []unodered{{a: 1, b: "2"}, {a: 2, b: "a"}}
			// Sort[unodered](arr)   can't compile: unodered does not implement constraints.Ordered
			fmt.Println(arr)
		})
	})

	t.Run("sort by less", func(t *testing.T) {
		t.Run("[]int", func(t *testing.T) {
			arr := []int{4, 3, 1, 6, 9, 3, 6, 7}
			SortByLessFun[int](arr, func(a, b int) bool { return a < b })
			fmt.Println(arr)
		})
		t.Run("[]string", func(t *testing.T) {
			arr := []string{"ba", "cca", "a", "bba", "ca", "acc"}
			SortByLessFun[string](arr, func(a, b string) bool { return a < b })
			fmt.Println(arr)
		})
		t.Run("[]string len", func(t *testing.T) {
			arr := []string{"ba", "cca", "a", "bba", "ca", "adcc"}
			SortByLessFun[string](arr, func(a, b string) bool { return len(a) < len(b) })
			fmt.Println(arr)
		})
		t.Run("[]struct", func(t *testing.T) {
			type unodered struct {
				a int
				b string
			}
			arr := []unodered{{a: 2, b: "b"}, {a: 3, b: "a"}, {a: 1, b: "c"}}
			SortByLessFun[unodered](arr, func(a, b unodered) bool { return a.a < b.a })
			fmt.Println(arr)
		})
	})

	t.Run("sort by key", func(t *testing.T) {
		t.Run("[]int", func(t *testing.T) {
			arr := []int{4, 3, 1, 6, 9, 3, 6, 7}
			SortByKey[int](arr, func(a int) int { return a })
			fmt.Println(arr)
		})
		t.Run("[]int square", func(t *testing.T) {
			arr := []int{4, -3, 1, 6, -9, 3, -6, 7}
			SortByKey[int](arr, func(a int) int { return a * a })
			fmt.Println(arr)
		})
		t.Run("[]string len", func(t *testing.T) {
			arr := []string{"ba", "cca", "a", "bba", "ca", "adcc"}
			SortByKey[string](arr, func(a string) int { return len(a) })
			fmt.Println(arr)
		})
		t.Run("[]struct", func(t *testing.T) {
			type unodered struct {
				a int
				b string
			}
			arr := []unodered{{a: 1, b: "b"}, {a: 3, b: "a"}, {a: 2, b: "c"}}
			SortByKey[unodered](arr, func(a unodered) string { return a.b })
			fmt.Println(arr)
		})
	})
}
