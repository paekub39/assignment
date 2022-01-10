package httpServer

import (
	"assignment/question_2/cashier"
)

type InsertCashInput struct {
	CashList []InsertCashData `json:"cashList"`
}

type InsertCashData struct {
	Value  float32 `json:"value"`
	Amount uint8   `json:"amount"`
}

type InsertResponse struct {
	Error  *InsertErrorResponse  `json:"error"`
	Result *InsertResultResponse `json:"result"`
}

type InsertErrorResponse struct {
	Msg int `json:"msg"`
}

type InsertResultResponse struct {
	Fail     int       `json:"fail"`
	Success  int       `json:"success"`
	FailList []float32 `json:"failList"`
}

func insertCash(c *cashier.Cashier, input *InsertCashInput) (InsertResponse, cashier.Cashier) {
	insertCashList := input.CashList
	var failList []float32
	fail := 0
	success := 0
	for _, cash := range insertCashList {
		_, err := c.AddCash(cash.Value, cash.Amount)
		if err != nil {
			failList = append(failList, cash.Value)
			fail = fail + 1
		} else {
			success = success + 1
		}
	}
	result := InsertResultResponse{
		Fail:     fail,
		Success:  success,
		FailList: failList,
	}
	return InsertResponse{
		Error:  nil,
		Result: &result,
	}, *c
}
