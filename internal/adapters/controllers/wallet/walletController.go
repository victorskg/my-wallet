package wallet

import (
	"github.com/go-chi/chi/v5"
	"github.com/victorskg/my-wallet/internal/adapters/controllers"
	"github.com/victorskg/my-wallet/internal/adapters/controllers/wallet/handlers"
)

type Controller struct {
	createWalletHandler   handlers.CreateWalletHandler
	makeInvestmentHandler handlers.MakeInvestmentHandler
}

func NewWalletController(createWalletHandler handlers.CreateWalletHandler,
	makeInvestmentHandler handlers.MakeInvestmentHandler) controllers.Controller {
	return Controller{
		createWalletHandler:   createWalletHandler,
		makeInvestmentHandler: makeInvestmentHandler,
	}
}

func (c Controller) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", c.createWalletHandler.CreateWallet)
	r.Route("/{walletID}", func(r chi.Router) {
		r.Route("/stocks", func(r chi.Router) {
			r.Post("/buy", c.makeInvestmentHandler.MakeInvestment)
		})
	})

	return r
}
