package api

import (
	"bufio"
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"gosearch/pkg/crawler"
	"gosearch/pkg/engine"
	"gosearch/pkg/index/fakeindex"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var (
	s    = &Service{}
	data = []crawler.Document{
		crawler.Document{
			ID:    uint64(1),
			Title: "Как использовать git?",
			URL:   "http://localhost1",
		},
		crawler.Document{
			ID:    uint64(2),
			Title: "Прикладное применение подорожника как лекарство",
			URL:   "http://localhost2",
		},
		crawler.Document{
			ID:    uint64(3),
			Title: "Криптовалюта как будущее финансовой системы?",
			URL:   "http://localhost3",
		},
	}
)

type Case struct {
	Method  string
	Path    string
	Status  int
	Result  []crawler.Document
	Payload map[string]interface{}
	wantErr bool
	Error   string
}

func TestMain(m *testing.M) {
	fi := fakeindex.New(data)
	s = New(mux.NewRouter(), engine.New(fi, data))
	s.endpoints()

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestSearch(t *testing.T) {
	tests := []Case{
		Case{
			Method:  http.MethodGet,
			Path:    "/api/search/git",
			Status:  http.StatusOK,
			wantErr: false,
			Result: []crawler.Document{
				crawler.Document{
					ID:    uint64(1),
					Title: "Как использовать git?",
					URL:   "http://localhost1",
				},
			},
		},
		Case{
			Method:  http.MethodGet,
			Path:    "/api/search/go",
			Status:  http.StatusInternalServerError,
			wantErr: true,
			Error:   "document not found\n",
		},
	}

	for idx, tt := range tests {
		var req *http.Request
		if tt.Method == http.MethodPost {
			payload, _ := json.Marshal(tt.Payload)
			req = httptest.NewRequest(tt.Method, tt.Path, bytes.NewBuffer(payload))
		}
		if tt.Method == http.MethodGet {
			req = httptest.NewRequest(tt.Method, tt.Path, &bufio.Reader{})
		}
		rr := httptest.NewRecorder()

		s.Router.ServeHTTP(rr, req)

		if rr.Code != tt.Status {
			t.Fatalf("[%d] >>> код неверен: получили %d, а хотели %d", idx, rr.Code, tt.Status)
		}

		if tt.wantErr {
			resp, err := ioutil.ReadAll(rr.Body)
			if err != nil {
				t.Fatal(err.Error())
			}
			if string(resp) != tt.Error {
				t.Fatalf("[%d] >>> ответ неверен: получили %s, а хотели %s", idx, resp, tt.Error)
			}
		}

		if !tt.wantErr {
			var resp []crawler.Document
			err := json.NewDecoder(rr.Body).Decode(&resp)
			if err != nil {
				t.Fatalf("[%d] >>> ответ неверен: не удается преобразовать ответ в структуру документа: %s", idx, err.Error())
			}

			if tt.Result[0] != resp[0] {
				t.Fatalf("[%d] >>> ответ неверен: получили %+v, а хотели %+v", idx, resp[0], tt.Result[0])
			}
		}
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
	tests := []Case{
		Case{
			Method:  http.MethodPost,
			Path:    "/api/create_doc",
			Status:  http.StatusOK,
			wantErr: false,
			Payload: map[string]interface{}{
				"url":   "http://localhost",
				"title": "New record",
			},
		},
		Case{
			Method:  http.MethodPost,
			Path:    "/api/create_doc",
			Status:  http.StatusInternalServerError,
			wantErr: true,
			Error:   "cannot unmarshal",
			Payload: map[string]interface{}{
				"url":   1,
				"title": "New record",
			},
		},
	}

	for idx, tt := range tests {
		var req *http.Request
		if tt.Method == http.MethodPost {
			payload, _ := json.Marshal(tt.Payload)
			req = httptest.NewRequest(tt.Method, tt.Path, bytes.NewBuffer(payload))
		}
		if tt.Method == http.MethodGet {
			req = httptest.NewRequest(tt.Method, tt.Path, &bufio.Reader{})
		}
		rr := httptest.NewRecorder()

		s.Router.ServeHTTP(rr, req)

		if rr.Code != tt.Status {
			t.Fatalf("[%d] >>> код неверен: получили %d, а хотели %d", idx, rr.Code, tt.Status)
		}

		if tt.wantErr {
			resp, err := ioutil.ReadAll(rr.Body)
			if err != nil {
				t.Fatal(err.Error())
			}
			if !strings.Contains(string(resp), tt.Error) {
				t.Fatalf("[%d] >>> ответ неверен: получили %s, а хотели %s", idx, resp, tt.Error)
			}
		}

		if !tt.wantErr {
			_, err := s.engine.Search(tt.Payload["title"].(string))
			if err != nil {
				t.Fatalf("[%d] >>> результат неверен: не удается найти добавленный документ: %s", idx, err.Error())
			}
		}
	}
}

func TestDelete(t *testing.T) {
	tests := []Case{
		Case{
			Method:  http.MethodPost,
			Path:    "/api/delete_doc",
			Status:  http.StatusOK,
			wantErr: false,
			Result: []crawler.Document{
				crawler.Document{
					ID:    uint64(1),
					Title: "Как использовать git?",
					URL:   "http://localhost1",
				},
			},
			Payload: map[string]interface{}{
				"id": 1,
			},
		},
		Case{
			Method:  http.MethodPost,
			Path:    "/api/delete_doc",
			Status:  http.StatusInternalServerError,
			wantErr: true,
			Error:   "cannot unmarshal",
			Payload: map[string]interface{}{
				"id": "1",
			},
		},
	}

	for idx, tt := range tests {
		var req *http.Request
		if tt.Method == http.MethodPost {
			payload, _ := json.Marshal(tt.Payload)
			req = httptest.NewRequest(tt.Method, tt.Path, bytes.NewBuffer(payload))
		}
		if tt.Method == http.MethodGet {
			req = httptest.NewRequest(tt.Method, tt.Path, &bufio.Reader{})
		}
		rr := httptest.NewRecorder()

		s.Router.ServeHTTP(rr, req)

		if rr.Code != tt.Status {
			t.Fatalf("[%d] >>> код неверен: получили %d, а хотели %d", idx, rr.Code, tt.Status)
		}

		if tt.wantErr {
			resp, err := ioutil.ReadAll(rr.Body)
			if err != nil {
				t.Fatal(err.Error())
			}
			if !strings.Contains(string(resp), tt.Error) {
				t.Fatalf("[%d] >>> ответ неверен: получили %s, а хотели %s", idx, resp, tt.Error)
			}
		}

		if !tt.wantErr {
			_, err := s.engine.Search(tt.Result[0].Title)
			if err == nil {
				t.Fatalf("[%d] >>> результат неверен: удается найти удаленный документ", idx)
			}
		}
	}
}

func TestUpdate(t *testing.T) {
	tests := []Case{
		Case{
			Method:  http.MethodPost,
			Path:    "/api/update_doc",
			Status:  http.StatusOK,
			wantErr: false,
			Payload: map[string]interface{}{
				"id":    3,
				"url":   "http://localhost3",
				"title": "Золото как будущее финансовой системы?",
			},
		},
		Case{
			Method:  http.MethodPost,
			Path:    "/api/update_doc",
			Status:  http.StatusInternalServerError,
			wantErr: true,
			Error:   "cannot unmarshal",
			Payload: map[string]interface{}{
				"id":    "1",
				"url":   "http://localhost",
				"title": "New record",
			},
		},
	}

	for idx, tt := range tests {
		var req *http.Request
		if tt.Method == http.MethodPost {
			payload, _ := json.Marshal(tt.Payload)
			req = httptest.NewRequest(tt.Method, tt.Path, bytes.NewBuffer(payload))
		}
		if tt.Method == http.MethodGet {
			req = httptest.NewRequest(tt.Method, tt.Path, &bufio.Reader{})
		}
		rr := httptest.NewRecorder()

		s.Router.ServeHTTP(rr, req)

		if rr.Code != tt.Status {
			res, _ := ioutil.ReadAll(rr.Body)
			t.Log(string(res))
			t.Fatalf("[%d] >>> код неверен: получили %d, а хотели %d", idx, rr.Code, tt.Status)
		}

		if tt.wantErr {
			resp, err := ioutil.ReadAll(rr.Body)
			if err != nil {
				t.Fatal(err.Error())
			}
			if !strings.Contains(string(resp), tt.Error) {
				t.Fatalf("[%d] >>> ответ неверен: получили %s, а хотели %s", idx, resp, tt.Error)
			}
		}

		if !tt.wantErr {
			_, err := s.engine.Search(tt.Payload["title"].(string))
			if err != nil {
				t.Fatalf("[%d] >>> результат неверен: не удается найти обновленный документ: %s", idx, err.Error())
			}
		}
	}
}
