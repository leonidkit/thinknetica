package webapp

import (
	"encoding/json"
	"gosearch/pkg/crawler"
	"gosearch/pkg/engine"
	"gosearch/pkg/index/fakeindex"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
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
)

type Case struct {
	Method string
	Path   string
	Status int
	Result map[string]interface{}
}

func TestIndex(t *testing.T) {
	wa := New(*engn)
	ts := httptest.NewServer(wa)
	defer ts.Close()

	tests := []Case{
		{
			Method: "GET",
			Path:   "/index",
			Status: http.StatusOK,
			Result: nil,
		},
		{
			Method: "POST",
			Path:   "/index",
			Status: http.StatusMethodNotAllowed,
			Result: map[string]interface{}{
				"error": "method not allowed",
			},
		},
		{
			Method: "GET",
			Path:   "/indexx",
			Status: http.StatusNotFound,
			Result: map[string]interface{}{
				"error": "unknown endpoint",
			},
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

			var result interface{}
			err = json.Unmarshal(body, &result)
			if err != nil {
				t.Errorf("[%d] unmarshaling response error: %v", idx, err)
				continue
			}

			if !reflect.DeepEqual(result, tt.Result) {
				t.Errorf("[%d] results not match\nGot: %#v\nExpected: %#v", idx, result, tt.Result)
				continue
			}
		}
	}
}

func TestDocs(t *testing.T) {
	wa := New(*engn)
	ts := httptest.NewServer(wa)
	defer ts.Close()

	tests := []Case{
		{
			Method: "GET",
			Path:   "/index",
			Status: http.StatusOK,
			Result: nil,
		},
		{
			Method: "POST",
			Path:   "/docs",
			Status: http.StatusMethodNotAllowed,
			Result: map[string]interface{}{
				"error": "method not allowed",
			},
		},
		{
			Method: "GET",
			Path:   "/docss",
			Status: http.StatusNotFound,
			Result: map[string]interface{}{
				"error": "unknown endpoint",
			},
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

			var result interface{}
			err = json.Unmarshal(body, &result)
			if err != nil {
				t.Errorf("[%d] unmarshaling response error: %v", idx, err)
				continue
			}

			if !reflect.DeepEqual(result, tt.Result) {
				t.Errorf("[%d] results not match\nGot: %#v\nExpected: %#v", idx, result, tt.Result)
				continue
			}
		}
	}
}
