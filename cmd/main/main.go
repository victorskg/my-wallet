package main

import (
	"github.com/victorskg/my-wallet/internal/usecases"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/victorskg/my-wallet/internal/adapters/controllers"
	stockController "github.com/victorskg/my-wallet/internal/adapters/controllers/stock"
	stockHandlers "github.com/victorskg/my-wallet/internal/adapters/controllers/stock/handlers"
	walletController "github.com/victorskg/my-wallet/internal/adapters/controllers/wallet"
	walletHandlers "github.com/victorskg/my-wallet/internal/adapters/controllers/wallet/handlers"
	"github.com/victorskg/my-wallet/internal/adapters/gateways/client"
	"github.com/victorskg/my-wallet/internal/adapters/gateways/database"
	"github.com/victorskg/my-wallet/internal/adapters/gateways/database/entity"
	"github.com/victorskg/my-wallet/internal/usecases/stock"
	dbRepo "github.com/victorskg/my-wallet/pkg/database"
)

const (
	databaseURL      = "localhost"
	databasePort     = "5432"
	databaseDriver   = "postgres"
	databaseName     = "my_wallet"
	databaseSchema   = "my_wallet"
	databaseUser     = "local_user"
	databasePassword = "local_pwd"
)

func main() {
	logger := log.Default()
	r := chi.NewRouter()
	logger.Println("Starting MY-WALLET")

	r.Use(middleware.Logger)

	r.Mount("/stocks", StockController().Routes())
	r.Mount("/wallet", WalletController().Routes())
	logger.Println("MY-WALLET listening on localhost:3333")

	err := http.ListenAndServe(":3333", r)
	if err != nil {
		log.Panicf("Error starting MY-WALLET. Caused by: %s", err.Error())
		return
	}
}

func StockController() controllers.Controller {
	stockGateway := database.NewStockDatabaseGateway()
	saveStock := stock.NewSaveStockUseCase(stockGateway)
	saveStocksHandler := stockHandlers.NewSaveStocksHandler(saveStock)
	return stockController.NewStockController(saveStocksHandler)
}

func WalletController() controllers.Controller {
	walletPsqlRepository := dbRepo.NewRepository(databaseURL, databasePort, databaseDriver, databaseName,
		databaseUser, databasePassword, "wallet", databaseSchema, entity.WalletEntity{})
	walletGateway := database.NewWalletDatabaseGateway(walletPsqlRepository)

	stockGateway := database.NewStockDatabaseGateway()
	dividendsGateway := client.NewDividendClientGateway()

	createWallet := usecases.NewCreateWalletUseCase(walletGateway)
	createWalletHandler := walletHandlers.NewCreateWalletHandler(createWallet)
	getWallet := usecases.NewGetWalletUseCase(walletGateway)
	getWalletHandler := walletHandlers.NewGetWalletHandler(getWallet)
	makeInvestment := usecases.NewMakeInvestmentUseCase(walletGateway, stockGateway)
	makeInvestmentHandler := walletHandlers.NewMakeInvestmentHandler(makeInvestment)
	importDividends := usecases.NewImportDividendsUseCase(stockGateway, walletGateway, dividendsGateway)
	importDividendsHandler := walletHandlers.NewImportDividendsHandler(importDividends)
	return walletController.NewWalletController(getWalletHandler, createWalletHandler, makeInvestmentHandler, importDividendsHandler)
}
