package stocks

import (
	"context"
	"errors"
	"math"

	"github.com/alexoh554/control-center/sqlc"
)

type Client struct{}

// Get stock purchases grouped by symbol, regardless of purchase date
func (c *Client) GetBySymbol(ctx context.Context, queries sqlc.Queries) (map[string]Holding, error) {
	stockPurchases, err := queries.GetStockPurchases(ctx)
	if err != nil {
		return nil, errors.New("error getting stock purchases")
	}

	holdings := make(map[string]Holding)
	for _, stock := range stockPurchases {
		if holding, ok := holdings[stock.Symbol]; ok {
			// Need to update total shares, cost, and average cost
			holding.Shares += int(stock.Quantity)
			holding.BookCostCents += int(stock.TotalPriceCents)
			holding.AveragePriceCents = int(math.Round(float64(holding.BookCostCents) / float64(holding.Shares)))
		} else {
			holdings[stock.Symbol] = Holding{
				Symbol:            stock.Symbol,
				Shares:            int(stock.Quantity),
				BookCostCents:     int(stock.TotalPriceCents),
				AveragePriceCents: int(stock.PriceCents),
			}
		}
	}

	return holdings, nil
}
