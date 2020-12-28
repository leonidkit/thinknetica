package api

import (
	"encoding/json"
	"gosearch/pkg/engine"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Service struct {
	Router *mux.Router
	engine engine.Service
}

func New(router *mux.Router, engine engine.Service) *Service {
	s := Service{
		Router: router,
		engine: engine,
	}
	s.endpoints()

	return &s
}

func (s *Service) endpoints() {
	s.Router.HandleFunc("/api/search/{query}", s.Search).Methods("GET")
	s.Router.HandleFunc("/api/docs", s.Docs).Methods("GET")
	s.Router.HandleFunc("/api/create", s.Create).Methods("POST")
	s.Router.HandleFunc("/api/delete", s.Delete).Methods("POST")
	s.Router.HandleFunc("/api/update", s.Update).Methods("POST")
}

func (s *Service) Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	query := mux.Vars(r)["query"]
	result, err := s.engine.Search(query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(result)
}

func (s *Service) Docs(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	result := s.engine.Tree.Index()
	json.NewEncoder(w).Encode(result)
}

func (s *Service) Create(w http.ResponseWriter, r *http.Request) {
	var req map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	r.Body.Close()

	url, ok := req["url"].(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	title, ok := req["title"].(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if url == "" || title == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.engine.Tree.Add(url, title)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *Service) Delete(w http.ResponseWriter, r *http.Request) {
	var req map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	r.Body.Close()

	id, ok := req["id"].(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if id != "" {
		iid, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		s.engine.Tree.Delete(uint64(iid))
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

func (s *Service) Update(w http.ResponseWriter, r *http.Request) {
	var req map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	r.Body.Close()

	id, ok := req["id"].(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	title, ok := req["title"].(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	url, ok := req["url"].(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if id != "" && title != "" && url != "" {
		iid, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		s.engine.Tree.Update(uint64(iid), url, title)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}
