ipackage main

import (
	"fmt"
)

/*
13. 罗马数字转整数

给定一个罗马数字，将其转换成整数。

返回的结果要求在 1 到 3999 的范围内。


思路： 具体解法请参考 https://baike.baidu.com/item/%E7%BD%97%E9%A9%AC%E6%95%B0%E5%AD%97/772296?fr=aladdin
*/
func main() {
	romanToInt := romanToInt("MCMXCVI")
	fmt.Println("romanToInt:", romanToInt)
}

func romanToInt(s string) int {
	intNum := 0
	var lastChar byte
	for index, char := range s {
		tempNum := 0
		if index > 0 {
			lastChar = s[index-1]
		}
		switch char {
		case 'I':
			tempNum = 1
		case 'V':
			tempNum = 5
		case 'X':
			tempNum = 10
		case 'L':
			tempNum = 50
		case 'C':
			tempNum = 100
		case 'D':
			tempNum = 500
		case 'M':
			tempNum = 1000
		}

		tempLastNum := 0
		switch lastChar {
		case 'I':
			tempLastNum = 1
		case 'X':
			tempLastNum = 10
		case 'C':
			tempLastNum = 100
		}
		if (lastChar == 'I' || lastChar == 'X' || lastChar == 'C') && tempLastNum > 0 && tempLastNum < tempNum {
			intNum = tempNum + intNum - tempLastNum*2
		} else {
			intNum = tempNum + intNum
		}
	}

	return intNum
}

