/* --------------------------------------------------------------------------
 * APUNTES:
 * 		   Para demostrar el uso de las funciones <toLower>, <toUpper>,
 *		   <Replace>, <ReplaceAll>, <Split> y <Join> se verificara si una
 *		   palabra es palíndroma ó no.
 *
 *
 * IMPORTANTE:
 *  			  - <Split(<STRING>, <SEPARATOR>)> = Convierte un <String> a un <Array>.
 *  			  - <Join(<ARRAY>, <SEPARATOR>)> = Convierte un <Array> a un <String>.
-------------------------------------------------------------------------- */

package main

import (
	"fmt"
	"strings"
)

func reverseString(word string) string {
	wordArray := strings.Split(word, "")
	reverseWordArray := make([]string, 0)

	for i := len(wordArray) - 1; i >= 0; i-- {
		reverseWordArray = append(reverseWordArray, wordArray[i])
	}

	return strings.Join(reverseWordArray, "")
}

func isPalindrome(word string) bool {
	word = strings.ToLower(word)
	// word = strings.Replace(word, "") // <Replace(<STRING>, <OLD CHARACTER>, <NEW CHARACTER>, <HOW MUCH OLD CHARACTERS REPLACE>)>
	word = strings.ReplaceAll(word, " ", "") // <ReplaceAll(<STRING>, <OLD CHARACTER>, <NEW CHARACTER>)>
	reverseWord := reverseString(word)

	return word == reverseWord
}

func main() {
	var word string = "Luz Azul"

	if isPalindrome(word) {
		fmt.Println("¡Es palíndroma!")
	} else {
		fmt.Println("¡No es palíndroma!")
	}
}
