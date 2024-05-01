package types

// Raw JSON, used to bypass encoding/decoding, in case encoding/decoding is done out of band (e.g. with proto)
// or should be skipped for some reason
type rawJsonBytes struct {
	Bytes []byte
}

func (b *rawJsonBytes) MarshalJSON() ([]byte, error) {
	return b.Bytes, nil
}

func (b *rawJsonBytes) UnmarshalJSON(payload []byte) error {
	b.Bytes = payload
	return nil
}

func RawJsonBytes(bytes []byte) *rawJsonBytes {
	return &rawJsonBytes{bytes}
}
