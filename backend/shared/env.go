package shared

import (
	"errors"
	"os"
	"strings"
)

var DevMode bool
var AuthServiceHostname string
var JwtSecret string

func getDevMode() bool {
	secret := os.Getenv("DEV_MODE")
	return strings.ToLower(secret) == "true"
}

func getAuthServiceHostname() (string, error) {
	pgUrl := os.Getenv("AUTH_SERVICE_HOSTNAME")
	if pgUrl != "" {
		return pgUrl, nil
	}
	if DevMode {
		return "localhost", nil
	}
	return "", errors.New("AUTH_SERVICE_HOSTNAME env is not set")
}

func getJwtSecret() (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret != "" {
		return secret, nil
	}
	if DevMode {
		return "n13470t9hav8sd-q90h1f9hweonapmcefni", nil
	}
	return "", errors.New("JWT_SECRET env is not set")
}

func EnsureEnvAreSet() error {
	DevMode = getDevMode()
	var err error // to avoid shadowing

	AuthServiceHostname, err = getAuthServiceHostname()
	if err != nil {
		return err
	}

	JwtSecret, err = getJwtSecret()
	if err != nil {
		return err
	}

	return nil
}
