package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/victorskg/my-wallet/pkg/http/response"

	"github.com/victorskg/my-wallet/internal/usecases/wallet"
)

type responseBody struct {
	statusCode int `json:"-"`
	message    string
	details    []responseBodyDetails
}

type responseBodyDetails struct {
	ticker string
	error  string
}

type ImportDividendsHandler struct {
	importDividends wallet.ImportDividends
}

func NewImportDividendsHandler(importDividends wallet.ImportDividends) ImportDividendsHandler {
	return ImportDividendsHandler{
		importDividends: importDividends,
	}
}

func (h ImportDividendsHandler) ImportDividends(w http.ResponseWriter, r *http.Request) {
	walletID := r.Context().Value("walletID").(uuid.UUID)
	output := h.importDividends.Execute(walletID)
	rBody := h.createResponseBody(output)

	response.WriteJSONResponse(w, rBody, rBody.statusCode)
}

func (h ImportDividendsHandler) createResponseBody(output wallet.ImportDividendsOutput) responseBody {
	var rBody = responseBody{message: output.Message}
	switch output.Status {
	case wallet.Success:
		rBody.statusCode = http.StatusCreated
		break
	case wallet.PartialSuccess:
		rBody.statusCode = http.StatusPartialContent
		break
	case wallet.Error:
		rBody.statusCode = http.StatusBadRequest
		break
	}

	if rBody.statusCode != http.StatusCreated {
		for _, outputError := range output.Errors {
			rBody.details = append(rBody.details, responseBodyDetails{
				ticker: outputError.Ticker,
				error:  outputError.Error.Error(),
			})
		}
	}

	return rBody
}
