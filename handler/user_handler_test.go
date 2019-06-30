package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/app/realworld/model/user"

	"github.com/stretchr/testify/require"

	"github.com/app/realworld/handler/mock"
	"github.com/app/realworld/handler/resource"
	"github.com/app/realworld/router"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	expectedEmail := "example@sample.com"
	expectedPassword := "password"
	expectedUsername := "Joe"
	expectedToken := "test_token"

	// request bodyの用意
	requestBody := new(resource.LoginRequest)
	requestBody.User.Email = expectedEmail
	requestBody.User.Password = expectedPassword

	j, err := json.Marshal(requestBody)
	require.NoError(t, err)

	// mockの用意
	mock := new(mock.AuthMiddle)
	loginInfo := user.LoginInfo{
		Email:    expectedEmail,
		Username: expectedUsername,
		Bio:      "",
	}
	mock.On("Authenticate", expectedEmail, expectedPassword).
		Return(loginInfo, expectedToken, nil)

	// TestServerとRequesterの準備
	r := router.New()
	hand := NewHandler(mock)
	hand.Register(r)

	server := httptest.NewServer(r)
	defer server.Close()

	req := httptest.NewRequest("POST", "/users/login", strings.NewReader(string(j)))
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	var actual resource.LoginResponse
	assert.NoError(t, json.Unmarshal(res.Body.Bytes(), &actual))

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, expectedEmail, actual.User.Email)
	assert.Equal(t, expectedUsername, actual.User.Username)
	assert.Equal(t, expectedToken, actual.User.Token)
}
