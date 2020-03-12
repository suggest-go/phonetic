package phonetic

import "strings"

const maxCodeLen = 4

type doubleMetaphone struct {
}

// inspired by https://github.com/apache/commons-codec/blob/master/src/main/java/org/apache/commons/codec/language/DoubleMetaphone.java
func (d doubleMetaphone) Encode(source string) (string, error) {
	chars := clean(source)

	if len(chars) == 0 {
		return "", nil
	}

	b := &strings.Builder{}
	index := 0

	for index < len(chars) && b.Len() < maxCodeLen {
		switch ch := chars[index]; ch {
		case 'A', 'E', 'I', 'O', 'U', 'Y':
			index = handleAEIOUY(b, index)
		case 'B', 'F', 'K', 'N', 'Q', 'V':
			index = handleBFKNQV(b, chars, index)
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
			b.WriteRune('M')
			// TODO finish me
		case 'P':
			index = handleP(b, chars, index)
		case 'R':
			index = handleR(b, chars, index)
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

	return "", nil
}

func handleAEIOUY(b *strings.Builder, index int) int {
	if index == 0 {
		b.WriteRune('A')
	}

	return index + 1
}

var mapBFKNQV = map[rune][2]rune{
	'B': [2]rune{'P', 'B'},
	'F': [2]rune{'F', 'F'},
	'K': [2]rune{'K', 'K'},
	'N': [2]rune{'N', 'N'},
	'Q': [2]rune{'K', 'Q'},
	'V': [2]rune{'F', 'V'},
}

func handleBFKNQV(b *strings.Builder, chars []rune, index int) int {
	ch := chars[index]
	curr, next := mapBFKNQV[ch]

	b.WriteRune(curr)

	if index+1 < len(chars) && chars[index+1] == next {
		index += 2
	} else {
		index++
	}

	return index
}

func handleC(b *strings.Builder, chars []rune, index int) int {

}

func handleD(b *strings.Builder, chars []rune, index int) int {
	if contains(chars, index, 2, []rune{'D', 'G'}) {
		if contains(chars, index+2, 1, []rune{'I'}, []rune{'E'}, []rune{'Y'}) {
			b.WriteRune('J')
			index += 3
		} else {
			b.WriteString("TK")
			index += 2
		}
	} else if contains(chars, index, 2, []rune{'D', 'T'}, []rune{'D', 'D'}) {
		b.WriteRune('T')
		index += 2
	} else {
		b.WriteRune('T')
		index++
	}

	return index
}

func handleG(b *strings.Builder, chars []rune, index int) int {

}

func handleH(b *strings.Builder, chars []rune, index int) int {
}

func handleJ(b *strings.Builder, chars []rune, index int) int {

}

func handleL(b *strings.Builder, chars []rune, index int) int {

}

func handleP(b *strings.Builder, chars []rune, index int) int {

}

func handleR(b *strings.Builder, chars []rune, index int) int {

}

func handleS(b *strings.Builder, chars []rune, index int) int {

}

func handleT(b *strings.Builder, chars []rune, index int) int {

}

func handleW(b *strings.Builder, chars []rune, index int) int {

}

func handleX(b *strings.Builder, chars []rune, index int) int {

}

func handleZ(b *strings.Builder, chars []rune, index int) int {

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
