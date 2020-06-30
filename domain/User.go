package domain

type User struct {
	ID           int    `json:"id"`
	Source       string `json:"source"`
	Phone        string `json:"phone"`
	FullName     string `json:"full_name"`
	Address      string `json:"address"`
	FacebookID   string `json:"facebook_id"`
	Gender       string `json:"gender"`
	Link         string `json:"link"`
	ProfileImage string `json:"profile_image"`
	Tag          string `json:"tag"`
	AccountID    string `json:"account_id"`
}
