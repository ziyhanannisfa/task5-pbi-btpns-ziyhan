package helpers

import (
	"net/http"
	"strings"
)

func GetTokenFromHeader(request *http.Request) string {
	tokenHeader := request.Header.Get("Authorization")

	if tokenHeader == "" {
		return ""
	}

	tokenParts := strings.Split(tokenHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return ""
	}

	return tokenParts[1]
}
