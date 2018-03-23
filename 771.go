package main

import "fmt"

/*
You're given strings J representing the types of stones that are jewels, and S representing the stones you have.  Each character in S is a type of stone you have.  You want to know how many of the stones you have are also jewels.

The letters in J are guaranteed distinct, and all characters in J and S are letters. Letters are case sensitive, so "a" is considered a different type of stone from "A".

你得到的字符串J代表的是宝石的类型，S代表你拥有的宝石。 S中的每个字符都是你拥有的一种石头。你想知道你有多少宝石也是宝石。

J中的字母保证不同，J和S中的所有字符都是字母。字母区分大小写，所以“a”被认为是与“A”不同类型的石头。
*/
func main() {
	numJewelsInStones := numJewelsInStones("aA", "aAAbbbb")
	fmt.Println("Output is", numJewelsInStones)
}

func numJewelsInStones(J string, S string) int {
	count := 0
	for _, ch1 := range S {
		for _, ch2 := range J {
			if ch1 == ch2 {
				count++
			}
		}
	}
	return count
}

