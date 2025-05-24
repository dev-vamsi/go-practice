package main

import "fmt"

func main() {
	myText := TextObject{text: "Vamsi Krishna Tupakula"}
	fmt.Println(myText.toLowerCase())
	fmt.Println(myText.toUpperCase())
	myText.printWordCount()
	myText.printCharacterCount()
	fmt.Println(myText.getReversedString())
	fmt.Println(myText.isPalindrome())
}
