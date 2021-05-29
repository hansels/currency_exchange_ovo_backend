package api

import (
	myRouter "github.com/hansels/currency_exchange_ovo_backend/common/router"
	"github.com/hansels/currency_exchange_ovo_backend/src/usecase"
	"github.com/julienschmidt/httprouter"
)

func (a *API) Register(router *httprouter.Router) {
	router.GET("/ping", myRouter.HandleNow("/ping", a.Ping))
	router.GET("/currencies", myRouter.HandleNow("/currencies", a.Currencies))

	router.POST("/currency", myRouter.HandleNow("/currency", a.Currency))
	router.POST("/calculate", myRouter.HandleNow("/calculate", a.Calculate))
}

type API struct {
	Module *usecase.Module
}

func New(module *usecase.Module) *API {
	return &API{Module: module}
}
