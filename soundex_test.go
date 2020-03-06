package phonetic

import (
	"errors"
	"fmt"
	"testing"
)

func TestSoundex(t *testing.T) {
	testCases := []struct {
		source string
		hash   string
		err    error
	}{
		{
			source: "WILLIAMS",
			hash:   "W452",
			err:    nil,
		},
		{
			source: "BARAGWANATH",
			hash:   "B625",
			err:    nil,
		},
		{
			source: "DONNELL",
			hash:   "D540",
			err:    nil,
		},
		{
			source: "LLOYD",
			hash:   "L300",
			err:    nil,
		},
		{
			source: "WOOLCOCK",
			hash:   "W422",
			err:    nil,
		},
		{
			source: "ROBERT",
			hash:   "R163",
			err:    nil,
		},
		{
			source: "RUPERT",
			hash:   "R163",
			err:    nil,
		},
		{
			source: "MICHAEL",
			hash:   "M240",
			err:    nil,
		},
		{
			source: "test",
			hash:   "T230",
			err:    nil,
		},
		{
			source: "TeST",
			hash:   "T230",
			err:    nil,
		},
		{
			source: "brown",
			hash:   "B650",
			err:    nil,
		},
	}

	for i, testCase := range testCases {
		testCase := testCase

		t.Run(fmt.Sprintf("testCase #%d", i+1), func(t *testing.T) {
			encoder := NewSoundexEncoder()
			hash, err := encoder.Encode(testCase.source)

			if !errors.Is(err, testCase.err) {
				t.Errorf("expected error %v, got %v", testCase.err, err)
			}

			if hash != testCase.hash {
				t.Errorf("expected hash %v, got %v", testCase.hash, hash)
			}
		})
	}
}

func BenchmarkSoundex(b *testing.B) {
	words := []string{
		"inpdendence",
		"approxmiation",
		"testing",
		"accuracy",
	}

	encoder := NewSoundexEncoder()

	for i := 0; i < b.N; i++ {
		source := words[i%len(words)]

		if _, err := encoder.Encode(source); err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}
