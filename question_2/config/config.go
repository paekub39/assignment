package config

import (
	"assignment/question_2/storage"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

const (
	storageFileName = "./storage.csv"
)

func WriteTemp(s *storage.Storage) {
	recordFile, err := os.Create(storageFileName)
	if err != nil {
		fmt.Println("An error encountered ::", err)
	}

	writer := csv.NewWriter(recordFile)
	var csvData = [][]string{
		{strconv.FormatFloat(float64(storage.OneThousandCash), 'E', -1, 32), strconv.FormatFloat(float64(storage.FiveHundredCash), 'E', -1, 32), strconv.FormatFloat(float64(storage.OneHundredCash), 'E', -1, 32), strconv.FormatFloat(float64(storage.FiftyCash), 'E', -1, 32), strconv.FormatFloat(float64(storage.TwentyCash), 'E', -1, 32), strconv.FormatFloat(float64(storage.TenCash), 'E', -1, 32), strconv.FormatFloat(float64(storage.FiveCash), 'E', -1, 32), strconv.FormatFloat(float64(storage.OneCash), 'E', -1, 32), strconv.FormatFloat(float64(storage.PointTwoFiveCash), 'E', -1, 32)},
		{strconv.FormatFloat(float64(s.CashList[storage.OneThousandCash].CurrentAmount), 'E', -1, 32), strconv.FormatFloat(float64(s.CashList[storage.FiveHundredCash].CurrentAmount), 'E', -1, 32), strconv.FormatFloat(float64(s.CashList[storage.OneHundredCash].CurrentAmount), 'E', -1, 32), strconv.FormatFloat(float64(s.CashList[storage.FiftyCash].CurrentAmount), 'E', -1, 32), strconv.FormatFloat(float64(s.CashList[storage.TwentyCash].CurrentAmount), 'E', -1, 32), strconv.FormatFloat(float64(s.CashList[storage.TenCash].CurrentAmount), 'E', -1, 32), strconv.FormatFloat(float64(s.CashList[storage.FiveCash].CurrentAmount), 'E', -1, 32), strconv.FormatFloat(float64(s.CashList[storage.OneCash].CurrentAmount), 'E', -1, 32), strconv.FormatFloat(float64(s.CashList[storage.PointTwoFiveCash].CurrentAmount), 'E', -1, 32)},
	}
	err = writer.WriteAll(csvData) // returns error
	if err != nil {
		fmt.Println("An error encountered ::", err)
	}
}

func ReadTemp() (*storage.Storage, error) {
	recordFile, err := os.Open(storageFileName)
	if err != nil {
		fmt.Println("An error encountered ::", err)
		return nil, err
	}
	csvReader := csv.NewReader(recordFile)
	data, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("An error encountered ::", err)
		return nil, err
	}
	st := new(storage.Storage)
	for i, line := range data {
		if i > 0 {
			cashList := make(map[float32]storage.CashTemplate)
			for j, field := range line {
				value, err := strconv.ParseFloat(field, 32)
				if err != nil {
					value = 0
				}
				val := uint8(value)

				if j == 0 {
					newCash := storage.CashTemplate{
						Limit:         storage.OneThousandLimit,
						CurrentAmount: val,
					}
					cashList[storage.OneThousandCash] = newCash
				} else if j == 1 {
					newCash := storage.CashTemplate{
						Limit:         storage.FiveHundredLimit,
						CurrentAmount: val,
					}
					cashList[storage.FiveHundredCash] = newCash
				} else if j == 2 {
					newCash := storage.CashTemplate{
						Limit:         storage.OneHundredLimit,
						CurrentAmount: val,
					}
					cashList[storage.OneHundredCash] = newCash
				} else if j == 3 {
					newCash := storage.CashTemplate{
						Limit:         storage.FiftyLimit,
						CurrentAmount: val,
					}
					cashList[storage.FiftyCash] = newCash
				} else if j == 4 {
					newCash := storage.CashTemplate{
						Limit:         storage.TwentyLimit,
						CurrentAmount: val,
					}
					cashList[storage.TwentyCash] = newCash
				} else if j == 5 {
					newCash := storage.CashTemplate{
						Limit:         storage.TenLimit,
						CurrentAmount: val,
					}
					cashList[storage.TenCash] = newCash
				} else if j == 6 {
					newCash := storage.CashTemplate{
						Limit:         storage.FiveLimit,
						CurrentAmount: val,
					}
					cashList[storage.FiveCash] = newCash
				} else if j == 7 {
					newCash := storage.CashTemplate{
						Limit:         storage.OneLimit,
						CurrentAmount: val,
					}
					cashList[storage.OneCash] = newCash
				} else if j == 8 {
					newCash := storage.CashTemplate{
						Limit:         storage.PointTwoFiveLimit,
						CurrentAmount: val,
					}
					cashList[storage.PointTwoFiveCash] = newCash
				}
			}
			st.CashList = cashList
			break
		}
	}
	return st, nil
}
