package handlers

import (
	"github.com/victorskg/my-wallet/internal/usecases"
	"net/http"

	"github.com/google/uuid"
	"github.com/victorskg/my-wallet/pkg/http/response"
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
	importDividends usecases.ImportDividends
}

func NewImportDividendsHandler(importDividends usecases.ImportDividends) ImportDividendsHandler {
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

func (h ImportDividendsHandler) createResponseBody(output usecases.ImportDividendsOutput) responseBody {
	var rBody = responseBody{message: output.Message}
	switch output.Status {
	case usecases.Success:
		rBody.statusCode = http.StatusCreated
		break
	case usecases.PartialSuccess:
		rBody.statusCode = http.StatusPartialContent
		break
	case usecases.Error:
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
