package internal_test

import (
	"bytes"
	"encoding/json"
	"msmc/auth-service/internal"
	"msmc/auth-service/test"
	"net/http/httptest"
	"reflect"
	"testing"
)

var testUserInfoData = test.SeedUsers[0]

func TestUserInfo(t *testing.T) {
	test.SeedDb(t)

	userData := login(t, internal.LoginDto{
		Email:    testUserInfoData.Email,
		Password: testUserInfoData.Password,
	})

	userPublicData := internal.UserPublicData{
		Username: userData.User.Username,
		Email:    userData.User.Email,
		Id:       userData.User.Id,
	}

	checkUserData(t, userPublicData, internal.UserPublicData{Id: userPublicData.Id})
	checkUserData(t, userPublicData, internal.UserPublicData{Email: userPublicData.Email})
	checkUserData(t, userPublicData, internal.UserPublicData{Username: userPublicData.Username})

	test.ClearDb(t)
}

func checkUserData(t *testing.T, expected, partial internal.UserPublicData) {
	if !reflect.DeepEqual(expected, getUserData(t, partial)) {
		t.Fatalf("User data is not equal for %+v", partial)
	}
}

func getUserData(t *testing.T, reqBody internal.UserPublicData) internal.UserPublicData {
	app := internal.GetApp(test.GetDb(t))

	reqBodyRaw, err := json.Marshal(reqBody)
	handleErr(t, err)

	response, err := app.Test(httptest.NewRequest("POST", "/get-user-data", bytes.NewBuffer(reqBodyRaw)))
	handleErr(t, err)

	var userData internal.UserPublicData

	err = json.NewDecoder(response.Body).Decode(&userData)
	handleErr(t, err)

	return userData
}
