package api

import (
	"encoding/json"
	"gosearch/pkg/crawler"
	"gosearch/pkg/engine"
	"net/http"

	"github.com/gorilla/mux"
)

type Service struct {
	Router *mux.Router
	engine *engine.Service
}

func New(router *mux.Router, engine *engine.Service) *Service {
	s := Service{
		Router: router,
		engine: engine,
	}
	s.endpoints()

	return &s
}

func (s *Service) endpoints() {
	s.Router.HandleFunc("/api/search/{query}", s.Search).Methods(http.MethodGet)
	s.Router.HandleFunc("/api/docs", s.Docs).Methods(http.MethodGet)
	s.Router.HandleFunc("/api/create_doc", s.Create).Methods(http.MethodPost)
	s.Router.HandleFunc("/api/delete_doc", s.Delete).Methods(http.MethodPost)
	s.Router.HandleFunc("/api/update_doc", s.Update).Methods(http.MethodPost)
}

func (s *Service) Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	query := mux.Vars(r)["query"]
	result, err := s.engine.Search(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func (s *Service) Docs(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(s.engine.Data)
}

func (s *Service) Create(w http.ResponseWriter, r *http.Request) {
	var req crawler.Document

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	r.Body.Close()

	if req.Title == "" || req.URL == "" {
		http.Error(w, "\"url\", \"title\" fields should not be empty", http.StatusBadRequest)
		return
	}

	err = s.engine.Tree.Add(req.URL, req.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (s *Service) Delete(w http.ResponseWriter, r *http.Request) {
	var req crawler.Document

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	r.Body.Close()

	if req.ID == 0 {
		http.Error(w, "\"id\" field should not be empty", http.StatusBadRequest)
		return
	}

	err = s.engine.Tree.Delete(req.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (s *Service) Update(w http.ResponseWriter, r *http.Request) {
	var req crawler.Document

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	r.Body.Close()

	if req.ID == 0 || req.URL == "" || req.Title == "" {
		http.Error(w, "\"url\", \"title\", \"id\" fields should not be empty", http.StatusBadRequest)
		return
	}

	err = s.engine.Tree.Update(req.ID, req.URL, req.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
