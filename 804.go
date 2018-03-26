package main

import (
	"bytes"
	"fmt"
)

/*
804. Unique Morse Code Words

International Morse Code defines a standard encoding where each letter is mapped to a series of dots and dashes, as follows: "a" maps to ".-", "b" maps to "-...", "c" maps to "-.-.", and so on.

For convenience, the full table for the 26 letters of the English alphabet is given below:

[".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."]
Now, given a list of words, each word can be written as a concatenation of the Morse code of each letter. For example, "cab" can be written as "-.-.-....-", (which is the concatenation "-.-." + "-..." + ".-"). We'll call such a concatenation, the transformation of a word.

Return the number of different transformations among all words we have.

Example:
Input: words = ["gin", "zen", "gig", "msg"]
Output: 2
Explanation:
The transformation of each word is:
"gin" -> "--...-."
"zen" -> "--...-."
"gig" -> "--...--."
"msg" -> "--...--."

There are 2 different transformations, "--...-." and "--...--.".


思路： 做好映射后，依次比较就好了，如果在已经存在的 slice 就不加入，最后返回 slice 的长度即可
*/

func main() {
	root := []string{"gin", "zen", "gig", "msg"}
	uniqueMorseRepresentations := uniqueMorseRepresentations(root)
	fmt.Println("uniqueMorseRepresentations:", uniqueMorseRepresentations)
}

func uniqueMorseRepresentations(words []string) int {
	morseMapping := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	uniqueMorseSlice := []string{}
	for _, str := range words {
		strBuffer := new(bytes.Buffer)
		for _, char := range str {
			strBuffer.WriteString(morseMapping[char-97])
		}
		tempStr := strBuffer.String()
		hasStr := sliceHasStr(uniqueMorseSlice, tempStr)
		if hasStr == false {
			uniqueMorseSlice = append(uniqueMorseSlice, tempStr)
		}
	}

	return len(uniqueMorseSlice)
}

func sliceHasStr(strSlice []string, checkStr string) bool {
	hasThisStr := false
	for _, str := range strSlice {
		if str == checkStr {
			hasThisStr = true
			break
		}
	}

	return hasThisStr
}

