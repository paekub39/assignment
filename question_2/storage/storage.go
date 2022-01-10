package storage

import (
	chError "assignment/question_2/error"
	"fmt"
	"math"
)

const (
	OneThousandCash  float32 = 1000
	FiveHundredCash  float32 = 500
	OneHundredCash   float32 = 100
	FiftyCash        float32 = 50
	TwentyCash       float32 = 20
	TenCash          float32 = 10
	FiveCash         float32 = 5
	OneCash          float32 = 1
	PointTwoFiveCash float32 = 0.25

	OneThousandLimit  uint8 = 10
	FiveHundredLimit  uint8 = 20
	OneHundredLimit   uint8 = 15
	FiftyLimit        uint8 = 20
	TwentyLimit       uint8 = 30
	TenLimit          uint8 = 20
	FiveLimit         uint8 = 20
	OneLimit          uint8 = 20
	PointTwoFiveLimit uint8 = 50
)

type Storage struct {
	CashList map[float32]CashTemplate
}

type CashTemplate struct {
	Limit         uint8
	CurrentAmount uint8
}

func GetAllowCash() []float32 {
	result := []float32{OneThousandCash, FiveHundredCash, OneHundredCash, FiftyCash, TwentyCash, TenCash, FiveCash, OneCash, PointTwoFiveCash}
	return result
}

func GetCashLimit(cash float32) (uint8, error) {
	switch cash {
	case OneThousandCash:
		return OneThousandLimit, nil
	case FiveHundredCash:
		return FiveHundredLimit, nil
	case OneHundredCash:
		return OneHundredLimit, nil
	case FiftyCash:
		return FiftyLimit, nil
	case TwentyCash:
		return TwentyLimit, nil
	case TenCash:
		return TenLimit, nil
	case FiveCash:
		return FiveLimit, nil
	case OneCash:
		return OneLimit, nil
	case PointTwoFiveCash:
		return PointTwoFiveLimit, nil
	default:
		msg := fmt.Sprintf("%f dose not support", cash)
		err := &msg
		return 0, chError.CashierError{
			Msg: *err,
		}
	}
}

func IsInAllowCash(cash float32) bool {
	cashList := GetAllowCash()
	result := false
	for _, c := range cashList {
		if c == cash {
			result = true
			break
		}
	}
	return result
}

func (s *Storage) AddCash(cash float32, amount uint8) (*Storage, error) {
	if !IsInAllowCash(cash) {
		msg := fmt.Sprintf("%f dose not support", cash)
		err := &msg
		return s, chError.CashierError{
			Msg: *err,
		}
	}
	cashList := s.CashList
	cashValue, ok := cashList[cash]
	if !ok {
		cashLimit, _ := GetCashLimit(cash)
		if cashLimit >= amount {
			newCash := CashTemplate{
				Limit:         cashLimit,
				CurrentAmount: amount,
			}
			s.CashList[cash] = newCash
			return s, nil
		} else {
			msg := fmt.Sprintf("CashTemplate amount are over limit. The limit of %f is %d", cash, cashValue.Limit)
			err := &msg
			return s, chError.CashierError{
				Msg: *err,
			}
		}
	} else {
		newAmount := cashValue.CurrentAmount + amount
		if cashValue.Limit >= newAmount {
			cashTemp := CashTemplate{
				Limit:         cashValue.Limit,
				CurrentAmount: newAmount,
			}
			s.CashList[cash] = cashTemp
			return s, nil
		} else {
			msg := fmt.Sprintf("CashTemplate amount are over limit. The limit of %f is %d", cash, cashValue.Limit)
			err := &msg
			return s, chError.CashierError{
				Msg: *err,
			}
		}
	}
}

