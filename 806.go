package main

import "fmt"

/*
806. 写字符串需要的行数
我们要把给定的字符串 S 从左到右写到每一行上，每一行有一个最大的宽度 100 个单位，如果我们在写每个字母的时候会使这行超过了 100 个单位，那么我们应该把这个字母写到下一行。我们给定了一个数组 widths ，这个数组 widths[0] 代表 'a' 需要的单位， widths[1] 代表 'b' 需要的单位，...， widths[25] 代表 'z' 需要的单位.

现在回答两个问题：至少包含一个字符的 S 需要多少行，和最后一行需要的单位？返回你的答案是长度为 2 的 int 类型的列表

Example :
Input:
widths = [10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10]
S = "abcdefghijklmnopqrstuvwxyz"
Output: [3, 60]
Explanation:
所有的字符拥有相同的占用单位10. 书写所有的 26 个字母，
我们需要两个整行和占用60个单位的一行
Example :
Input:
widths = [4,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10]
S = "bbbcccdddaaa"
Output: [2, 4]
Explanation:
除去字母'a'所有的字符都是相同的单位10，并且字符串 "bbbcccdddaa" 将会覆盖 9 * 10 + 2 * 4 = 98 个单位.
最后一个字母 'a' ，他讲会被写到第二行，因为第一行只剩下 2 个单位了。
所以，这个答案是 2 行，第二行有4个单位宽


注:

字符串 S 的长度在 [1, 1000] 的范围.
S 只包含小写字母.
widths 是长度为 26的数组.
widths[i] 值的范围在 [2, 10].

思路，依次遍历每一个字符，而且我们知道只包含小写字母，所以根据 ascii 码可以的到每个字母得索引，得到宽度后，
计算行宽，如果宽度大于100 就移动到下一行即可
*/
func main() {
	arr := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	numberOfLines := numberOfLines(arr, "abcdefghijklmnopqrstuvwxyz")
	fmt.Println("Output is", numberOfLines)
}

func numberOfLines(widths []int, S string) []int {
	var returnArr []int
	lines := 0
	lineLength := 0
	if len(S) > 0 {
		lines = lines + 1
	}
	for _, char := range S {
		if lineLength+widths[char-97] > 100 {
			lines = lines + 1
			lineLength = 0
		}
		lineLength = lineLength + widths[char-97]
	}

	returnArr = append(returnArr, lines)
	returnArr = append(returnArr, lineLength)
	return returnArr
}

