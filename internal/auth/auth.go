package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/ItsTas/BlogAgregator/internal/client"
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

func ExtractImageUrl(item client.Item) string {
	for _, media := range item.Media {
		if media.Type == "image/jpeg" || media.Type == "image/png" {
			return media.URL
		}
	}

	for _, thumb := range item.Thumbnail {
		if thumb.Type == "image/jpeg" || thumb.Type == "image/png" {
			return thumb.URL
		}
	}

	for _, enclosure := range item.Enclosure {
		if enclosure.Type == "image/jpeg" || enclosure.Type == "image/png" {
			return enclosure.URL
		}
	}

	return ""
}
