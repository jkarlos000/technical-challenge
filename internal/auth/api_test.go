package auth

import (
	"context"
	"github.com/jkarlos000/technical-challenge/internal/errors"
	"github.com/jkarlos000/technical-challenge/internal/test"
	"github.com/jkarlos000/technical-challenge/pkg/log"
	"net/http"
	"testing"
)

type mockService struct{}

func (m mockService) Login(ctx context.Context, username, password string) (string, error) {
	if username == "test" && password == "pass" {
		return "token-100", nil
	}
	return "", errors.Unauthorized("")
}

func TestAPI(t *testing.T) {
	logger, _ := log.NewForTest()
	router := test.MockRouter(logger)
	RegisterHandlers(router.Group(""), mockService{}, logger)

	tests := []test.APITestCase{
		{"success", "POST", "/login", `{"username":"test","password":"pass"}`, nil, http.StatusOK, `{"token":"token-100"}`},
		{"bad credential", "POST", "/login", `{"username":"test","password":"wrong pass"}`, nil, http.StatusUnauthorized, ""},
		{"bad json", "POST", "/login", `"username":"test","password":"wrong pass"}`, nil, http.StatusBadRequest, ""},
	}
	for _, tc := range tests {
		test.Endpoint(t, router, tc)
	}
}
