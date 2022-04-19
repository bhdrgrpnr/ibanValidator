package api

import (
	"IbanValidator/internal/service"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestValidIban(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/validateIban/DE89370400440532013000", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler := NewHandler()
	rr := httptest.NewRecorder()
	service.InitService()
	handler.ServeHTTP(rr, req)
	if rr.Result().StatusCode != http.StatusOK {
		t.Errorf("unexpected status %d", rr.Result().StatusCode)
		return
	}

	var response Response
	json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Equal(t, response.Result, true)
	assert.Equal(t, response.ErrorExp, "")
}

func TestInvalidIban(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/validateIban/DE893704004405320130001", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler := NewHandler()
	rr := httptest.NewRecorder()
	service.InitService()
	handler.ServeHTTP(rr, req)
	if rr.Result().StatusCode != http.StatusBadRequest {
		t.Errorf("unexpected status %d", rr.Result().StatusCode)
		return
	}

	var response Response
	json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Equal(t, response.Result, false)
	assert.True(t, strings.Contains(response.ErrorExp, "Iban should be with length"))
}

func TestInvalidIbanWrongCountry(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/validateIban/SE893704004405320130001", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler := NewHandler()
	rr := httptest.NewRecorder()
	service.InitService()
	handler.ServeHTTP(rr, req)
	if rr.Result().StatusCode != http.StatusBadRequest {
		t.Errorf("unexpected status %d", rr.Result().StatusCode)
		return
	}

	var response Response
	json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Equal(t, response.Result, false)
	assert.True(t, strings.Contains(response.ErrorExp, "Country not supported"))
}
