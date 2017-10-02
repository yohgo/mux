package mux

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/yohgo/pastry"
)

// Operation is a function the executes a specific operation on a request
type Operation func(w http.ResponseWriter, r *http.Request, parameters ...interface{}) bool

// ValidateJWT is an operation that validate the JWT token on a request
func ValidateJWT(w http.ResponseWriter, r *http.Request, parameters ...interface{}) bool {
	var resolver pastry.JSONResolver

	// Attempting to load the HMAC key
	HMACKey, ok := parameters[0].(string)
	if HMACKey == "" || !ok {
		resolver.ResolveJSONError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), "Access to the requested resource is denied")
		return false
	}

	// Attempting to read the JWT access token string
	accessToken := map[bool]string{true: r.FormValue("access_token"), false: r.Header.Get("access_token")}[r.Header.Get("access_token") == ""]
	if accessToken == "" {
		resolver.ResolveJSONError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), "Access to the requested resource is denied")
		return false
	}

	// Attempting to parse the JWT access token string
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}

		return []byte(HMACKey), nil
	})

	if err != nil {
		resolver.ResolveJSONError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), "Access to the requested resource is denied")
		return false
	}

	claims := token.Claims.(jwt.MapClaims)

	userID, ok := claims["uid"].(uint64)
	if !ok {
		resolver.ResolveJSONError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), "Access to the requested resource is denied")
		return false
	}

	userRole, ok := claims["iur"].(string)
	if !ok {
		resolver.ResolveJSONError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), "Access to the requested resource is denied")
		return false
	}

	userPermissions, ok := claims["iup"].([]string)
	if !ok {
		resolver.ResolveJSONError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), "Access to the requested resource is denied")
		return false
	}

	r.Header.Add("user_id", fmt.Sprintf("%d", userID))
	r.Header.Add("user_role", userRole)
	r.Header.Add("user_permissions", strings.Join(userPermissions, ","))

	return true
}
