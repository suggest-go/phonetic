package phonetic

// Encoder provides API for encoding purposes
type Encoder interface {
	// Encode encodes a string and returns an encoded result
	Encode(source string) (string, error)
}

// Decoder provides API for decoding purposes
type Decoder interface {
	// Decode decodes a string and returns a decoded result
	Decode(source string) (string, error)
}
