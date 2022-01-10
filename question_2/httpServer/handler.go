package httpServer

import (
	"assignment/question_2/cashier"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	Cashier *cashier.Cashier
}

func (h *Handler) Initialize(ch *cashier.Cashier) {
	h.Cashier = ch
}

func (h *Handler) GetAllRemaining(c echo.Context) error {
	ch := h.Cashier
	resp := getAllCashRemaining(ch)
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) InsertCash(c echo.Context) error {
	ch := h.Cashier
	input := new(InsertCashInput)

	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp, chNew := insertCash(ch, input)
	h.Cashier = &chNew

	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) SetCash(c echo.Context) error {
	ch := h.Cashier
	input := new(SetCashInput)

	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp, chNew := setCash(ch, input)
	h.Cashier = &chNew
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) ReduceCash(c echo.Context) error {
	ch := h.Cashier
	input := new(ReduceCashInput)

	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp, chNew := reduceCash(ch, input)
	h.Cashier = &chNew
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) PayByCash(c echo.Context) error {
	ch := h.Cashier
	input := new(PaidInput)

	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp, chNew := paidCash(ch, input)
	h.Cashier = &chNew
	return c.JSON(http.StatusOK, resp)
}
