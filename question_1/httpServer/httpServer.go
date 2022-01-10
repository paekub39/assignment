package httpServer

import (
	"github.com/labstack/echo/v4"
)

func Initialize() {
	e := echo.New()
	h := Handler{}
	e.GET("/get-xyz-by-position", h.GetXYZByPosition)
	e.GET("/get-xyz-by-remove-know-data", h.GetXYZByRemoveKnowData)
	e.Logger.Fatal(e.Start(":1110"))
}
