package cryptop

import (
	"encoding/hex"
	"errors"
	"regexp"
	"strings"
)

// cipher direction
const (
	DirectionEncrypt = iota
	DirectionDecrypt
)

// validation modes
const (
	ModeGenerate = iota
	ModeValidate
)

// encryption modes
const (
	ModeECB = "ECB"
	ModeCBC = "CBC"
)

// Pack returns the packed representation of an expanded hex string which is provided
func Pack(input string) (string, error) {

	// validate the input
	if len(input) == 0 {
		return "", errors.New("Input string has zero length")
	}

	if len(input)%2 != 0 {
		return "", errors.New("Input string is an uneven length, use only full bytes")
	}

	upperInput := strings.ToUpper(input)

	match, err := regexp.MatchString("^[0-9A-F]+$", upperInput)

	if err != nil {
		return "", err
	}

	if match == false {
		return "", errors.New("Input string contains invalid characters, use hex only (0-9 A-F)")
	}

	// decode!
	result, err := hex.DecodeString(upperInput)

	if err != nil {
		return "", errors.New("Failed to decode the given hex")
	}

	return string(result), nil
}

// Expand returns the expanded representation of the packed hex bytes which are provided
func Expand(input []byte) (string, error) {

	// validate the input
	if len(input) == 0 {
		return "", errors.New("Input string has zero length")
	}

	// encode!
	result := hex.EncodeToString(input)

	if len(result) == 0 {
		return "", errors.New("Failed to encode the given hex string")
	}

	return strings.ToUpper(result), nil
}

// XOR returns the XOR result of the two given byte slices, they are expected to have valid hexadecimal contents, equal lengths, and packed
func XOR(i1, i2 []byte) []byte {

	length := len(i1)
	r := make([]byte, length)

	for i := 0; i < length; i++ {
		r[i] = i1[i] ^ i2[i]
	}

	return r
}

// NumericOnly examines the given string and returns true if all bytes are numeric characters
func NumericOnly(s string) bool {

	match, err := regexp.MatchString("^[0-9]+$", s)

	if err != nil {
		return false
	}

	return match
}
