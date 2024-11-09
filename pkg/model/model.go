package model

import "time"

type Transaction struct {
	Date    time.Time `json:"date" bson:"date_time"`
	TransNo int       `json:"trans_no" bson:"trans_no"`
	Credit  float64   `json:"credit" bson:"credit"`
	Debit   float64   `json:"debit" bson:"debit"`
	Detail  string    `json:"detail" bson:"detail"`
}

type TransactionResponse struct {
	Transactions []Transaction `json:"transactions"`
	Total        int           `json:"total"`
	Page         int           `json:"page"`
	TotalElement int           `json:"total_element"`
	TotalPage    int           `json:"total_page"`
}
