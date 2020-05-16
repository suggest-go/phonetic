package phonetic

import (
	"testing"
)

func TestHandle(t *testing.T) {
	testCases := []struct {
		name           string
		chars          []rune
		index          int
		handler        func(b doubleMetaphoneResult, chars runes, index int) int
		expectedIndex  int
		expectedString string
	}{
		{
			name:           "handleBB",
			chars:          runes{'B', 'B'},
			index:          0,
			handler:        handleBFKNRQV,
			expectedIndex:  2,
			expectedString: "P",
		},
		{
			name:           "handleBP",
			chars:          runes{'B', 'P'},
			index:          0,
			handler:        handleBFKNRQV,
			expectedIndex:  1,
			expectedString: "P",
		},
		{
			name:           "handleNN",
			chars:          runes{'N', 'N'},
			index:          0,
			handler:        handleBFKNRQV,
			expectedIndex:  2,
			expectedString: "N",
		},
		{
			name:           "handleDG",
			chars:          runes{'E', 'D', 'G', 'E'},
			index:          1,
			handler:        handleD,
			expectedIndex:  4,
			expectedString: "J",
		},
		{
			name:           "handleDG",
			chars:          runes{'E', 'D', 'G', 'A', 'R'},
			index:          1,
			handler:        handleD,
			expectedIndex:  3,
			expectedString: "TK",
		},
		{
			name:           "handleDT",
			chars:          runes{'D', 'T'},
			index:          0,
			handler:        handleD,
			expectedIndex:  2,
			expectedString: "T",
		},
		{
			name:           "handleD",
			chars:          runes{'D', 'K'},
			index:          0,
			handler:        handleD,
			expectedIndex:  1,
			expectedString: "T",
		},
		{
			name:           "handleH before vowels",
			chars:          runes{'A', 'H', 'E'},
			index:          1,
			handler:        handleH,
			expectedIndex:  3,
			expectedString: "H",
		},
		{
			name:           "handleH first and next vowel",
			chars:          runes{'H', 'E', 'Y'},
			index:          0,
			handler:        handleH,
			expectedIndex:  2,
			expectedString: "H",
		},
		{
			name:           "handleH ignore",
			chars:          runes{'W', 'H', 'A', 'T'},
			index:          1,
			handler:        handleH,
			expectedIndex:  2,
			expectedString: "",
		},
		{
			name:           "handlePH",
			chars:          runes{'P', 'H', 'A', 'R'},
			index:          0,
			handler:        handleP,
			expectedIndex:  2,
			expectedString: "F",
		},
		{
			name:           "handleP",
			chars:          runes{'P', 'O', 'M'},
			index:          0,
			handler:        handleP,
			expectedIndex:  1,
			expectedString: "P",
		},
		{
			name:           "handlePP",
			chars:          runes{'P', 'P', 'O'},
			index:          0,
			handler:        handleP,
			expectedIndex:  2,
			expectedString: "P",
		},
		{
			name:           "handleWR",
			chars:          runes{'W', 'R', 'A', 'N', 'K'},
			index:          0,
			handler:        handleW,
			expectedIndex:  2,
			expectedString: "R",
		},
		{
			name:           "handleWPolish",
			chars:          runes{'F', 'I', 'L', 'I', 'P', 'O', 'W', 'I', 'C', 'Z'},
			index:          6,
			handler:        handleW,
			expectedIndex:  10,
			expectedString: "TS",
		},
		{
			name:           "handleX",
			chars:          runes{'X', 'E', 'R'},
			index:          0,
			handler:        handleX,
			expectedIndex:  1,
			expectedString: "S",
		},
		{
			name:           "handleXaux",
			chars:          runes{'B', 'R', 'E', 'A', 'U', 'X'},
			index:          5,
			handler:        handleX,
			expectedIndex:  6,
			expectedString: "",
		},
		{
			name:           "handleX",
			chars:          runes{'R', 'E', 'X'},
			index:          2,
			handler:        handleX,
			expectedIndex:  3,
			expectedString: "KS",
		},
		{
			name:           "handleXX",
			chars:          runes{'R', 'E', 'X', 'X'},
			index:          2,
			handler:        handleX,
			expectedIndex:  4,
			expectedString: "KS",
		},
		{
			name:           "handleZH",
			chars:          runes{'Z', 'H', 'A', 'G'},
			index:          0,
			handler:        handleZ,
			expectedIndex:  2,
			expectedString: "J",
		},
		{
			name:           "handleZZ",
			chars:          runes{'Z', 'Z', 'E', 'R', 'O'},
			index:          0,
			handler:        handleZ,
			expectedIndex:  2,
			expectedString: "S",
		},
		{
			name:           "handleZ",
			chars:          runes{'Z', 'E', 'R', 'O'},
			index:          0,
			handler:        handleZ,
			expectedIndex:  1,
			expectedString: "S",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			b := newDoubleMetaphoneResult()
			actualIndex := testCase.handler(b, testCase.chars, testCase.index)

			if actualIndex != testCase.expectedIndex {
				t.Errorf("test fail, expected %d, got %d", testCase.expectedIndex, actualIndex)
			}

			actualString := b.primary.String()

			if actualString != testCase.expectedString {
				t.Errorf("test fail, expected %s, got %s", testCase.expectedString, actualString)
			}
		})
	}
}
