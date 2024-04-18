package v1

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/ahmadmilzam/go/internal/usecase"
	"github.com/ahmadmilzam/go/pkg/httpres"
	"github.com/gin-gonic/gin"
)

type PingRoute struct {
	usecase usecase.AppUsecaseInterface
}

func NewAccountRoute(handler *gin.RouterGroup, u usecase.AppUsecaseInterface) {
	route := &PingRoute{u}
	h := handler.Group("/ping")
	{
		h.GET("/:id", route.pong)
	}
}

func (route *PingRoute) pong(ctx *gin.Context) {
	var params usecase.CreateAccountReqParams
	c := context.Background()

	if err := ctx.ShouldBindJSON(&params); err != nil {
		err = fmt.Errorf("%s: r.createAccount: %w", httpres.GenericBadRequest, err)
		msg := "Fail to parse request data"
		ctx.Set("msg", msg)
		ctx.Set("err", err)
		ctx.JSON(
			httpres.GetStatusCode(err),
			httpres.GenerateErrResponse(err, msg),
		)
		return
	}

	isValid, err := params.Validate()
	if !isValid {
		msg := "Invalid request data"
		ctx.Set("msg", msg)
		ctx.Set("err", err)
		ctx.JSON(
			httpres.GetStatusCode(err),
			httpres.GenerateErrResponse(err, msg),
		)
		return
	}

	aw, err := route.usecase.CreateAccount(c, params)

	if err != nil {
		msg := "Fail to create account"
		ctx.Set("msg", msg)
		ctx.Set("err", err)
		ctx.JSON(
			httpres.GetStatusCode(err),
			httpres.GenerateErrResponse(err, msg),
		)
		return
	}

	ctx.JSON(http.StatusCreated, httpres.GenerateOK(aw))
}

func (route *AccountRoute) getAccount(ctx *gin.Context) {
	var req usecase.GetAccountReqParams
	c := context.Background()

	if err := ctx.ShouldBindUri(&req); err != nil {
		er := errors.New("bad param phone")
		err := fmt.Errorf("%s: r.getAccount: %w", httpres.GenericBadRequest, er)
		msg := "Fail to parse request data"
		ctx.Set("msg", msg)
		ctx.Set("err", err)
		ctx.JSON(
			httpres.GetStatusCode(err),
			httpres.GenerateErrResponse(err, msg),
		)
		return
	}

	account, err := route.usecase.GetAccount(c, req.Phone)

	if err != nil {
		var msg string
		sc := httpres.GetStatusCode(err)
		if sc >= 500 {
			msg = "Internal server error"
		} else {
			msg = "Account not found"
		}

		ctx.Set("msg", msg)
		ctx.Set("err", err)
		ctx.JSON(
			sc,
			httpres.GenerateErrResponse(err, msg),
		)
		return
	}

	ctx.JSON(http.StatusOK, httpres.GenerateOK(account))
}
