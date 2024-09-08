package internal_test

import (
	"bytes"
	"encoding/json"
	"msmc/auth-service/internal"
	"msmc/auth-service/test"
	"net/http/httptest"
	"testing"
)

func TestLoginRemote(t *testing.T) {
	test.SeedDb(t)

	userData := login(t, internal.LoginDto{
		Email:    testUserInfoData.Email,
		Password: testUserInfoData.Password,
	})

	loginRemote(t, internal.LoginRemoteDto{
		Token:     userData.Token,
		RemoteUrl: "http://localhost:3001",
	})

	test.ClearDb(t)
}

func loginRemote(t *testing.T, reqBody internal.LoginRemoteDto) internal.TokenResponse {
	app := internal.GetApp(test.GetDb(t))

	reqBodyRaw, err := json.Marshal(reqBody)
	handleErr(t, err)

	response, err := app.Test(httptest.NewRequest("POST", "/login-remote", bytes.NewBuffer(reqBodyRaw)))
	handleErr(t, err)

	var tokenResponse internal.TokenResponse

	err = json.NewDecoder(response.Body).Decode(&tokenResponse)
	handleErr(t, err)

	return tokenResponse
}
