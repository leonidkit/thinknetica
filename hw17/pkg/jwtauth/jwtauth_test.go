package jwtauth

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type Case struct {
	Method   string
	Code     int
	Request  interface{}
	Response string
	Path     string
}

func TestAuth(t *testing.T) {
	api := &AuthServer{}
	api.Endpoints()

	tcases := []Case{
		Case{
			Method:   "GET",
			Code:     http.StatusMethodNotAllowed,
			Request:  nil,
			Response: "",
			Path:     "/auth",
		},
		Case{
			Method:   "POST",
			Code:     http.StatusInternalServerError,
			Request:  "{error : json}",
			Response: "unmarshaling input json error",
			Path:     "/auth",
		},
		Case{
			Method: "POST",
			Code:   http.StatusBadRequest,
			Request: map[string]string{
				"login":    "leonid_kit@mail.ru",
				"password": "fskjdf",
			},
			Response: "incorrect username or password",
			Path:     "/auth",
		},
		Case{
			Method: "POST",
			Code:   http.StatusNotFound,
			Request: map[string]string{
				"login":    "leonid_kit1@mail.ru",
				"password": "qwerty",
			},
			Response: "user not found",
			Path:     "/auth",
		},
		Case{
			Method: "POST",
			Code:   http.StatusOK,
			Request: map[string]string{
				"login":    "leonid_kit@mail.ru",
				"password": "qwerty",
			},
			Response: "",
			Path:     "/auth",
		},
		Case{
			Method: "POST",
			Code:   http.StatusBadRequest,
			Request: map[string]string{
				"login":    "leonid_kit@mail.ru",
				"password": "123",
			},
			Response: "incorrect username or password",
			Path:     "/auth",
		},
	}

	for idx, tt := range tcases {
		reqstr, err := json.Marshal(tt.Request)
		if err != nil {
			log.Fatal(err.Error())
		}

		req := httptest.NewRequest(tt.Method, "/auth", bytes.NewBuffer(reqstr))
		rr := httptest.NewRecorder()

		api.Handler.ServeHTTP(rr, req)
		if rr.Code != tt.Code {
			log.Fatalf("[%d] expected http status %v, got %v", idx, tt.Code, rr.Code)
		}

		// пропуск для кейса, когда не ожидаем что-либо от сервера в теле ответа
		if tt.Response == "" {
			continue
		}

		var respJSON interface{}
		err = json.NewDecoder(rr.Body).Decode(&respJSON)
		if err != nil {
			if err != io.EOF {
				log.Fatalf("[%d] json response was expected, got unmarshaling error: %v", idx, err.Error())
			}
		}

		// тестирование кейса с ошибкой
		if rr.Code != 200 {
			respMap, ok := respJSON.(map[string]interface{})
			if !ok {
				log.Fatalf("[%d] map[string]string response was expected, got %v", idx, respJSON)
			}

			if !strings.Contains(respMap["error"].(string), tt.Response) {
				log.Fatalf("[%d] expected respons contains %v, got %v", idx, tt.Response, respMap["error"])
			}
		}
	}
}
