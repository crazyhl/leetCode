package main

import (
	"bytes"
	"fmt"
)

/*
344. 反转字符串
请编写一个函数，其功能是将输入的字符串反转过来。


思路： 把字符串拆分反向遍历，然后用 buffer 拼起来就 ok 了
*/
func main() {
	reverseString := reverseString("hello")
	fmt.Println("reverseString:", reverseString)
}

func reverseString(s string) string {
	b := bytes.Buffer{}
	for i := len(s); i > 0; i-- {
		fmt.Println(s[i-1])
		b.WriteString(string(s[i-1]))
	}

	return b.String()
}

