package main

import "fmt"

// https://leetcode.com/problems/product-of-the-last-k-numbers/

// Implement the class ProductOfNumbers that supports two methods:
// 1. add(int num)
//   Adds the number num to the back of the current list of numbers.
// 2. getProduct(int k)
//   Returns the product of the last k numbers in the current list.
// You can assume that always the current list has at least k numbers.
// At any time, the product of any contiguous sequence of numbers will fit into
// a single 32-bit integer without overflowing.
// Example:
//   Input: ["ProductOfNumbers","add","add","add","add","add","getProduct","getProduct","getProduct",
//          "add","getProduct"]
//          [[],[3],[0],[2],[5],[4],[2],[3],[4],[8],[2]]
//   Output:
//     [null,null,null,null,null,null,20,40,0,null,32]
//   Explanation:
//     ProductOfNumbers productOfNumbers = new ProductOfNumbers();
//     productOfNumbers.add(3);        // [3]
//     productOfNumbers.add(0);        // [3,0]
//     productOfNumbers.add(2);        // [3,0,2]
//     productOfNumbers.add(5);        // [3,0,2,5]
//     productOfNumbers.add(4);        // [3,0,2,5,4]
//     productOfNumbers.getProduct(2); // return 20. The product of the last 2 numbers is 5 * 4 = 20
//     productOfNumbers.getProduct(3); // return 40. The product of the last 3 numbers is 2 * 5 * 4 = 40
//     productOfNumbers.getProduct(4); // return 0. The product of the last 4 numbers is 0 * 2 * 5 * 4 = 0
//     productOfNumbers.add(8);        // [3,0,2,5,4,8]
//     productOfNumbers.getProduct(2); // return 32. The product of the last 2 numbers is 4 * 8 = 32
// Constraints:
//    There will be at most 40000 operations considering both add and getProduct.
//    0 <= num <= 100
//    1 <= k <= 40000

type ProductOfNumbers struct {
	buf  []int
	zind int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		buf:  make([]int, 0),
		zind: 0,
	}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.zind = 0
		this.buf = this.buf[:0]
		return
	}
	if len(this.buf) == 0 {
		this.buf = append(this.buf, num)
	} else {
		num = num * this.buf[len(this.buf)-1]
		this.buf = append(this.buf, num)
	}
	this.zind++
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if k > this.zind {
		return 0
	}
	if k == this.zind {
		return this.buf[len(this.buf)-1]
	}
	return this.buf[len(this.buf)-1] / this.buf[len(this.buf)-k-1]
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	productOfNumbers := Constructor()
	productOfNumbers.Add(3)                     // [3]
	fmt.Println(productOfNumbers.GetProduct(1)) // return 3
	productOfNumbers.Add(0)                     // [3,0]
	fmt.Println(productOfNumbers.GetProduct(1)) // return 0
	fmt.Println(productOfNumbers.GetProduct(2)) // return 0
	productOfNumbers.Add(2)                     // [3,0,2]
	productOfNumbers.Add(5)                     // [3,0,2,5]
	productOfNumbers.Add(4)                     // [3,0,2,5,4]
	fmt.Println(productOfNumbers.GetProduct(2)) // return 20. The product of the last 2 numbers is 5 * 4 = 20
	fmt.Println(productOfNumbers.GetProduct(3)) // return 40. The product of the last 3 numbers is 2 * 5 * 4 = 40
	fmt.Println(productOfNumbers.GetProduct(4)) // return 0. The product of the last 4 numbers is 0 * 2 * 5 * 4 = 0
	productOfNumbers.Add(8)                     // [3,0,2,5,4,8]
	fmt.Println(productOfNumbers.GetProduct(1)) // return 8.
	fmt.Println(productOfNumbers.GetProduct(2)) // return 32.
	fmt.Println(productOfNumbers.GetProduct(3)) // return 160.
	fmt.Println(productOfNumbers.GetProduct(4)) // return 320.
	fmt.Println(productOfNumbers.GetProduct(5)) // return 0.
	fmt.Println(productOfNumbers.GetProduct(6)) // return 0.
}
