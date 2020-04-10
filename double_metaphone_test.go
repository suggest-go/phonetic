package phonetic

import (
	"strings"
	"testing"
)

func TestHandle(t *testing.T) {
	testCases := []struct {
		name           string
		chars          []rune
		index          int
		handler        func(b *strings.Builder, chars []rune, index int) int
		expectedIndex  int
		expectedString string
	}{
		{
			name:           "handleBB",
			chars:          []rune{'B', 'B'},
			index:          0,
			handler:        handleBFKNRQV,
			expectedIndex:  2,
			expectedString: "P",
		},
		{
			name:           "handleBP",
			chars:          []rune{'B', 'P'},
			index:          0,
			handler:        handleBFKNRQV,
			expectedIndex:  1,
			expectedString: "P",
		},
		{
			name:           "handleNN",
			chars:          []rune{'N', 'N'},
			index:          0,
			handler:        handleBFKNRQV,
			expectedIndex:  2,
			expectedString: "N",
		},
		{
			name:           "handleDG",
			chars:          []rune{'E', 'D', 'G', 'E'},
			index:          1,
			handler:        handleD,
			expectedIndex:  4,
			expectedString: "J",
		},
		{
			name:           "handleDG",
			chars:          []rune{'E', 'D', 'G', 'A', 'R'},
			index:          1,
			handler:        handleD,
			expectedIndex:  3,
			expectedString: "TK",
		},
		{
			name:           "handleDT",
			chars:          []rune{'D', 'T'},
			index:          0,
			handler:        handleD,
			expectedIndex:  2,
			expectedString: "T",
		},
		{
			name:           "handleD",
			chars:          []rune{'D', 'K'},
			index:          0,
			handler:        handleD,
			expectedIndex:  1,
			expectedString: "T",
		},
		{
			name:           "handleH before vowels",
			chars:          []rune{'A', 'H', 'E'},
			index:          1,
			handler:        handleH,
			expectedIndex:  3,
			expectedString: "H",
		},
		{
			name:           "handleH first and next vowel",
			chars:          []rune{'H', 'E', 'Y'},
			index:          0,
			handler:        handleH,
			expectedIndex:  2,
			expectedString: "H",
		},
		{
			name:           "handleH ignore",
			chars:          []rune{'W', 'H', 'A', 'T'},
			index:          1,
			handler:        handleH,
			expectedIndex:  2,
			expectedString: "",
		},
		{
			name:           "handlePH",
			chars:          []rune{'P', 'H', 'A', 'R'},
			index:          0,
			handler:        handleP,
			expectedIndex:  2,
			expectedString: "F",
		},
		{
			name:           "handleP",
			chars:          []rune{'P', 'O', 'M'},
			index:          0,
			handler:        handleP,
			expectedIndex:  1,
			expectedString: "P",
		},
		{
			name:           "handlePP",
			chars:          []rune{'P', 'P', 'O'},
			index:          0,
			handler:        handleP,
			expectedIndex:  2,
			expectedString: "P",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			b := &strings.Builder{}
			actualIndex := testCase.handler(b, testCase.chars, testCase.index)

			if actualIndex != testCase.expectedIndex {
				t.Errorf("test fail, expected %d, got %d", testCase.expectedIndex, actualIndex)
			}

			actualString := b.String()

			if actualString != testCase.expectedString {
				t.Errorf("test fail, expected %s, got %s", testCase.expectedString, actualString)
			}
		})
	}
}
