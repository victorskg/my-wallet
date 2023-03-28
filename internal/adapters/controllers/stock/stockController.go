package stock

import (
	"github.com/go-chi/chi/v5"
	"github.com/victorskg/my-wallet/internal/adapters/controllers"
	"github.com/victorskg/my-wallet/internal/adapters/controllers/stock/handlers"
)

type Controller struct {
	saveStocksHandler handlers.SaveStocksHandler
}

func NewStockController(saveStocksHandler handlers.SaveStocksHandler) controllers.Controller {
	return Controller{
		saveStocksHandler: saveStocksHandler,
	}
}

func (c Controller) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", c.saveStocksHandler.SaveStocks)

	return r
}
