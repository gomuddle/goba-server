package goba_server

import (
	"encoding/base64"
	"errors"
	"strings"
)

// BASIC is the prefix of basic authorization headers.
const BASIC = "Basic"

var (
	// ErrEmptyAuthHeader occurs when an authorization header is empty.
	ErrEmptyAuthHeader    = errors.New("authorization header is empty")

	// ErrAuthHeaderNotBasic occurs when an authorization header is not basic.
	ErrAuthHeaderNotBasic = errors.New("authorization header type must be \"basic\"")

	// ErrInvalidAuthHeader occurs when an authorization header is invalid.
	ErrInvalidAuthHeader  = errors.New("invalid authorization header")
)

// credentialsFromHeader decodes header into Credentials.
func credentialsFromHeader(header string) (*Credentials, error) {
	if header == "" {
		return nil, ErrEmptyAuthHeader
	}
	if !isBasicAuth(header) {
		return nil, ErrAuthHeaderNotBasic
	}
	if len(header) < len(BASIC)+1 {
		return nil, ErrInvalidAuthHeader
	}
	return decodeCredentials(header[len(BASIC)+1:])
}

// isBasicAuth reports whether header is a basic authorization header.
func isBasicAuth(header string) bool {
	return strings.HasPrefix(header, BASIC)
}

// decodeCredentials decodes the given encoded credentials into Credentials.
func decodeCredentials(encCredentials string) (*Credentials, error) {
	rawCreds, err := base64.StdEncoding.DecodeString(encCredentials)
	if err != nil {
		return nil, err
	}
	creds := string(rawCreds)

	sepPos := strings.IndexRune(creds, ':')
	if sepPos < 0 {
		return nil, ErrInvalidAuthHeader
	}

	return &Credentials{
		Username: creds[:sepPos],
		Password: creds[sepPos+1:],
	}, nil
}
