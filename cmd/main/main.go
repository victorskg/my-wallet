package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/victorskg/my-wallet/internal/adapters/controllers"
	stockController "github.com/victorskg/my-wallet/internal/adapters/controllers/stock"
	stockHandlers "github.com/victorskg/my-wallet/internal/adapters/controllers/stock/handlers"
	walletController "github.com/victorskg/my-wallet/internal/adapters/controllers/wallet"
	walletHandlers "github.com/victorskg/my-wallet/internal/adapters/controllers/wallet/handlers"
	"github.com/victorskg/my-wallet/internal/adapters/database"
	"github.com/victorskg/my-wallet/internal/usecases/stock"
	"github.com/victorskg/my-wallet/internal/usecases/wallet"
)

func main() {
	logger := log.Default()
	r := chi.NewRouter()
	logger.Println("Starting MY-WALLET")

	r.Use(middleware.Logger)

	r.Mount("/stocks", StockController().Routes())
	r.Mount("/wallet", WalletController().Routes())
	logger.Println("MY-WALLET listening on localhost:3333")

	http.ListenAndServe(":3333", r)
}

func StockController() controllers.Controller {
	stockGateway := database.NewStockDatabaseGateway()
	saveStock := stock.NewSaveStockUseCase(stockGateway)
	saveStocksHandler := stockHandlers.NewSaveStocksHandler(saveStock)
	return stockController.NewStockController(saveStocksHandler)
}

func WalletController() controllers.Controller {
	walletGateway := database.NewWalletDatabaseGateway()
	stockGateway := database.NewStockDatabaseGateway()
	createWallet := wallet.NewCreateWalletUseCase(walletGateway)
	createWalletHandler := walletHandlers.NewCreateWalletHandler(createWallet)
	makeInvestment := wallet.NewMakeInvestmentUseCase(walletGateway, stockGateway)
	makeInvestmentHandler := walletHandlers.NewMakeInvestmentHandler(makeInvestment)
	return walletController.NewWalletController(createWalletHandler, makeInvestmentHandler)
}
