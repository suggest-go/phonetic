package phonetic

import "strings"

const (
	maxCodeLen = 4
	vowels     = "AEIOUY"
)

type runes []rune

func (c runes) at(index int) rune {
	ch := rune(-1 << 31)

	if index < len(c) && index >= 0 {
		ch = c[index]
	}

	return ch
}

func (c runes) contains(start, length int, criteria ...string) bool {
	result := false

	if start >= 0 && start+length <= len(c) {
		target := c[start : start+length]

		for _, candidate := range criteria {
			if equal(target, candidate) {
				result = true
				break
			}
		}
	}

	return result
}

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
	r.appendStringPrimary(str)
	r.appendStringAlternate(str)
}

func (r doubleMetaphoneResult) appendStringPrimary(str string) {
	r.primary.WriteString(str)
}

func (r doubleMetaphoneResult) appendStringAlternate(str string) {
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

func handleBFKNRQV(b doubleMetaphoneResult, chars runes, index int) int {
	ch := chars[index]
	mapped := mapBFKNRQV[ch]
	curr, next := mapped[0], mapped[1]

	b.append(curr)

	if chars.at(index+1) == next {
		index += 2
	} else {
		index++
	}

	return index
}

func handleC(b doubleMetaphoneResult, chars runes, index int) int {
	return index
}

func handleD(b doubleMetaphoneResult, chars runes, index int) int {
	if chars.contains(index, 2, "DG") {
		if chars.contains(index+2, 1, "I", "E", "Y") {
			b.append('J')
			index += 3
		} else {
			b.appendString("TK")
			index += 2
		}
	} else if chars.contains(index, 2, "DT", "DD") {
		b.append('T')
		index += 2
	} else {
		b.append('T')
		index++
	}

	return index
}

func handleG(b doubleMetaphoneResult, chars runes, index int) int {
	return index
}

func handleH(b doubleMetaphoneResult, chars runes, index int) int {
	if (index == 0 || isVowel(chars.at(index-1))) && isVowel(chars.at(index+1)) {
		b.append('H')
		index += 2
	} else {
		index++
	}

	return index
}

func handleJ(b doubleMetaphoneResult, chars runes, index int) int {
	return index
}

func handleL(b doubleMetaphoneResult, chars runes, index int) int {
	return index
}

func handleP(b doubleMetaphoneResult, chars runes, index int) int {
	if chars.at(index+1) == 'H' {
		b.append('F')
		index += 2
	} else {
		b.append('P')

		if chars.contains(index+1, 1, "P", "B") {
			index++
		}

		index++
	}

	return index
}

func handleS(b doubleMetaphoneResult, chars runes, index int) int {
	return index
}

func handleT(b doubleMetaphoneResult, chars runes, index int) int {
	return index
}

func handleW(b doubleMetaphoneResult, chars runes, index int) int {
	if chars.contains(index, 2, "WR") {
		b.append('R')
		index += 2
	} else {
		if index == 0 && (isVowel(chars.at(index+1)) || chars.contains(index, 2, "WH")) {
			if isVowel(chars.at(index + 1)) {
				b.appendPrimary('A')
				b.appendAlternate('F')
			} else {
				b.append('A')
			}

			index++
		} else if index == len(chars)-1 && isVowel(chars.at(index-1)) ||
			chars.contains(index-1, 5, "EWSKI", "EWSKY", "OWSKI", "OWSKY") ||
			chars.contains(0, 3, "SCH") {
			b.appendAlternate('F')
			index++
		} else if chars.contains(index, 4, "WICZ", "WITZ") {
			b.appendStringPrimary("TS")
			b.appendStringAlternate("FX")
			index += 4
		} else {
			index++
		}
	}

	return index
}

func handleX(b doubleMetaphoneResult, chars runes, index int) int {
	if index == 0 {
		b.append('S')
		index++
	} else {
		if !(index == len(chars)-1 &&
			(chars.contains(index-3, 3, "IAU", "EAU") || chars.contains(index-2, 2, "AU", "OU"))) {
			b.appendString("KS")
		}

		if chars.contains(index+1, 1, "C", "X") {
			index++
		}

		index++
	}

	return index
}

func handleZ(b doubleMetaphoneResult, chars runes, index int) int {
	if chars.at(index+1) == 'H' {
		b.append('J')
		index += 2
	} else {
		if chars.contains(index+1, 2, "ZO", "ZI", "ZA") {
			b.appendString("STS")
		} else {
			b.append('S')
		}

		if chars.at(index+1) == 'Z' {
			index += 2
		} else {
			index++
		}
	}

	return index
}

func equal(a []rune, b string) bool {
	if len(a) != len(b) {
		return false
	}

	result := true

	for i, ch := range b {
		result = ch == a[i]

		if !result {
			break
		}
	}

	return result
}

func isVowel(ch rune) bool {
	return strings.ContainsRune(vowels, ch)
}
