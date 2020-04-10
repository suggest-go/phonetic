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
			handler:        handleBFKNQV,
			expectedIndex:  2,
			expectedString: "P",
		},
		{
			name:           "handleBP",
			chars:          []rune{'B', 'P'},
			index:          0,
			handler:        handleBFKNQV,
			expectedIndex:  1,
			expectedString: "P",
		},
		{
			name:           "handleNN",
			chars:          []rune{'N', 'N'},
			index:          0,
			handler:        handleBFKNQV,
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
