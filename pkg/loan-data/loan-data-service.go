package loandata

import "github.com/aom31/generatepdfgo/internal/models"

type LoanData struct {
	data models.ReceiptDocInfo
}

func NewLoanData(dataReceipt models.ReceiptDocInfo) *LoanData {
	return &LoanData{
		data: dataReceipt,
	}
}

func (loan *LoanData) CalculatePrincipalBalanceRemaining() float32 {
	totalPrincipalBalanceRemaining := loan.data.LoanInfo.PrincipalBalance - loan.data.LoanInfo.PrincipalPaymentAmount
	return totalPrincipalBalanceRemaining
}

func (loan *LoanData) CalculateSummaryPayment() float32 {
	totalSummaryPayment := loan.data.LoanInfo.PrincipalPaymentAmount + loan.data.LoanInfo.InterestRatePayment
	return totalSummaryPayment
}

func (loan *LoanData) CalculateSummaryAmountRemaining() float32 {
	totalSummaryBalanceRemaining := loan.data.LoanInfo.PrincipalBalance - loan.data.LoanInfo.SummaryPayment
	return totalSummaryBalanceRemaining
}
