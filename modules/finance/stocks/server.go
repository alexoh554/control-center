package stocks

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/alexoh554/control-center/common"
	"github.com/alexoh554/control-center/sqlc"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Server struct {
	*sqlc.Queries
}

func NewServer(queries *sqlc.Queries) *Server {
	return &Server{Queries: queries}
}

func (s *Server) Register(router *mux.Router) {
	router.HandleFunc("/finance/stocks", s.Create).Methods("POST")
	router.HandleFunc("/finance/stocks/{id}", s.Delete).Methods("DELETE")
}

// Creates an entry denoting a stock purchase at a given time (not real time)
func (s *Server) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Symbol          string    `json:"symbol"`
		PriceCents      int32     `json:"price_cents"`
		Quantity        int32     `json:"quantity"`
		PurchasedAt     time.Time `json:"purchased_at"`
		TotalPriceCents int32     `json:"total_price_cents"`
		Broker          string    `json:"broker"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var broker sql.NullString
	if req.Broker != "" {
		broker = sql.NullString{String: req.Broker, Valid: true}
	}

	stockPurchase, err := s.CreateStockPurchase(r.Context(), sqlc.CreateStockPurchaseParams{
		Symbol:          req.Symbol,
		PriceCents:      req.PriceCents,
		Quantity:        req.Quantity,
		PurchasedAt:     req.PurchasedAt,
		TotalPriceCents: req.TotalPriceCents,
		Broker:          broker,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.JsonResponse(w, stockPurchase)
}

func (s *Server) Delete(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stockPurchaseID, err := uuid.Parse(req.ID)
	if err != nil {
		http.Error(w, "Invalid UUID format", http.StatusBadRequest)
	}

	err = s.DeleteStockPurchase(r.Context(), stockPurchaseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.JsonResponse(w, map[string]string{"message": "Stock purchase deleted successfully"})
}
