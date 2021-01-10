package api

import (
	"bytes"
	"encoding/json"
	"hw21/pkg/storage"
	"hw21/pkg/storage/memstore"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

var s = &Service{}

type Case struct {
	name    string
	method  string
	path    string
	status  int
	payload []storage.Film
	result  []storage.Film
	wantRes bool
	err     string
	wantErr bool
}

func TestMain(m *testing.M) {
	router := mux.NewRouter()
	strg := memstore.New()

	s = &Service{
		Router: router,
		strg:   strg,
	}

	s.endpoints()

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestAPI(t *testing.T) {
	tcases := []Case{
		{
			name:   "Добавление фильма",
			method: http.MethodPost,
			path:   "/film",
			status: http.StatusOK,
			payload: []storage.Film{
				{
					Title:       "Зачарованные",
					ReleaseDate: 1610143502,
					BoxOffice:   12312312.123,
					Rating:      "PG-13",
					StudioID:    1,
				},
			},
			wantRes: false,
			wantErr: false,
		},
		{
			name:    "Добавление фильма по неверному пути",
			method:  http.MethodPost,
			path:    "/film/1",
			status:  http.StatusMethodNotAllowed,
			wantRes: false,
			wantErr: false,
		},
		{
			name:   "Обновление фильма",
			method: http.MethodPut,
			path:   "/film",
			status: http.StatusOK,
			payload: []storage.Film{
				{
					ID:          1,
					Title:       "Зачарованные",
					ReleaseDate: 1610143502,
					BoxOffice:   12312312.123,
					Rating:      "PG-13",
					StudioID:    1,
				},
			},
			wantRes: false,
			wantErr: false,
		},
		{
			name:   "Удаление фильма",
			method: http.MethodDelete,
			path:   "/film",
			status: http.StatusOK,
			payload: []storage.Film{
				{
					ID: 1,
				},
			},
			wantRes: false,
			wantErr: false,
		},
		{
			name:   "Получение фильма",
			method: http.MethodGet,
			path:   "/film/3",
			status: http.StatusOK,
			result: []storage.Film{
				{
					Title: "Старикам здесь не место",
				},
			},
			wantRes: true,
			wantErr: false,
		},
		{
			name:    "Получение фильма с неверным запросом",
			method:  http.MethodGet,
			path:    "/film/3a",
			status:  http.StatusBadRequest,
			wantRes: false,
			err:     "id must be integer",
			wantErr: true,
		},
	}

	for _, tt := range tcases {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			var payload []byte
			rr := httptest.NewRecorder()

			if !reflect.ValueOf(tt.payload).IsZero() {
				switch tt.method {
				case http.MethodPost:
					payload, _ = json.Marshal(tt.payload)
				default:
					payload, _ = json.Marshal(tt.payload[0])
				}

				req = httptest.NewRequest(tt.method, tt.path, bytes.NewBuffer(payload))
			} else {
				req = httptest.NewRequest(tt.method, tt.path, &bytes.Buffer{})
			}

			s.Router.ServeHTTP(rr, req)

			if rr.Code != tt.status {
				log.Fatalf("ошибка: ожидался статус %d, получен %d", tt.status, rr.Code)
			}

			if tt.wantErr {
				body, _ := ioutil.ReadAll(rr.Body)
				if !strings.Contains(string(body), tt.err) {
					log.Fatalf("ошибка: ожидался ответ c ошибкой %s, получен %s", tt.err, string(body))
				}
			}

			if tt.wantRes {
				res := []storage.Film{}
				json.NewDecoder(rr.Body).Decode(&res)

				if tt.result[0].Title != res[0].Title {
					log.Fatalf("ошибка: ожидался ответ %+v, получен %+v", tt.result, res)
				}
			}
		})
	}
}
