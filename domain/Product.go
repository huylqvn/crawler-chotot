package domain

type Product struct {
	ID        string `json:"id"`
	Source    string `json:"source"`
	AccountID string `json:"account_id"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	Price     int    `json:"price"`
	Image     string `json:"image"`
	Tag       string `json:"tag"`
}
