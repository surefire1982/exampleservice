package basic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/surefire1982/exampleservice/pkg/entity"
	"github.com/surefire1982/exampleservice/pkg/user"

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
	userHandler *UserHandler
	userID1     string
	userID2     string
}

// setup the userHandler
func (suite *BasicTestSuite) SetupTest() {
	repo := user.NewInMemRepository()
	userSvc := user.NewService(repo)
	suite.userHandler = NewUserHandler(*userSvc)

	// setup the test users
	user1 := &entity.User{
		Username: "username1",
		Email:    "username1@email.com",
	}
	suite.userID1, _ = suite.userHandler.userSvc.Store(user1)

	user2 := &entity.User{
		Username: "username2",
		Email:    "username2@email.com",
	}
	suite.userID2, _ = suite.userHandler.userSvc.Store(user2)
}

func (suite *BasicTestSuite) TestCreateUser() {
	t := suite.T()
	rr := httptest.NewRecorder()

	// setup user
	newUser := entity.User{
		Username: "newUsername",
		Email:    "newUsername@email.com",
	}
	reqBody := new(bytes.Buffer)
	json.NewEncoder(reqBody).Encode(newUser)

	req := httptest.NewRequest("POST", "/", reqBody)

	rctx := chi.NewRouteContext()

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler := http.HandlerFunc(suite.userHandler.Create)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, fmt.Sprintf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK))

	var resp entity.UserResponse

	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, resp)
	assert.NotNil(t, resp.UserID)

	// check created
	foundUser, err := suite.userHandler.userSvc.Find(resp.UserID)
	assert.Nil(t, err)
	assert.NotNil(t, foundUser)
	assert.Equal(t, resp.UserID, foundUser.UserID)
}

func (suite *BasicTestSuite) TestFind() {
	t := suite.T()
	rr := httptest.NewRecorder()

	req := httptest.NewRequest("GET", "/", nil)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("userID", suite.userID1)

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler := http.HandlerFunc(suite.userHandler.Find)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusOK, rr.Code, fmt.Sprintf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK))

	var receivedUser entity.User

	if err := json.Unmarshal(rr.Body.Bytes(), &receivedUser); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, suite.userID1, receivedUser.UserID, "Received and Expected users should be equal")

}

func (suite *BasicTestSuite) TestFindInvalidID() {
	t := suite.T()
	rr := httptest.NewRecorder()

	req := httptest.NewRequest("GET", "/", nil)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("userID", "nonexistentuserID")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler := http.HandlerFunc(suite.userHandler.Find)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusNotFound, rr.Code, fmt.Sprintf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusNotFound))
}

func (suite *BasicTestSuite) TestDeleteUser() {
	t := suite.T()
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/", nil)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("userID", suite.userID1)

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler := http.HandlerFunc(suite.userHandler.Delete)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusOK, rr.Code, fmt.Sprintf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK))

	var resp entity.UserResponse

	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, resp)
	assert.NotNil(t, resp.UserID)

	// search for deleted user
	user, err := suite.userHandler.userSvc.Find(suite.userID1)

	assert.Nil(t, user)
	assert.Equal(t, entity.ErrNotFound.Error(), err.Error())

}

func (suite *BasicTestSuite) TestFindAll() {
	t := suite.T()
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	rctx := chi.NewRouteContext()

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler := http.HandlerFunc(suite.userHandler.FindAll)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusOK, rr.Code, fmt.Sprintf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK))

	var receivedUsers []entity.User

	if err := json.Unmarshal(rr.Body.Bytes(), &receivedUsers); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(receivedUsers))
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(BasicTestSuite))
}
