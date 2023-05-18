package wallet

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/victorskg/my-wallet/internal/adapters/controllers"
	"github.com/victorskg/my-wallet/internal/adapters/controllers/wallet/handlers"
	"github.com/victorskg/my-wallet/pkg/http/response"
)

type Controller struct {
	getWalletHandler       handlers.GetWalletHandler
	createWalletHandler    handlers.CreateWalletHandler
	makeInvestmentHandler  handlers.MakeInvestmentHandler
	importDividendsHandler handlers.ImportDividendsHandler
}

func NewWalletController(
	getWalletHandler handlers.GetWalletHandler,
	createWalletHandler handlers.CreateWalletHandler,
	makeInvestmentHandler handlers.MakeInvestmentHandler,
	importDividendsHandler handlers.ImportDividendsHandler) controllers.Controller {
	return Controller{
		getWalletHandler:       getWalletHandler,
		createWalletHandler:    createWalletHandler,
		makeInvestmentHandler:  makeInvestmentHandler,
		importDividendsHandler: importDividendsHandler,
	}
}

func (c Controller) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", c.createWalletHandler.CreateWallet)
	r.Route("/{walletID}", func(r chi.Router) {
		r.Use(WithIDContext)
		r.Get("/", c.getWalletHandler.GetWallet)
		r.Route("/stocks", func(r chi.Router) {
			r.Post("/buy", c.makeInvestmentHandler.MakeInvestment)
		})
		r.Post("/import-dividends", c.importDividendsHandler.ImportDividends)
	})

	return r
}

func WithIDContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		walletID, err := uuid.Parse(chi.URLParam(r, "walletID"))
		if err != nil {
			response.WriteResponseMessage(w, "O ID enviado não corresponde a um ID válido.",
				http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), "walletID", walletID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
