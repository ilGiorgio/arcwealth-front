package domain

type ApiResponse[T any] struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	Data      T      `json:"data"`
	Count     int    `json:"count"`
	Timestamp string `json:"timestamp"`
}

type Currency struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	CurrencyCode string `json:"currencyCode"`
}
