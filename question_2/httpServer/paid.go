package httpServer

import (
	"assignment/question_2/cashier"
	"assignment/question_2/storage"
)

type PaidInput struct {
	Paid  []PaidData `json:"paid"`
	Price float32    `json:"price"`
}

type PaidData struct {
	Value  float32 `json:"value"`
	Amount uint8   `json:"amount"`
}

type PaidResponse struct {
	Error  *PaidErrorResponse  `json:"error"`
	Result *PaidResultResponse `json:"result"`
}

type PaidErrorResponse struct {
	Msg string `json:"msg"`
}

type PaidResultResponse struct {
	ChangeList  []PaidCashResponse `json:"changeList"`
	TotalChange float32            `json:"totalChange"`
}

type PaidCashResponse struct {
	Value  float32 `json:"value"`
	Amount uint8   `json:"amount"`
}

func paidCash(c *cashier.Cashier, input *PaidInput) (PaidResponse, cashier.Cashier) {
	cashList := input.Paid
	mapCash := make(map[float32]uint8)
	for _, cash := range cashList {
		mapCash[cash.Value] = cash.Amount
	}

	changeMap, err := c.PayByCash(input.Price, mapCash)
	if err != nil {
		per := PaidErrorResponse{
			Msg: err.Error(),
		}
		resp := PaidResponse{
			Error:  &per,
			Result: nil,
		}
		return resp, *c
	}
	changeMapTemp := *changeMap
	var pcrList []PaidCashResponse
	totalChange := float32(0)
	allowanceCashList := storage.GetAllowCash()
	for _, allowance := range allowanceCashList {
		c, ok := changeMapTemp[allowance]
		if !ok {
			continue
		}
		totalChange = totalChange + (allowance * float32(c))
		prs := PaidCashResponse{
			Value:  allowance,
			Amount: c,
		}
		pcrList = append(pcrList, prs)
	}

	result := PaidResultResponse{
		ChangeList:  pcrList,
		TotalChange: totalChange,
	}

	resp := PaidResponse{
		Error:  nil,
		Result: &result,
	}

	return resp, *c
}
