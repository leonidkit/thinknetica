package jwtauth

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type authInfo struct {
	Login string `json:"login"`
	Pass  string `json:"password"`
}

type User struct {
	Login    string
	Password string
	Role     string
}

var (
	db = []User{
		User{
			Login:    "leonid_kit@mail.ru",
			Password: "d8578edf8458ce06fbc5bb76a58c5ca4",
			Role:     "admin",
		},
		User{
			Login:    "andrcop@gmail.ru",
			Password: "d8578edf8458ce06fbc5bb76a58c5ca4",
			Role:     "user",
		},
	}
	secretKey = "yd0jWLynmm"
)

type AuthServer struct {
	Handler http.Handler
}

func JSONError(w http.ResponseWriter, resp interface{}, code int) {
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

func (as *AuthServer) Auth(w http.ResponseWriter, r *http.Request) {
	h := md5.New()
	ai := authInfo{}
	err := json.NewDecoder(r.Body).Decode(&ai)
	if err != nil {
		resp := map[string]string{
			"error": "unmarshaling input json error: " + err.Error(),
		}

		JSONError(w, resp, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	for _, u := range db {
		_, err := h.Write([]byte(ai.Pass))
		if err != nil {
			resp := map[string]string{
				"error": err.Error(),
			}

			JSONError(w, resp, http.StatusInternalServerError)
			return
		}
		pass := hex.EncodeToString(h.Sum(nil))

		if u.Login == ai.Login {
			if u.Password != pass {
				resp := map[string]string{
					"error": "incorrect username or password",
				}

				JSONError(w, resp, http.StatusBadRequest)
				return
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"role": u.Role,
			})

			tokenString, err := token.SignedString([]byte(secretKey))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}

			resp := map[string]string{
				"token": tokenString,
			}
			err = json.NewEncoder(w).Encode(resp)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}
			return
		}
	}

	resp := map[string]string{
		"error": "user not found",
	}

	JSONError(w, resp, http.StatusNotFound)
	return
}

func (as *AuthServer) Endpoints() {
	router := mux.NewRouter()
	router.HandleFunc("/auth", as.Auth).Methods(http.MethodPost)

	as.Handler = handlers.LoggingHandler(os.Stdout, router)
}
