package model

type CompanyInfo struct {
	Floor     string  		`example:"3" json:"floor"`
	Name      string 	`example:"喜登數位" json:"name"`
}

type ApiResponse struct {
	Code    int
	Massage string
	Data    interface{}
}