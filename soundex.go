package phonetic

import (
	"errors"
	"fmt"
)

// ErrCharIsNotMapped tells that the given character could not be processed
// by soundex algorithm
var ErrCharIsNotMapped = errors.New("character is not mapped")

// https://nlp.stanford.edu/IR-book/html/htmledition/phonetic-correction-1.html
//
// charsMapping maps a letter to the corresponding digit
//
// 'A', 'E', 'I', 'O', 'U', 'H', 'W', 'Y' -> 0
// 'B', 'F', 'P', 'V' -> 1
// 'C', 'G', 'J', 'K', 'Q', 'S', 'X', 'Z' -> 2
// 'D', 'T' -> 3
// 'L' -> 4
// 'M', 'N' -> 5
// 'R' -> 6
var charsMapping = []rune("01230120022455012623010202")

type soundex struct {
}

// NewSoundexEncoder creates a new instance of soundex encoder
func NewSoundexEncoder() Encoder {
	return &soundex{}
}

// inspired by https://github.com/apache/commons-codec/blob/master/src/main/java/org/apache/commons/codec/language/Soundex.java
func (s soundex) Encode(source string) (string, error) {
	if len(source) == 0 {
		return "", nil
	}

	chars := []rune(source)
	lastDigit, err := mapChar(chars[0])

	if err != nil {
		return "", err
	}

	hash := [4]rune{chars[0], '0', '0', '0'}
	count := 1

	for i := 1; i < len(chars) && count < len(hash); i++ {
		digit, err := mapChar(chars[i])

		if err != nil {
			return "", err
		}

		if digit != '0' && digit != lastDigit {
			hash[count] = digit
			count++
		}

		lastDigit = digit
	}

	return string(hash[:]), nil
}

func mapChar(ch rune) (rune, error) {
	index := ch - 'A'

	if index < 0 || int(index) > len(charsMapping) {
		return 0, fmt.Errorf("%w: %#U", ErrCharIsNotMapped, ch)
	}

	return charsMapping[index], nil
}
