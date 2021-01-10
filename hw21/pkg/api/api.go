package api

import (
	"encoding/json"
	"fmt"
	"hw21/pkg/storage"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Service struct {
	Router *mux.Router
	strg   storage.Interface
}

func New(router *mux.Router, strg storage.Interface) *Service {
	s := &Service{
		Router: router,
		strg:   strg,
	}
	s.endpoints()

	return s
}

func (s *Service) endpoints() {
	s.Router.HandleFunc("/film", s.addFilmsHandler).Methods(http.MethodPost, http.MethodOptions)
	s.Router.HandleFunc("/film/{id}", s.filmsHandler).Methods(http.MethodGet, http.MethodOptions)
	s.Router.HandleFunc("/film", s.updateFilmHandler).Methods(http.MethodPut, http.MethodOptions)
	s.Router.HandleFunc("/film", s.deleteFilmHandler).Methods(http.MethodDelete, http.MethodOptions)
}

func (s *Service) addFilmsHandler(w http.ResponseWriter, r *http.Request) {
	var films []storage.Film

	err := json.NewDecoder(r.Body).Decode(&films)
	if err != nil {
		http.Error(w, "reading body error: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	fmt.Println(films)

	err = s.strg.AddFilms(films)
	if err != nil {
		http.Error(w, "inserting data error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Service) deleteFilmHandler(w http.ResponseWriter, r *http.Request) {
	var film storage.Film

	err := json.NewDecoder(r.Body).Decode(&film)
	if err != nil {
		http.Error(w, "reading body error: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if film.ID == 0 {
		http.Error(w, "you must pass the record id", http.StatusInternalServerError)
		return
	}

	err = s.strg.DeleteFilm(film)
	if err != nil {
		http.Error(w, "deleting data error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Service) updateFilmHandler(w http.ResponseWriter, r *http.Request) {
	var film storage.Film

	err := json.NewDecoder(r.Body).Decode(&film)
	if err != nil {
		http.Error(w, "reading body error: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if film.ID == 0 {
		http.Error(w, "you must pass the record id", http.StatusInternalServerError)
		return
	}

	err = s.strg.UpdateFilm(film)
	if err != nil {
		http.Error(w, "updating data error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Service) filmsHandler(w http.ResponseWriter, r *http.Request) {
	var films []storage.Film
	var err error

	vars := mux.Vars(r)

	if _, ok := vars["id"]; !ok {
		films, err = s.strg.Films(0)
		if err != nil {
			http.Error(w, "selecting data error: "+err.Error(), http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(films)
		return
	}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "id must be integer: "+err.Error(), http.StatusBadRequest)
		return
	}

	films, err = s.strg.Films(int64(id))
	if err != nil {
		http.Error(w, "selecting data error: "+err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(films)
}
