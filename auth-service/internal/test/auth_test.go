package internal_test

import (
	"bytes"
	"encoding/json"
	"msmc/auth-service/internal"
	"msmc/auth-service/test"
	"net/http/httptest"
	"testing"
)

func TestAuth(t *testing.T) {
	test.SeedDb(t)

	register(t, testAuthUser)

	loginDtoWithoutEmail := RegisterDtoToLoginDto(testAuthUser)
	loginDtoWithoutEmail.Email = ""
	login(t, loginDtoWithoutEmail)

	loginDtoWithoutUsername := RegisterDtoToLoginDto(testAuthUser)
	loginDtoWithoutUsername.Username = ""
	login(t, loginDtoWithoutUsername)

	// Test login both with email and username
	login(t, RegisterDtoToLoginDto(testAuthUser))

	test.ClearDb(t)
}

var testAuthUser = internal.RegisterDto{
	Email:    "auth_test@example.com",
	Username: "AuthTest",
	Password: "12345678",
}

func RegisterDtoToLoginDto(dto internal.RegisterDto) internal.LoginDto {
	return internal.LoginDto{
		Email:    dto.Email,
		Username: dto.Username,
		Password: dto.Password,
	}
}

func login(t *testing.T, user internal.LoginDto) internal.AuthResponse {
	app := internal.GetApp(test.GetDb(t))

	reqBodyRaw, err := json.Marshal(user)
	handleErr(t, err)

	response, err := app.Test(httptest.NewRequest("POST", "/login", bytes.NewBuffer(reqBodyRaw)))
	handleErr(t, err)

	var tokenRes internal.AuthResponse

	err = json.NewDecoder(response.Body).Decode(&tokenRes)
	handleErr(t, err)

	return tokenRes
}

func register(t *testing.T, user internal.RegisterDto) internal.AuthResponse {
	app := internal.GetApp(test.GetDb(t))

	reqBodyRaw, err := json.Marshal(user)
	handleErr(t, err)

	response, err := app.Test(httptest.NewRequest("POST", "/register", bytes.NewBuffer(reqBodyRaw)))
	handleErr(t, err)

	var tokenRes internal.AuthResponse

	err = json.NewDecoder(response.Body).Decode(&tokenRes)
	handleErr(t, err)

	return tokenRes
}

func handleErr(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("%v", err)
	}
}
