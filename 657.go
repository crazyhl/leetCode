package main

import "fmt"

/*
657. Judge Route Circle
Initially, there is a Robot at position (0, 0). Given a sequence of its moves, judge if this robot makes a circle, which means it moves back to the original place.

The move sequence is represented by a string. And each move is represent by a character. The valid robot moves are R (Right), L (Left), U (Up) and D (down). The output should be true or false representing whether the robot makes a circle.

初始化一个机器人站在0,0坐标上，传递一个队列让这个机器人移动，最后就是判定移动后，这个机器人是否能回到原点。

我的方案是，设定两个坐标，上和下改变其中一个坐标的值 上 +1 下 -1 ，左右改变另一个坐标的值，左 +1 右 -1

然后遍历传入的字符串队列，最后判定两个坐标是否都是0 就 ok 了
*/
func main() {
	judgeCircle := judgeCircle("LL")
	fmt.Println("Output is", judgeCircle)
}

func judgeCircle(moves string) bool {
	leftPosition := 0
	rightPosition := 0
	for _, ch1 := range moves {
		switch ch1 {
		case 'U':
			leftPosition++
		case 'D':
			leftPosition--
		case 'L':
			rightPosition++
		case 'R':
			rightPosition--
		}
	}
	if leftPosition == 0 && rightPosition == 0 {
		return true
	} else {
		return false
	}
}

