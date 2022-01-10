package httpServer

import "assignment/question_2/cashier"

type SetCashInput struct {
	CashList []SetCashData `json:"cashList"`
}

type SetCashData struct {
	Value  float32 `json:"value"`
	Amount uint8   `json:"amount"`
}

type SetResponse struct {
	Error  *SetErrorResponse  `json:"error"`
	Result *SetResultResponse `json:"result"`
}

type SetErrorResponse struct {
	Msg int `json:"msg"`
}

type SetResultResponse struct {
	Fail     int       `json:"fail"`
	Success  int       `json:"success"`
	FailList []float32 `json:"failList"`
}

func setCash(c *cashier.Cashier, input *SetCashInput) (SetResponse, cashier.Cashier) {
	setCashList := input.CashList
	var failList []float32
	fail := 0
	success := 0
	for _, cash := range setCashList {
		_, err := c.SetCash(cash.Value, cash.Amount)
		if err != nil {
			failList = append(failList, cash.Value)
			fail = fail + 1
		} else {
			success = success + 1
		}
	}
	result := SetResultResponse{
		Fail:     fail,
		Success:  success,
		FailList: failList,
	}
	return SetResponse{
		Error:  nil,
		Result: &result,
	}, *c
}
