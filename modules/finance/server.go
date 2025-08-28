package finance

import (
	"database/sql"

	"github.com/alexoh554/control-center/modules/finance/stocks"
	"github.com/alexoh554/control-center/sqlc"
	"github.com/gorilla/mux"
)

type Server struct {
	portfolio stocks.Server
}

func NewServer(db *sql.DB) *Server {
	queries := sqlc.New(db)
	return &Server{
		portfolio: *stocks.NewServer(queries),
	}
}

func (s *Server) Register(router *mux.Router) {
	s.portfolio.Register(router)
}
