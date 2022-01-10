package httpServer

import (
	"assignment/question_2/cashier"
	"github.com/labstack/echo/v4"
)

func Initialize() {
	e := echo.New()
	ch := cashier.NewCashier()
	h := Handler{}
	h.Initialize(&ch)
	e.GET("/get-all-cash-remaining", h.GetAllRemaining)
	e.PUT("/insert-cash", h.InsertCash)
	e.PUT("/set-cash", h.SetCash)
	e.PUT("/reduce-cash", h.ReduceCash)
	e.POST("/pay-by-cash", h.PayByCash)
	e.Logger.Fatal(e.Start(":1111"))
}
