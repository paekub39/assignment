package httpServer

import (
	"assignment/question_1/core"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
}

func (h *Handler) GetXYZByPosition(c echo.Context) error {
	resp := core.GetXYZByPosition()
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetXYZByRemoveKnowData(c echo.Context) error {
	resp := core.GetXYZByRemoveKnowData()
	return c.JSON(http.StatusOK, resp)
}
