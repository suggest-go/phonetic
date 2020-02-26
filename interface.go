package phonetic

// Encoder provides API for encoding purposes
type Encoder interface {
	// Encode encodes a string and returns an encoded result
	Encode(source string) (string, error)
}
