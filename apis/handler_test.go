package apis

import (
	"bytes"
	"fmt"
	"github.com/vpiyush/getir-go-app/daos"
	"github.com/vpiyush/getir-go-app/services"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetRecords_Success(t *testing.T) {
	var jsonStr = []byte(`{"startDate": "2016-01-26","endDate": "2018-02-02","minCount": 2700,"maxCount": 3000 }`)
	req, err := http.NewRequest("POST", "/records", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Executing Test Get Records")
	//env := Env{records: &mockRecordModel{}}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetRecords)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetRecords_InvalidRequest(t *testing.T) {
	var jsonStr = []byte(`{"startDate": "2016-01-26","endDate": "2018-02-02","minCount": 2700,"maxCount": "hello" }`)
	req, err := http.NewRequest("POST", "/records", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Executing Test Get Records")
	//env := Env{records: &mockRecordModel{}}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetRecords)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestGetPair_KeyNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/pair?key=active", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandlePair)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestInsertPair_Success(t *testing.T) {
	var jsonStr = []byte(`{"key":"active-tabs","value":"getir"}`)
	req, err := http.NewRequest("POST", "/pair", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandlePair)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"key":"active-tabs","value":"getir"}`
	if strings.TrimSuffix(rr.Body.String(), "\n") != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestInsertPair_KeyAlreadExists(t *testing.T) {
	var jsonStr = []byte(`{"key":"active-tabs","value":"getir"}`)
	req, err := http.NewRequest("POST", "/pair", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandlePair)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusForbidden {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusForbidden)
	}
}

func TestGetPair_Success(t *testing.T) {
	s := services.NewPairService(daos.NewPairDAO())
	s.Insert("active-tabs", "getir")
	req, err := http.NewRequest("GET", "/pair?key=active-tabs", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandlePair)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"key":"active-tabs","value":"getir"}`
	if strings.TrimSuffix(rr.Body.String(), "\n") != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestInsertPair_InvalidRequest(t *testing.T) {
	var jsonStr = []byte(`{"key":1234,"value":"getir"}`)
	req, err := http.NewRequest("POST", "/pair", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandlePair)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
