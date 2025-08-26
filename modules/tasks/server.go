package tasks

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/alexoh554/control-center/common"
	"github.com/alexoh554/control-center/sqlc"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Server struct {
	*sqlc.Queries
}

func NewServer(db *sql.DB) *Server {
	queries := sqlc.New(db)
	return &Server{Queries: queries}
}

func (s *Server) Register(router *mux.Router) {
	router.HandleFunc("/tasks/create", s.Create).Methods("POST")
	router.HandleFunc("/tasks/{taskID}", s.Update).Methods("PATCH")
	router.HandleFunc("/tasks/{taskID}", s.Delete).Methods("DELETE")
}

func (s *Server) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := s.CreateTask(r.Context(), sqlc.CreateTaskParams{
		Title:       req.Title,
		Description: sql.NullString{String: req.Description, Valid: true},
		Status:      sql.NullString{String: req.Status, Valid: true},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.JsonResponse(w, task)
}

func (s *Server) Update(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID          string `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	taskID, err := uuid.Parse(req.ID)
	if err != nil {
		http.Error(w, "Invalid UUID format", http.StatusBadRequest)
		return
	}

	task, err := s.UpdateTask(r.Context(), sqlc.UpdateTaskParams{
		ID:          taskID,
		Title:       req.Title,
		Description: sql.NullString{String: req.Description, Valid: true},
		Status:      sql.NullString{String: req.Status, Valid: true},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.JsonResponse(w, task)
}

func (s *Server) Delete(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	taskID, err := uuid.Parse(req.ID)
	if err != nil {
		http.Error(w, "Invalid UUID format", http.StatusBadRequest)
		return
	}

	task, err := s.DeleteTask(r.Context(), taskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.JsonResponse(w, task)
}
