package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAuthAPI(header http.Header) (string, error) {
	val := header.Get("Authorization")

	if val == "" {
		return "", errors.New("No authetication found")
	}

	values := strings.Split(val, " ")

	if len(values) != 2 || values[0] != "ApiKey" {
		return "", errors.New("Malformed authentication")
	}

	return values[1], nil
}
