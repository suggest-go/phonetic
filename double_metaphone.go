package phonetic

import "strings"

const (
	maxCodeLen = 4
	vowels     = "AEIOUY"
)

type doubleMetaphone struct {
}

type doubleMetaphoneResult struct {
	primary   *strings.Builder
	alternate *strings.Builder
}

func newDoubleMetaphoneResult() doubleMetaphoneResult {
	return doubleMetaphoneResult{
		&strings.Builder{},
		&strings.Builder{},
	}
}

func (r doubleMetaphoneResult) append(ch rune) {
	r.appendPrimary(ch)
	r.appendAlternate(ch)
}

func (r doubleMetaphoneResult) appendPrimary(ch rune) {
	r.primary.WriteRune(ch)
}

func (r doubleMetaphoneResult) appendAlternate(ch rune) {
	r.alternate.WriteRune(ch)
}

func (r doubleMetaphoneResult) appendString(str string) {
	r.primary.WriteString(str)
	r.alternate.WriteString(str)
}

func (r doubleMetaphoneResult) string() string {
	return r.primary.String() + "#" + r.alternate.String()
}

func (r doubleMetaphoneResult) isComplete() bool {
	return r.primary.Len() >= maxCodeLen && r.alternate.Len() >= maxCodeLen
}

// inspired by https://github.com/apache/commons-codec/blob/master/src/main/java/org/apache/commons/codec/language/DoubleMetaphone.java
func (d doubleMetaphone) Encode(source string) (string, error) {
	chars := clean(source)

	if len(chars) == 0 {
		return "", nil
	}

	b := newDoubleMetaphoneResult()
	index := 0

	for index < len(chars) && !b.isComplete() {
		switch ch := chars[index]; ch {
		case 'A', 'E', 'I', 'O', 'U', 'Y':
			index = handleAEIOUY(b, index)
		case 'B', 'F', 'K', 'N', 'R', 'Q', 'V':
			index = handleBFKNRQV(b, chars, index)
		case 'C':
			index = handleC(b, chars, index)
		case 'D':
			index = handleD(b, chars, index)
		case 'G':
			index = handleG(b, chars, index)
		case 'H':
			index = handleH(b, chars, index)
		case 'J':
			index = handleJ(b, chars, index)
		case 'L':
			index = handleL(b, chars, index)
		case 'M':
			b.append('M')
			// TODO finish me
		case 'P':
			index = handleP(b, chars, index)
		case 'S':
			index = handleS(b, chars, index)
		case 'T':
			index = handleT(b, chars, index)
		case 'W':
			index = handleW(b, chars, index)
		case 'X':
			index = handleX(b, chars, index)
		case 'Z':
			index = handleZ(b, chars, index)
		default:
			index++
		}
	}

	return b.string(), nil
}

func handleAEIOUY(b doubleMetaphoneResult, index int) int {
	if index == 0 {
		b.append('A')
	}

	return index + 1
}

var mapBFKNRQV = map[rune][]rune{
	'B': {'P', 'B'},
	'F': {'F', 'F'},
	'K': {'K', 'K'},
	'N': {'N', 'N'},
	'R': {'R', 'R'},
	'Q': {'K', 'Q'},
	'V': {'F', 'V'},
}

func handleBFKNRQV(b doubleMetaphoneResult, chars []rune, index int) int {
	ch := chars[index]
	mapped := mapBFKNRQV[ch]
	curr, next := mapped[0], mapped[1]

	b.append(curr)

	if index+1 < len(chars) && chars[index+1] == next {
		index += 2
	} else {
		index++
	}

	return index
}

func handleC(b doubleMetaphoneResult, chars []rune, index int) int {
	return index
}

func handleD(b doubleMetaphoneResult, chars []rune, index int) int {
	if contains(chars, index, 2, []rune{'D', 'G'}) {
		if contains(chars, index+2, 1, []rune{'I'}, []rune{'E'}, []rune{'Y'}) {
			b.append('J')
			index += 3
		} else {
			b.appendString("TK")
			index += 2
		}
	} else if contains(chars, index, 2, []rune{'D', 'T'}, []rune{'D', 'D'}) {
		b.append('T')
		index += 2
	} else {
		b.append('T')
		index++
	}

	return index
}

func handleG(b doubleMetaphoneResult, chars []rune, index int) int {
	return index
}

func handleH(b doubleMetaphoneResult, chars []rune, index int) int {
	if (index == 0 || isVowel(chars[index-1])) && index+1 < len(chars) && isVowel(chars[index+1]) {
		b.append('H')
		index += 2
	} else {
		index++
	}

	return index
}

func handleJ(b doubleMetaphoneResult, chars []rune, index int) int {
	return index
}

func handleL(b doubleMetaphoneResult, chars []rune, index int) int {
	return index
}

func handleP(b doubleMetaphoneResult, chars []rune, index int) int {
	if index+1 < len(chars) && chars[index+1] == 'H' {
		b.append('F')
		index += 2
	} else {
		b.append('P')

		if contains(chars, index+1, 1, []rune{'P'}, []rune{'B'}) {
			index++
		}

		index++
	}

	return index
}

func handleS(b doubleMetaphoneResult, chars []rune, index int) int {
	return index
}

func handleT(b doubleMetaphoneResult, chars []rune, index int) int {
	return index
}

func handleW(b doubleMetaphoneResult, chars []rune, index int) int {
	return index
}

func handleX(b doubleMetaphoneResult, chars []rune, index int) int {
	return index
}

func handleZ(b doubleMetaphoneResult, chars []rune, index int) int {
	if index+1 < len(chars) && chars[index+1] == 'H' {
		b.append('J')
		index += 2
	} else {
		if contains(chars, index+1, 2, []rune{'Z', 'O'}, []rune{'Z', 'I'}, []rune{'Z', 'A'}) {
			b.appendString("STS")
		} else {
			b.append('S')
		}

		if index+1 < len(chars) && chars[index+1] == 'Z' {
			index += 2
		} else {
			index++
		}
	}

	return index
}

func contains(chars []rune, start, length int, criteria ...[]rune) bool {
	result := false

	if start >= 0 && start+length <= len(chars) {
		target := chars[start : start+length]

		for _, candidate := range criteria {
			if equal(target, candidate) {
				result = true
				break
			}
		}
	}

	return result
}

func equal(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}

	result := true

	for i := 0; i < len(a) && result; i++ {
		result = a[i] == b[i]
	}

	return result
}

func isVowel(ch rune) bool {
	return strings.ContainsRune(vowels, ch)
}
