package structs

type Transaction struct {
	RequestId int `json:"request_id" form:"request_id" query:"request_id" xml:"request_id"`
	TransactionRecords []TransactionRecord `json:"data"`
}

type TransactionRecord struct {
	Customer string `json:"customer"`
	Quantity int `json:"quantity"`
	Price float64 `json:"price"`
}

type TransactionFilter struct {
	Limit int `json:"limit" form:"limit" query:"limit"`
	Offset int `json:"offset" form:"offset" query:"offset"`
	Filter string `json:"filter" form:"filter" query:"filter"`
}