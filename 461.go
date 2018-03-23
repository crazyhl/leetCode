package main

import "fmt"

/*
461. Hamming Distance
The Hamming distance between two integers is the number of positions at which the corresponding bits are different.

Given two integers x and y, calculate the Hamming distance.

计算两个数字的汉明距离，汉明距离的说明可以参考 https://baike.baidu.com/item/%E6%B1%89%E6%98%8E%E8%B7%9D%E7%A6%BB/475174

说下我的方案，把两个数字进行异或运算，这样就把不同的数字取出来了，然后循环对比每一位 是要这个位是1 就把计数器加一
然后把数字右移一位，直到这个异或结果是0 的时候就说明没有1的位了，就可以退出循环了。计数器的结果就是汉明距离了
*/
func main() {
	hammingDistance := hammingDistance(1, 4)
	fmt.Println("hammingDistance is", hammingDistance)
}

func hammingDistance(x int, y int) int {
	result := x ^ y
	count := 0
	for {
		if (result & 1) == 1 {
			count++
		}
		result = result >> 1
		if result == 0 {
			break
		}
	}
	return count
}

