package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_generateSalt(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPost, "", nil)
	if err != nil {
		t.Error(err)
	}
	generateSalt(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("expected 200, instead %d", recorder.Code)
	}
	res := map[string]string{}
	if err := json.NewDecoder(recorder.Body).Decode(&res); err != nil {
		t.Error(err)
	}
	if _, ok := res["salt"]; !ok {
		t.Errorf("response body doesn't contain 'salt'")
	}
}
