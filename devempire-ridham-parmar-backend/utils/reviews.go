package utils

import (
	"net/url"
)

func DecodeString(encodedString string) (string, error) {
	decodedString, err := url.QueryUnescape(encodedString)
	return decodedString, err
}
