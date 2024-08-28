package pkg

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Mr-LvGJ/cloud-native-poc/pkg/common/setting"
)

func InitNecessity(ctx context.Context) {
}

func RunServer(_ context.Context) {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("Wellcome to %s!", setting.GlobalConfig().ServiceName))
	})
	e.Logger.Fatal(e.Start(":18080"))
}
