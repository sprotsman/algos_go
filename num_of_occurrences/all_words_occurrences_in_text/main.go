/*
Count how many items each word appears in a text.
Use the "Fields" function to split text into words.
Use the ToLower function to convert all words to lowercase.
Print the map.
*/
package main

import (
	"fmt"
	"strings"
)

const text = `
When, in disgrace with fortune and men's eyes,
I all alone beweep my outcast state,
And trouble deaf heaven with my bootless cries,
And look upon myself, and curse my fate,
Wishing me like to one more rich in hope,
Featured like him, like him with friends possessed,
Desiring this man's art and that man's scope,
With what I most enjoy contented least;
Yet in these thoughts myself almost despising,
Haply I think on thee - and then my state,
Like to the lark at break of day arising
From sullen earth, sings hymns at heaven's gate;
For thy sweet love rememb'red such wealth brings
That then I scorn to change my state with kings.
`

func main() {
	// create slice of strings
	words := strings.Fields(text)
	// create empty map from word to the count of the word
	counts := map[string]int{}  // word -> count
	for _, word := range words {
		counts[strings.ToLower(word)]++
	}
	fmt.Println(counts)
}
