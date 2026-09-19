package protocol

import (
	"encoding/hex"
	"strings"
)

func Encode(text string) string {
	return hex.EncodeToString([]byte(text))
}

func Decode(encoded string) (string, error) {
	clean := strings.TrimSpace(encoded)

	decoded, err := hex.DecodeString(clean)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}
