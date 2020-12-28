package api

import (
	"bufio"
	"bytes"
	"encoding/json"
	"gosearch/pkg/engine"
	"gosearch/pkg/index/fakeindex"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

var s = &Service{}

func TestMain(m *testing.M) {
	fi := fakeindex.New()
	data, _ := fi.Find("")
	s = New(mux.NewRouter(), *engine.New(fi, data))
	s.endpoints()

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestSearch(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/search/go", &bufio.Reader{})

	rr := httptest.NewRecorder()

	s.Router.ServeHTTP(rr, req)

	if rr.Code != 200 {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
}

func TestDocs(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/docs", &bufio.Reader{})

	rr := httptest.NewRecorder()

	s.Router.ServeHTTP(rr, req)

	if rr.Code != 200 {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
}

func TestCreate(t *testing.T) {
	q := map[string]string{
		"url":   "http://localhost",
		"title": "New record",
	}

	payload, _ := json.Marshal(q)
	req := httptest.NewRequest("POST", "/api/create", bytes.NewBuffer(payload))

	rr := httptest.NewRecorder()

	s.Router.ServeHTTP(rr, req)

	if rr.Code != 200 {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
}

func TestDelete(t *testing.T) {
	q := map[string]string{
		"id": "1234",
	}

	payload, _ := json.Marshal(q)
	req := httptest.NewRequest("POST", "/api/delete", bytes.NewBuffer(payload))

	rr := httptest.NewRecorder()

	s.Router.ServeHTTP(rr, req)

	if rr.Code != 200 {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
}

func TestUpdate(t *testing.T) {
	q := map[string]string{
		"id":    "1234",
		"url":   "http://localhost",
		"title": "New record",
	}

	payload, _ := json.Marshal(q)
	req := httptest.NewRequest("POST", "/api/update", bytes.NewBuffer(payload))

	rr := httptest.NewRecorder()

	s.Router.ServeHTTP(rr, req)

	if rr.Code != 200 {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
}
