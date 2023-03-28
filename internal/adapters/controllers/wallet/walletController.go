package wallet

import (
	"github.com/go-chi/chi/v5"
	"github.com/victorskg/my-wallet/internal/adapters/controllers"
	"github.com/victorskg/my-wallet/internal/adapters/controllers/wallet/handlers"
)

type Controller struct {
	createWalletHandler handlers.CreateWalletHandler
}

func NewWalletController(createWalletHandler handlers.CreateWalletHandler) controllers.Controller {
	return Controller{
		createWalletHandler: createWalletHandler,
	}
}

func (c Controller) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", c.createWalletHandler.CreateWallet)

	return r
}
