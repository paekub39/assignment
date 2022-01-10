package httpServer

import (
	"assignment/question_2/cashier"
	"assignment/question_2/storage"
)

type AllCashRemainingResponse struct {
	Error  *AllCashRemainingErrorResponse  `json:"error"`
	Result *AllCashRemainingResultResponse `json:"result"`
}

type AllCashRemainingErrorResponse struct {
	Msg int `json:"msg"`
}

type AllCashRemainingResultResponse struct {
	CashList []AllCashRemainingCashResponse `json:"cashList"`
	Total    float32                        `json:"total"`
}

type AllCashRemainingCashResponse struct {
	Value  float32 `json:"value"`
	Amount uint8   `json:"amount"`
}

func getAllCashRemaining(c *cashier.Cashier) AllCashRemainingResponse {
	total := c.Storage.GetAllRemainingCashAmountInStorage()
	mapCashList := c.GetAllRemainingCash()
	allowance := storage.GetAllowCash()
	var cashList []AllCashRemainingCashResponse
	for _, k := range allowance {
		v, ok := mapCashList[k]
		if ok {
			cr := AllCashRemainingCashResponse{
				Value:  k,
				Amount: v,
			}
			cashList = append(cashList, cr)
		} else {
			cr := AllCashRemainingCashResponse{
				Value:  k,
				Amount: 0,
			}
			cashList = append(cashList, cr)
		}
	}
	result := AllCashRemainingResultResponse{
		CashList: cashList,
		Total:    total,
	}
	return AllCashRemainingResponse{
		Error:  nil,
		Result: &result,
	}
}
