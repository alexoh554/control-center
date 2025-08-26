package todos

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/alexoh554/control-center/common"
	"github.com/alexoh554/control-center/sqlc"
	"github.com/gorilla/mux"
)

type Server struct {
	db *sql.DB
}

func NewServer(db *sql.DB) *Server {
	return &Server{db: db}
}

func (s *Server) Register(router *mux.Router) {
	router.HandleFunc("/todos", s.Create).Methods("POST")
}

func (s *Server) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title  string `json:"title"`
		Status string `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	queries := sqlc.New(s.db)
	todo, err := queries.CreateTodo(r.Context(), sqlc.CreateTodoParams{
		Title:  req.Title,
		Status: sql.NullString{String: req.Status, Valid: true},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.JsonResponse(w, todo)
}
