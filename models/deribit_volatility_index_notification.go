package models

type DeribitVolatilityIndexNotification struct {
	Timestamp  int64   `json:"timestamp"`
	Volatility float64 `json:"volatility"`
	IndexName  string  `json:"index_name"`
}