func (s *Storage) SetCash(cash float32, amount uint8) (*Storage, error) {

	if !IsInAllowCash(cash) {
		msg := fmt.Sprintf("%f dose not support", cash)
		err := &msg
		return nil, chError.CashierError{
			Msg: *err,
		}
	}

	cashLimit, _ := GetCashLimit(cash)
	if cashLimit >= amount {
		newCash := CashTemplate{
			Limit:         cashLimit,
			CurrentAmount: amount,
		}
		s.CashList[cash] = newCash
		return s, nil
	} else {
		msg := fmt.Sprintf("CashTemplate amount are over limit. The limit of %f is %d", cash, cashLimit)
		err := &msg
		return nil, chError.CashierError{
			Msg: *err,
		}
	}
}

func (s *Storage) ReduceCash(cash float32, amount uint8) (*Storage, error) {

	if !IsInAllowCash(cash) {
		msg := fmt.Sprintf("%f dose not support", cash)
		err := &msg
		return nil, chError.CashierError{
			Msg: *err,
		}
	}
	cashList := s.CashList
	cashValue, ok := cashList[cash]
	if !ok {
		cashLimit, _ := GetCashLimit(cash)
		newCash := CashTemplate{
			Limit:         cashLimit,
			CurrentAmount: 0,
		}
		if amount > newCash.CurrentAmount {
			msg := fmt.Sprintf("CashTemplate amount are over the current amount. The current amount of %f is %d. Input is %d", cash, newCash.CurrentAmount, amount)
			err := &msg
			return nil, chError.CashierError{
				Msg: *err,
			}
		} else {
			s.CashList[cash] = newCash
			return s, nil
		}
	} else {
		if cashValue.CurrentAmount >= amount {
			newAmount := cashValue.CurrentAmount - amount
			cashTemp := CashTemplate{
				Limit:         cashValue.Limit,
				CurrentAmount: newAmount,
			}
			s.CashList[cash] = cashTemp
			return s, nil
		} else {
			msg := fmt.Sprintf("CashTemplate amount are over the current amount. The current amount of %f is %d. Input is %d", cash, cashValue.CurrentAmount, amount)
			err := &msg
			return nil, chError.CashierError{
				Msg: *err,
			}
		}
	}
}

func (s Storage) CalculateChangeWithCurrentCash(input float32) (Storage, map[float32]uint8, float32) {
	cashAllowanceList := GetAllowCash()
	remaining := input
	var change = make(map[float32]uint8)
	for _, c := range cashAllowanceList {
		reduction, r := CalculateCashReductionAmountAndRemaining(remaining, c, float32(s.CashList[c].CurrentAmount))
		change[c] = reduction
		cashTemp := CashTemplate{
			Limit:         s.CashList[c].Limit,
			CurrentAmount: s.CashList[c].CurrentAmount - reduction,
		}
		s.CashList[c] = cashTemp
		remaining = r
	}
	return s, change, remaining
}

func CalculateCashReductionAmountAndRemaining(input float32, value float32, currentAmount float32) (uint8, float32) {
	cashReductionAmount := input / value
	paidRemaining := math.Mod(float64(input), float64(value))
	if currentAmount < cashReductionAmount {
		over := cashReductionAmount - currentAmount
		cashReductionAmount = currentAmount
		paidRemaining = paidRemaining + float64(over*value)
	}
	return uint8(cashReductionAmount), float32(paidRemaining)
}

func (s *Storage) GetAllRemainingCashAmountInStorage() float32 {
	cashList := s.CashList
	result := float32(0)
	for cash, cashValue := range cashList {
		result = result + (cash * float32(cashValue.CurrentAmount))
	}
	return result
}

func (s *Storage) RefreshCurrentAmount() {
	cashList := s.CashList
	allowanceList := GetAllowCash()
	for _, cash := range allowanceList {
		cashValue, ok := cashList[cash]
		if ok {
			if cashValue.CurrentAmount > cashValue.Limit {
				cashValue.CurrentAmount = cashValue.Limit
				s.CashList[cash] = cashValue
			}
		} else {
			lim, _ := GetCashLimit(cash)
			newCash := CashTemplate{
				Limit:         lim,
				CurrentAmount: 0,
			}
			s.CashList[cash] = newCash
		}
	}
}
