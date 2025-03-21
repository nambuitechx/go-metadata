package models

import "errors"

// Verify ssl
// Client SSL verification. Make sure to configure the SSLConfig if enabled.
var VerifySSL = map[string]int {"no-ssl": 0, "ignore": 1, "validate": 2}

func ValidateVerifySSL(verifySSL string) (int, error) {
	idx, ok := VerifySSL[verifySSL]

	if !ok {
		return -1, errors.New("invalid verifty ssl")
	}

	return idx, nil
}

// SSL Mode
// SSL Mode to connect to database.
var SSLMode = map[string]int {"disable": 0, "allow": 1, "prefer": 2, "require": 3, "verify-ca": 4, "verify-full": 5}

func ValidateSSLMode(sslMode string) (int, error) {
	idx, ok := SSLMode[sslMode]

	if !ok {
		return -1, errors.New("invalid ssl mode")
	}

	return idx, nil
}
