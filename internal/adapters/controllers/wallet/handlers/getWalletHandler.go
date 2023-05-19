package handlers

import (
	"github.com/google/uuid"
	"github.com/victorskg/my-wallet/internal/usecases/wallet"
	domainErrors "github.com/victorskg/my-wallet/pkg/error"
	"github.com/victorskg/my-wallet/pkg/http/response"
	"net/http"
)

type Output struct {
	Id          uuid.UUID `json:"id"`
	Description string    `json:"description"`
}

type GetWalletHandler struct {
	getWallet wallet.GetWallet
}

func NewGetWalletHandler(getWallet wallet.GetWallet) GetWalletHandler {
	return GetWalletHandler{
		getWallet: getWallet,
	}
}

func (h GetWalletHandler) GetWallet(w http.ResponseWriter, r *http.Request) {
	walletID := r.Context().Value("walletID").(uuid.UUID)
	wDomain, err := h.getWallet.Execute(walletID)
	if err != nil {
		responseStatus := http.StatusBadRequest
		_, isNotFoundErr := err.(domainErrors.NotFound)

		if isNotFoundErr {
			responseStatus = http.StatusNotFound
		}

		response.WriteResponseMessage(w, err.Error(), responseStatus)
		return
	}

	response.WriteJSONResponse(w, Output{
		Id:          wDomain.Id(),
		Description: wDomain.Description(),
	}, http.StatusOK)
}
