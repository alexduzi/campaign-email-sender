package endpoints

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	validMockToken string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMiwiZW1haWwiOiJhZG1pbkBhZG1pbi5jb20ifQ.y4Xm7i0FAVnt29hlSJVfAeumUArjUMID1jeRyqb-IW8"
	validMockEmail string = "admin@admin.com"
)

func Test_Auth_Unauthorized(t *testing.T) {
	assert := assert.New(t)

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("next handler should not be called")
	})

	handlerFunc := Auth(nextHandler)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusUnauthorized, res.Code)
	assert.Equal(res.Body.String(), "{\"error\":\"request does not contain an authorization header\"}\n")
}

func Test_Auth_WhenAuthorizationIsInvalid_ReturnError(t *testing.T) {
	assert := assert.New(t)

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("next handler should not be called")
	})

	ValidateToken = func(token string, ctx context.Context) (string, error) {
		return "", errors.New("invalid token")
	}
	handlerFunc := Auth(nextHandler)
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Authorization", "Bearer invalid_token")
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusUnauthorized, res.Code)
	assert.Equal(res.Body.String(), "{\"error\":\"invalid token\"}\n")
}

func Test_Auth_WhenAuthorizationIsValid_CallNextHandler(t *testing.T) {
	assert := assert.New(t)
	var email string
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email = r.Context().Value("email").(string)
	})

	ValidateToken = func(token string, ctx context.Context) (string, error) {
		return "admin@admin.com", nil
	}
	handlerFunc := Auth(nextHandler)
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Authorization", "Bearer "+validMockToken)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	assert.Equal(validMockEmail, email)
}
