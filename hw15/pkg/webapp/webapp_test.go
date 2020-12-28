package webapp

import (
	"gosearch/pkg/crawler"
	"gosearch/pkg/engine"
	"gosearch/pkg/index/fakeindex"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

var (
	client = &http.Client{Timeout: time.Second}
	indx   = fakeindex.New()
	data   = []crawler.Document{
		crawler.Document{
			ID:    uint64(1),
			Title: "Как использовать git?",
			URL:   "http://localhost",
		},
		crawler.Document{
			ID:    uint64(2),
			Title: "Прикладное применение подорожника как лекарство",
			URL:   "http://localhost",
		},
		crawler.Document{
			ID:    uint64(3),
			Title: "Криптовалюта как будущее финансовой системы?",
			URL:   "http://localhost",
		},
	}
	engn = engine.New(indx, data)
	host = "0.0.0.0"
	port = "8000"
)

type Case struct {
	Method string
	Path   string
	Status int
	Result string
}

func TestIndex(t *testing.T) {
	wa := New(*engn)
	ts := httptest.NewServer(wa.handlers())
	defer ts.Close()

	tests := []Case{
		{
			Method: "GET",
			Path:   "/index",
			Status: http.StatusOK,
			Result: "",
		},
		{
			Method: "POST",
			Path:   "/index",
			Status: http.StatusMethodNotAllowed,
			Result: "",
		},
		{
			Method: "GET",
			Path:   "/indexx",
			Status: http.StatusNotFound,
			Result: "page not found",
		},
	}

	for idx, tt := range tests {
		req, err := http.NewRequest(tt.Method, ts.URL+tt.Path, nil)
		if err != nil {
			t.Errorf("[%d] request generating error: %v", idx, err)
			continue
		}

		resp, err := client.Do(req)
		if err != nil {
			t.Errorf("[%d] request error: %v", idx, err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != tt.Status {
			t.Errorf("[%d] expected http status %v, got %v", idx, tt.Status, resp.StatusCode)
			continue
		}

		if tt.Status != http.StatusOK {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("[%d] reading body error: %v", idx, err)
				continue
			}
			sbody := string(body)

			if !strings.Contains(sbody, tt.Result) {
				t.Errorf("[%d] results not match\nGot: %#v\nExpected: %#v", idx, sbody, tt.Result)
				continue
			}
		}
	}
}

func TestDocs(t *testing.T) {
	wa := New(*engn)
	ts := httptest.NewServer(wa.handlers())
	defer ts.Close()

	tests := []Case{
		{
			Method: "GET",
			Path:   "/index",
			Status: http.StatusOK,
			Result: "",
		},
		{
			Method: "POST",
			Path:   "/docs",
			Status: http.StatusMethodNotAllowed,
			Result: "",
		},
		{
			Method: "GET",
			Path:   "/docss",
			Status: http.StatusNotFound,
			Result: "page not found",
		},
	}

	for idx, tt := range tests {
		req, err := http.NewRequest(tt.Method, ts.URL+tt.Path, nil)
		if err != nil {
			t.Errorf("[%d] request generating error: %v", idx, err)
			continue
		}

		resp, err := client.Do(req)
		if err != nil {
			t.Errorf("[%d] request error: %v", idx, err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != tt.Status {
			t.Errorf("[%d] expected http status %v, got %v", idx, tt.Status, resp.StatusCode)
			continue
		}

		if tt.Status != http.StatusOK {
			// пропуск кейса в ответе на который не приходит тело
			if tt.Result == "" {
				continue
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("[%d] reading body error: %v", idx, err)
				continue
			}
			sbody := string(body)

			if !strings.Contains(sbody, tt.Result) {
				t.Errorf("[%d] results not match\nGot: %#v\nExpected: %#v", idx, sbody, tt.Result)
				continue
			}
		}
	}
}
