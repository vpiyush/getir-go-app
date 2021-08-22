package apis

import (
	"bytes"
	"fmt"
	memdb "github.com/vpiyush/getir-go-app/inmemdb"
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

	fmt.Println(rr.Body.String())
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
			status, http.StatusOK)
	}

	fmt.Println(rr.Body.String())
}

func TestGetPair_Success(t *testing.T) {
	memdb.Cache.Insert("active-tabs", "getir")
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

func TestGetPair_KeyNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/pair?key=active", nil)
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
