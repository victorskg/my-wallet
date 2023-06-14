package handlers

import (
	"github.com/victorskg/my-wallet/internal/usecases"
	"github.com/victorskg/my-wallet/pkg/http/response"
	"net/http"

	"github.com/victorskg/my-wallet/pkg/http/json"
)

type Input struct {
	Description string `json:"description"`
}

type CreateWalletHandler struct {
	createWallet usecases.CreateWallet
}

func NewCreateWalletHandler(createWallet usecases.CreateWallet) CreateWalletHandler {
	return CreateWalletHandler{
		createWallet: createWallet,
	}
}

func (h CreateWalletHandler) CreateWallet(w http.ResponseWriter, r *http.Request) {
	var data Input
	if err := json.Deserialize[Input](&data, w, r); err != nil {
		return
	}

	wDomain, err := h.createWallet.Execute(data.Description)
	if err != nil {
		response.WriteResponseMessage(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("X-Wallet-ID", wDomain.Id().String())
	w.WriteHeader(http.StatusCreated)
}
