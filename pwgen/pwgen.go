// proutils-api/pwgen/pwgen.go
//
// Logic for generating passwords.

package pwgen

import (
	"crypto/rand"
	"math"
	"math/big"
	"unicode"
)

var (
	// digits is the number of digits to roll. This is determined by the
	// dictionary, in this case the "EFF Large Wordlist for Passphrases".
	digits = 5

	// sides is the number of sides on a die
	sides = big.NewInt(6)
)

// one-liner for standard error handling
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

// NewPassword takes the number of words and other flags to return single string password
func NewPassword(words int, addInt bool, addSpecial bool) (string, error) {
	var pw string

	// Get list of random words
	wordList, err := Generate(words)
	handleError(err)

	// Capitalize the first letter of each word
	for _, word := range wordList {
		// TODO: Look at using byte slice with copy for better performance
		pw += upperFirst(word)
	}

	// If requested, append a number to the end
	if addInt {
		i, err := rand.Int(rand.Reader, big.NewInt(9))
		handleError(err)
		pw += i.String()
	}

	// If requested, append a special character to the end
	if addSpecial {
		specials := "~=+%^*/()[]{}/!@#$?|" // 20 characters
		i, err := rand.Int(rand.Reader, big.NewInt(19))
		if err != nil {
			return "", err
		}
		pw += string(specials[i.Int64()])
	}

	return pw, nil
}

// upperFirst takes a string and returns it with the first letter capitalized
func upperFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// Generate generates a list of the given number of words.
func Generate(words int) ([]string, error) {
	list := make([]string, 0, words)
	seen := make(map[string]struct{}, words)

	for i := 0; i < words; i++ {
		n, err := RollWord(digits)
		if err != nil {
			return nil, err
		}

		word := WordAt(n)
		if _, ok := seen[word]; ok {
			i--
			continue
		}

		list = append(list, word)
		seen[word] = struct{}{}
	}

	return list, nil
}

// MustGenerate behaves like Generate, but panics on error.
func MustGenerate(words int) []string {
	res, err := Generate(words)
	if err != nil {
		panic(err)
	}
	return res
}

// WordAt retrieves the word at the given index.
func WordAt(i int) string {
	return words[i]
}

// RollDie rolls a single 6-sided die and returns a value between [1,6].
func RollDie() (int, error) {
	r, err := rand.Int(rand.Reader, sides)
	if err != nil {
		return 0, err
	}
	return int(r.Int64()) + 1, nil
}

// RollWord rolls and aggregates dice to represent one word in the list. The
// result is the index of the word in the list.
func RollWord(d int) (int, error) {
	var final int

	for i := d; i > 0; i-- {
		res, err := RollDie()
		if err != nil {
			return 0, err
		}

		final += res * int(math.Pow(10, float64(i-1)))
	}

	return final, nil
}
