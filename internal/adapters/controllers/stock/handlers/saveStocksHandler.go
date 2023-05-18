package handlers

import (
	"net/http"

	"github.com/victorskg/my-wallet/internal/domain/stock"
	usecases "github.com/victorskg/my-wallet/internal/usecases/stock"
	"github.com/victorskg/my-wallet/pkg/http/json"
	"github.com/victorskg/my-wallet/pkg/http/response"
)

type input struct {
	Stocks []struct {
		Ticker        string      `json:"ticker"`
		Name          string      `json:"name"`
		Type          stock.SType `json:"type"`
		Category      string      `json:"category"`
		SubCategory   string      `json:"sub_category"`
		Administrator string      `json:"administrator"`
		BookValue     float32     `json:"book_value"`
		Patrimony     float64     `json:"patrimony"`
		PVP           float32     `json:"pvp"`
	} `json:"stocks"`
}

type SaveStocksHandler struct {
	saveStocks usecases.SaveStock
}

func NewSaveStocksHandler(saveStocks usecases.SaveStock) SaveStocksHandler {
	return SaveStocksHandler{
		saveStocks: saveStocks,
	}
}

func (h SaveStocksHandler) SaveStocks(w http.ResponseWriter, r *http.Request) {
	var data input
	if err := json.Deserialize[input](&data, w, r); err != nil {
		return
	}

	var stocks []stock.Stock
	for _, s := range data.Stocks {
		stockDomain := stock.NewStock(s.Ticker, s.Name, s.Type, s.Category, s.SubCategory,
			s.Administrator, s.BookValue, s.Patrimony, s.PVP)
		stocks = append(stocks, *stockDomain)
	}

	if _, err := h.saveStocks.Execute(stocks...); err != nil {
		response.WriteResponseMessage(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
