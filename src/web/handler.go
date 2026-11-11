package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/DrDowellsHead/Go_Bootcamp_4/src/domain"
)

type Handler struct {
	service domain.Service
}

func NewHandler(service domain.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /game", h.create)
	mux.HandleFunc("GET /game/{id}", h.get)
	mux.HandleFunc("POST /game/{id}", h.play)
	return mux
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	game, err := h.service.Create()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "cannot create game")
		return
	}
	writeJSON(w, http.StatusCreated, toWeb(game, h.service))
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := validateID(id); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	game, err := h.service.Get(id)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, toWeb(game, h.service))
}

func (h *Handler) play(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := validateID(id); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 16*1024)
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		writeError(w, http.StatusBadRequest, fmt.Sprintf("cannot read request: %v", err))
		return
	}
	if err := validateBoardShape(body); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	var input Game
	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, fmt.Sprintf("invalid JSON: %v", err))
		return
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		writeError(w, http.StatusBadRequest, "request body must contain exactly one JSON object")
		return
	}
	if input.ID != "" && input.ID != id {
		writeError(w, http.StatusBadRequest, "game ID in body must match URL")
		return
	}
	input.ID = id
	game, err := h.service.Play(id, toDomain(input).Board)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, toWeb(game, h.service))
}

func validateID(id string) error {
	if len(id) != 36 {
		return errors.New("invalid UUID: expected 36 characters")
	}
	for index, character := range id {
		if index == 8 || index == 13 || index == 18 || index == 23 {
			if character != '-' {
				return errors.New("invalid UUID: missing separator")
			}
			continue
		}
		if !strings.ContainsRune("0123456789abcdefABCDEF", character) {
			return errors.New("invalid UUID: expected hexadecimal digits")
		}
	}
	return nil
}

func writeServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrNotFound):
		writeError(w, http.StatusNotFound, err.Error())
	case errors.Is(err, domain.ErrInvalidMove), errors.Is(err, domain.ErrGameFinished):
		writeError(w, http.StatusBadRequest, err.Error())
	default:
		writeError(w, http.StatusInternalServerError, "internal server error")
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, ErrorResponse{Error: message})
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}
