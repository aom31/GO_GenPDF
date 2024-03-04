package models

type ReceiptDocInfo struct {
	AddressTitle    AddressInfo  `json:"addressTitle"`
	ReceiptNameFrom string       `json:"receiptNameFrom"`
	Date            string       `json:"date"`
	ReceiptNumber   string       `json:"receiptNumber"`
	AccountNumber   string       `json:"accountNumber"`
	InterestRate    float32      `json:"interestRate"`
	LoanInfo        LoanListInfo `json:"loanInfo"`
}

type LoanListInfo struct {
	PrincipalBalance          float32 `json:"principalBalance"`
	PrincipalPaymentAmount    float32 `json:"principalPaymentAmount"`
	PrincipalBalanceRemaining float32 `json:"principalBalanceRemaining"`
	InterestRatePayment       float32 `json:"interestRatePayment"`
	Other                     float32 `json:"other"`
	SummaryPayment            float32 `json:"summaryPayment"`
	SummaryBalanceRemaining   float32 `json:"summaryBalanceRemaining"`
	PaymentWithinDate         string  `json:"paymentWithinDate"`
}
