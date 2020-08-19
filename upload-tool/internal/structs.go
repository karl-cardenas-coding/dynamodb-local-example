package internal

// A strcuture for the JSON
type Model struct {
	OrderID    string `json:"orderId"`
	CustomerID string `json:"customerId"`
	Shipped    string `json:"shipped,omitempty"`
	Email      string `json:"email"`
	Address    string `json:"address"`
	Cost       string `json:"cost"`
}
