package basic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// TODO: Add remaining tests for other endpoints

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type BasicTestSuite struct {
	suite.Suite
}

func (suite *BasicTestSuite) TestGetUser() {
	t := suite.T()
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("userID", "userID1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler := http.HandlerFunc(GetUser)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, rr.Code, http.StatusOK, fmt.Sprintf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK))

	// Check the response body is what we expect.
	expectedUser := User{
		UserID:   "userID1",
		Username: "GoUser",
		Email:    "gouser@gouser.com",
	}

	receivedUser := &User{}

	if err := json.Unmarshal(rr.Body.Bytes(), receivedUser); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expectedUser, *receivedUser, "Received and Expected users should be equal")

}

func (suite *BasicTestSuite) TestDeleteUser() {
	t := suite.T()
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/", nil)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("userID", "userID1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler := http.HandlerFunc(DeleteUser)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, rr.Code, http.StatusOK, fmt.Sprintf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK))

	// TODO: Add database check here to make sure entry was deleted

}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(BasicTestSuite))
}
