package quotes

// Contains minimal data about a quote
type Quote struct {
	Symbol string

	RegularMarketChangePercent float64
	RegularMarketPrice         float64
	RegularMarketChange        float64
	RegularMarketTime          int
}
