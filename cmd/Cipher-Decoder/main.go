package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	const a = 21
	const b = 15
	var g, p int
	var bKey = 1
	var sKey = 1
	var aKey int

	_, _ = fmt.Scanf("g is %d and p is %d", &g, &p)
	fmt.Println("OK")
	_, _ = fmt.Scanf("A is %d", &aKey)

	for i := 1; i <= b; i++ {
		bKey = (bKey * g) % p
		sKey = (sKey * aKey) % p
	}
	fmt.Printf("B is %d\n", bKey)

	fmt.Println(CaesarCipher("Will you marry me?", sKey))

	reader := bufio.NewReader(os.Stdin)
	ans, _ := reader.ReadBytes('\n')
	if string(ans) == CaesarCipher("Yeah, okay!\n", sKey) {
		fmt.Println(CaesarCipher("Great!", sKey))
	}
	if string(ans) == CaesarCipher("Let's be friends.\n", sKey) {
		fmt.Println(CaesarCipher("What a pity!", sKey))
	}
}

// CaesarCipher
// mes - message to cipher
// of - offset, how many characters to shift
func CaesarCipher(mes string, of int) string {
	res := make([]rune, len(mes))
	of = of % 26
	for i, letter := range mes {
		if !unicode.IsLetter(letter) {
			res[i] = letter
			continue
		}
		if unicode.IsUpper(letter) {
			if int(letter)+of <= 'Z' {
				res[i] = rune(int(letter) + of)
			} else {
				res[i] = rune(int(letter) + of - 26)
			}
		}
		if unicode.IsLower(letter) {
			if int(letter)+of <= 'z' {
				res[i] = rune(int(letter) + of)
			} else {
				res[i] = rune(int(letter) + of - 26)
			}
		}
	}
	return string(res)
}
