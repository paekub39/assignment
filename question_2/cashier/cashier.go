package cashier

import (
	"assignment/question_2/config"
	chError "assignment/question_2/error"
	"assignment/question_2/storage"
	"fmt"
)

type Cashier struct {
	Storage storage.Storage
}

func NewCashier() Cashier {
	st, e := config.ReadTemp()
	if e != nil {
		fmt.Println(e)
		cashList := make(map[float32]storage.CashTemplate)

		cashOneThousand := storage.CashTemplate{
			Limit:         storage.OneThousandLimit,
			CurrentAmount: 0,
		}
		cashList[storage.OneThousandCash] = cashOneThousand

		cashFiveHundred := storage.CashTemplate{
			Limit:         storage.FiveHundredLimit,
			CurrentAmount: 0,
		}
		cashList[storage.FiveHundredCash] = cashFiveHundred

		cashOneHundred := storage.CashTemplate{
			Limit:         storage.OneHundredLimit,
			CurrentAmount: 0,
		}
		cashList[storage.OneHundredCash] = cashOneHundred

		cashFifty := storage.CashTemplate{
			Limit:         storage.FiftyLimit,
			CurrentAmount: 0,
		}
		cashList[storage.FiftyCash] = cashFifty

		cashTwenty := storage.CashTemplate{
			Limit:         storage.TwentyLimit,
			CurrentAmount: 0,
		}
		cashList[storage.TwentyCash] = cashTwenty

		cashTen := storage.CashTemplate{
			Limit:         storage.TenLimit,
			CurrentAmount: 0,
		}
		cashList[storage.TenCash] = cashTen

		cashFive := storage.CashTemplate{
			Limit:         storage.FiveLimit,
			CurrentAmount: 0,
		}
		cashList[storage.FiveCash] = cashFive

		cashOne := storage.CashTemplate{
			Limit:         storage.OneLimit,
			CurrentAmount: 0,
		}
		cashList[storage.OneCash] = cashOne

		cashPointTwoFive := storage.CashTemplate{
			Limit:         storage.PointTwoFiveLimit,
			CurrentAmount: 0,
		}
		cashList[storage.PointTwoFiveCash] = cashPointTwoFive

		s := storage.Storage{
			CashList: cashList,
		}
		st = &s
		config.WriteTemp(st)
	}
	cashier := Cashier{
		Storage: *st,
	}
	return cashier
}

func (c *Cashier) AddCash(cash float32, amount uint8) (*Cashier, error) {
	s := c.Storage
	_, e := s.AddCash(cash, amount)
	if e != nil {
		return nil, e
	}
	c.Storage = s
	config.WriteTemp(&c.Storage)
	return c, nil
}

func (c *Cashier) SetCash(cash float32, amount uint8) (*Cashier, error) {
	s := c.Storage
	_, e := s.SetCash(cash, amount)
	if e != nil {
		return nil, e
	}
	c.Storage = s
	config.WriteTemp(&c.Storage)
	return c, nil
}

func (c *Cashier) ReduceCash(cash float32, amount uint8) (*Cashier, error) {
	s := c.Storage
	_, e := s.ReduceCash(cash, amount)
	if e != nil {
		return nil, e
	}
	c.Storage = s
	config.WriteTemp(&c.Storage)
	return c, nil
}

func (c *Cashier) PayByCash(price float32, paid map[float32]uint8) (*map[float32]uint8, error) {
	var paidAmount float32 = 0
	cashAllowance := storage.GetAllowCash()
	cashListTemp := c.Storage.CashList
	for _, allowance := range cashAllowance {
		tempAmount, tOk := cashListTemp[allowance]
		pAmount, pOk := paid[allowance]
		if pOk {
			paidAmount = paidAmount + (allowance * float32(pAmount))
			currentAmountTemp := pAmount
			lim, _ := storage.GetCashLimit(allowance)
			if tOk {
				currentAmountTemp = currentAmountTemp + tempAmount.CurrentAmount
			}
			newCash := storage.CashTemplate{
				Limit:         lim,
				CurrentAmount: currentAmountTemp,
			}
			cashListTemp[allowance] = newCash
		}
	}
	if paidAmount < price {
		msg := fmt.Sprintf("input money is less than the price")
		return nil, chError.CashierError{
			Msg: msg,
		}
	}
	storageTemp := storage.Storage{
		CashList: cashListTemp,
	}
	change := paidAmount - price
	newStorage, changeList, remaining := storageTemp.CalculateChangeWithCurrentCash(change)
	if remaining > 0 {
		msg := fmt.Sprintf("Note is not enough for the change")
		return nil, chError.CashierError{
			Msg: msg,
		}
	}
	newStorage.RefreshCurrentAmount()
	c.Storage = newStorage
	config.WriteTemp(&c.Storage)
	return &changeList, nil
}

func (c *Cashier) GetAllRemainingCash() map[float32]uint8 {
	result := make(map[float32]uint8)
	cashList := c.Storage.CashList
	allowanceCashList := storage.GetAllowCash()
	for _, allowanceCash := range allowanceCashList {
		cash, ok := cashList[allowanceCash]
		if ok {
			result[allowanceCash] = cash.CurrentAmount
		} else {
			currentAmount := uint8(0)
			_, _ = c.SetCash(allowanceCash, currentAmount)
			result[allowanceCash] = currentAmount
		}
	}
	return result
}
