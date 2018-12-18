package basic

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
)

func TestGetUser(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("userID", "userID1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler := http.HandlerFunc(GetUser)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expectedUser := User{
		UserID:   "userID1",
		Username: "GoUser",
		Email:    "gouser@gouser.com",
	}

	expectedByteArray, err := json.Marshal(expectedUser)
	if err != nil {
		t.Fatal(err)
	}

	expected := string(expectedByteArray)

	receivedUser := &User{}

	if err := json.Unmarshal(rr.Body.Bytes(), receivedUser); err != nil {
		t.Fatal(err)
	}

	if expectedUser != *receivedUser {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
