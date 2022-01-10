package httpServer

import "assignment/question_2/cashier"

type ReduceCashInput struct {
	CashList []ReduceCashData `json:"cashList"`
}

type ReduceCashData struct {
	Value  float32 `json:"value"`
	Amount uint8   `json:"amount"`
}

type ReduceResponse struct {
	Error  *ReduceErrorResponse  `json:"error"`
	Result *ReduceResultResponse `json:"result"`
}

type ReduceErrorResponse struct {
	Msg int `json:"msg"`
}

type ReduceResultResponse struct {
	Fail     int       `json:"fail"`
	Success  int       `json:"success"`
	FailList []float32 `json:"failList"`
}

func reduceCash(c *cashier.Cashier, input *ReduceCashInput) (ReduceResponse, cashier.Cashier) {
	deductCashList := input.CashList
	var failList []float32
	fail := 0
	success := 0
	for _, cash := range deductCashList {
		_, err := c.ReduceCash(cash.Value, cash.Amount)
		if err != nil {
			failList = append(failList, cash.Value)
			fail = fail + 1
		} else {
			success = success + 1
		}
	}
	result := ReduceResultResponse{
		Fail:     fail,
		Success:  success,
		FailList: failList,
	}
	return ReduceResponse{
		Error:  nil,
		Result: &result,
	}, *c
}
