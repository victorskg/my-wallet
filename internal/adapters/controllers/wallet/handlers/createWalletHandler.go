package handlers

import (
	"github.com/victorskg/my-wallet/pkg/http/response"
	"net/http"

	"github.com/victorskg/my-wallet/pkg/http/json"

	"github.com/victorskg/my-wallet/internal/usecases/wallet"
)

type Input struct {
	Description string `json:"description"`
}

type CreateWalletHandler struct {
	createWallet wallet.CreateWallet
}

func NewCreateWalletHandler(createWallet wallet.CreateWallet) CreateWalletHandler {
	return CreateWalletHandler{
		createWallet: createWallet,
	}
}

func (h CreateWalletHandler) CreateWallet(w http.ResponseWriter, r *http.Request) {
	var data Input
	if err := json.Deserialize[Input](&data, w, r); err != nil {
		return
	}

	wallet, err := h.createWallet.Execute(data.Description)
	if err != nil {
		response.WriteResponseMessage(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("X-Wallet-ID", wallet.Id().String())
	w.WriteHeader(http.StatusCreated)
}
