package test

import (
	"backend/internal"
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http/httptest"
	"os"
	"testing"
)

func setEnvs(envs map[string]string) error {
	for key, value := range envs {
		err := os.Setenv(key, value)
		if err != nil {
			return err
		}
	}
	return nil
}

func TestAuth(t *testing.T) {
	err := setEnvs(map[string]string{
		"DEV_MODE":              "false",
		"AUTH_SERVICE_HOSTNAME": "example.com", // testing uses example.com as default hostname
		"JWT_SECRET":            "test-secret",
	})
	if err != nil {
		t.Fatalf("Failed to set envs: %v", err)
	}

	app := internal.GetApp()
	reqBody := internal.LoginRemoteDto{UserId: 1}

	tokenString := getJwt(t, app, reqBody)
	payloadFromServer := getJwtPayload(t, app, tokenString)

	if payloadFromServer.UserID != reqBody.UserId {
		t.Fatalf("Expected user ID %d, got %d", reqBody.UserId, payloadFromServer.UserID)
	}
}

func getJwtPayload(t *testing.T, app *fiber.App, token string) internal.JwtPayload {
	reqDto := internal.GetJwtPayloadDto{Token: token}
	reqBody, err := json.Marshal(reqDto)
	if err != nil {
		t.Fatalf("Failed to encode request body: %v", err)
	}

	testReq := httptest.NewRequest("POST", "/get-jwt-payload", bytes.NewBuffer(reqBody))
	testReq.Header.Set("Content-Type", "application/json")

	res, err := app.Test(testReq)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}

	if res.StatusCode/100 != 2 {
		t.Fatalf("Expected status code 2xx, got %d", res.StatusCode)
	}

	body, err := JSONToMap(res.Body)
	_ = res.Body.Close()
	if err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	if body["error"] != nil {
		t.Fatalf("Error in response: %s", body["error"])
	}
	payload := body["payload"].(map[string]interface{})

	return internal.JwtPayload{
		Exp:    int64(payload["ExpiresAt"].(float64)),
		UserID: int(payload["user_id"].(float64)),
	}
}

func getJwt(t *testing.T, app *fiber.App, reqBody internal.LoginRemoteDto) string {
	reqBodyString, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	testReq := httptest.NewRequest("POST", "/login-remote", bytes.NewBuffer(reqBodyString))
	testReq.Header.Set("Content-Type", "application/json")

	res, err := app.Test(testReq)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}

	if res.StatusCode/100 != 2 {
		t.Logf("Response body: ")
		_ = printReadCloser(t, res.Body)
		t.Fatalf("Expected status code 2xx, got %d", res.StatusCode)
	}

	body, err := JSONToMap(res.Body)
	_ = res.Body.Close()
	if err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	token, ok := body["token"].(string)
	if !ok {
		t.Fatalf("Failed to get token from response body")
	}

	if token == "" {
		t.Fatalf("Token is empty")
	}

	return token
}

func JSONToMap(reader io.ReadCloser) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.NewDecoder(reader).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func printReadCloser(t *testing.T, reader io.ReadCloser) error {
	data, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	t.Log(string(data))
	return nil
}
