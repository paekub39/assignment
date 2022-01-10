package main

import (
	"assignment/question_2/storage"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	host := "http://localhost:1111"
	fmt.Println("Get All Cash Remaining")
	getAllCashRemaining(&host)
	fmt.Println("")
	fmt.Println("Insert All Cash By One")
	insertAllCashByOne(&host)
	fmt.Println("")
	fmt.Println("Get All Cash Remaining")
	getAllCashRemaining(&host)
	fmt.Println("")
	fmt.Println("Set All Cash To Max Amount")
	setAllCashToMaxAmount(&host)
	fmt.Println("")
	fmt.Println("Get All Cash Remaining")
	getAllCashRemaining(&host)
	fmt.Println("")
	fmt.Println("Reduce All Cash By Ten")
	reduceAllCashByOne(&host)
	fmt.Println("")
	fmt.Println("Get All Cash Remaining")
	getAllCashRemaining(&host)
	fmt.Println("")
	fmt.Println("Purchase Item: Price 3200, pay with 1000 * 4")
	paidWithCash(&host)
	fmt.Println("")
	fmt.Println("Get All Cash Remaining")
	getAllCashRemaining(&host)
}

func getAllCashRemaining(host *string) {
	fullPath := fmt.Sprintf("%s/get-all-cash-remaining", *host)
	resp, err := http.Get(fullPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	result := string(b)
	fmt.Printf("%s", result)
}

type InsertCashInput struct {
	CashList []InsertCashData `json:"cashList"`
}

type InsertCashData struct {
	Value  float32 `json:"value"`
	Amount uint8   `json:"amount"`
}

func insertAllCashByOne(host *string) {
	fullPath := fmt.Sprintf("%s/insert-cash", *host)
	cashList := storage.GetAllowCash()
	scdList := new([]InsertCashData)
	input := new(InsertCashInput)
	for _, cash := range cashList {
		scd := InsertCashData{
			Value:  cash,
			Amount: 1,
		}
		*scdList = append(*scdList, scd)
	}
	input.CashList = *scdList
	j, err := json.Marshal(input)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(http.MethodPut, fullPath, bytes.NewBuffer(j))
	if err != nil {
		panic(err)
	}
	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	// initialize http client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	result := string(b)
	fmt.Printf("%s\n", result)
}

type SetCashInput struct {
	CashList []SetCashData `json:"cashList"`
}

type SetCashData struct {
	Value  float32 `json:"value"`
	Amount uint8   `json:"amount"`
}

func setAllCashToMaxAmount(host *string) {
	fullPath := fmt.Sprintf("%s/set-cash", *host)
	cashList := storage.GetAllowCash()
	scdList := new([]SetCashData)
	input := new(SetCashInput)
	for _, cash := range cashList {
		lim, _ := storage.GetCashLimit(cash)
		scd := SetCashData{
			Value:  cash,
			Amount: lim,
		}
		*scdList = append(*scdList, scd)
	}
	input.CashList = *scdList
	j, err := json.Marshal(input)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(http.MethodPut, fullPath, bytes.NewBuffer(j))
	if err != nil {
		panic(err)
	}
	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	// initialize http client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	result := string(b)
	fmt.Printf("%s\n", result)
}

type ReduceCashInput struct {
	CashList []ReduceCashData `json:"cashList"`
}

type ReduceCashData struct {
	Value  float32 `json:"value"`
	Amount uint8   `json:"amount"`
}

func reduceAllCashByOne(host *string) {
	fullPath := fmt.Sprintf("%s/reduce-cash", *host)
	cashList := storage.GetAllowCash()
	scdList := new([]ReduceCashData)
	input := new(ReduceCashInput)
	for _, cash := range cashList {
		scd := ReduceCashData{
			Value:  cash,
			Amount: 10,
		}
		*scdList = append(*scdList, scd)
	}
	input.CashList = *scdList
	j, err := json.Marshal(input)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(http.MethodPut, fullPath, bytes.NewBuffer(j))
	if err != nil {
		panic(err)
	}

	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	// initialize http client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	result := string(b)
	fmt.Printf("%s\n", result)
}

type PaidInput struct {
	Paid  []PaidData `json:"paid"`
	Price float32    `json:"price"`
}

type PaidData struct {
	Value  float32 `json:"value"`
	Amount uint8   `json:"amount"`
}

func paidWithCash(host *string) {
	fullPath := fmt.Sprintf("%s/pay-by-cash", *host)
	price := float32(3200)
	pdList := new([]PaidData)
	pd1 := PaidData{
		Value:  1000,
		Amount: 4,
	}
	*pdList = append(*pdList, pd1)
	input := PaidInput{
		Paid:  *pdList,
		Price: price,
	}

	j, err := json.Marshal(input)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(http.MethodPost, fullPath, bytes.NewBuffer(j))
	if err != nil {
		panic(err)
	}

	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	// initialize http client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	result := string(b)
	fmt.Printf("%s\n", result)
}
