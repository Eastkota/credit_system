package model

type GenericSuccessData struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type GenericResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Error *Error      `json:"error,omitempty"`
}

type CustomerResult struct {
	Customer *Customer `json:"customer"`
}

type CreditResult struct {
	Credit *Credit `json:"credit"`
}

type BalanceResult struct {
	Balance *BalanceUpdate `json:"balance"`
}

type MultipleBalanceResult struct {
	Balances []BalanceUpdate `json:"balances"`
}
