package response

type AdListing struct {
	Total int  `json:"total"`
	Ads   []Ad `json:"ads"`
}

type Ad struct {
	AccountOid   string `json:"account_oid"`
	CategoryName string `json:"category_name"`
	Subject      string `json:"subject"`
	Body         string `json:"body"`
	Price        int    `json:"price"`
}

type Profile struct {
	Avatar     string `json:"avatar"`
	Address    string `json:"address"`
	FullName   string `json:"full_name"`
	Gender     string `json:"gender"`
	Phone      string `json:"phone"`
	FacebookID string `json:"facebook_id"`
}

type AdInfo struct {
	Ad struct {
		Body   string   `json:"body"`
		Images []string `json:"images"`
	} `json:"ad"`
}
