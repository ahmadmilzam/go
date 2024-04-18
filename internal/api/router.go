package api

import (
	"errors"
	"net/http"

	v1 "github.com/ahmadmilzam/go/internal/api/v1"
	"github.com/ahmadmilzam/go/internal/usecase"
	"github.com/ahmadmilzam/go/pkg/httpres"
	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine, u usecase.AppUsecaseInterface) {
	router.HandleMethodNotAllowed = true
	// K8s probe for kubernetes health checks -.
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "The server is up and running")
	})

	// Handling a page not found endpoint -.
	router.NoRoute(func(c *gin.Context) {
		c.JSON(
			http.StatusNotFound,
			httpres.GenerateErrResponse(errors.New("40400: Endpoint not found"), "Endpoint not found"),
		)
	})

	// Routers -.
	rgroupv1 := router.Group("/v1")
	{
		v1.NewAccountRoute(rgroupv1, u)
		v1.NewWalletRoute(rgroupv1, u)
		v1.NewTransferRoute(rgroupv1, u)
	}
}
