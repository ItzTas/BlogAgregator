package auth

import (
	"errors"
	"net/http"
	"strings"
)

func ParseAPIKey(headers http.Header) (string, error) {
	h := headers.Get("Authorization")
	if h == "" {
		return "", errors.New("no authorization header")
	}
	splitedHeader := strings.Split(h, " ")
	if len(splitedHeader) != 2 || splitedHeader[0] != "ApiKey" {
		return "", errors.New("invalid header formating")
	}
	return splitedHeader[1], nil
}
