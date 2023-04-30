package stock

import (
	"github.com/google/uuid"
	"github.com/victorskg/my-wallet/internal/domain/wallet"
	"github.com/victorskg/my-wallet/internal/gateways"
)

type OutputStatus string

const (
	Success        OutputStatus = "success"
	PartialSuccess OutputStatus = "partial_success"
	Error          OutputStatus = "error"
)

type OutputError struct {
	ticker string
	error  error
}

type Output struct {
	message string
	status  OutputStatus
	errors  []OutputError
}

type ImportDividends struct {
	stockGateway    gateways.StockGateway
	walletGateway   gateways.WalletGateway
	dividendGateway gateways.DividendGateway
}

func NewImportDividendsUseCase(
	stockGateway gateways.StockGateway,
	walletGateway gateways.WalletGateway,
	dividendGateway gateways.DividendGateway) ImportDividends {
	return ImportDividends{
		stockGateway:    stockGateway,
		walletGateway:   walletGateway,
		dividendGateway: dividendGateway,
	}
}

func (u ImportDividends) Execute(walletID uuid.UUID) Output {
	w, err := u.walletGateway.GetWallet(walletID)
	if err != nil {
		return Output{status: Error, message: err.Error()}
	}

	resultChannel := make(chan OutputError)
	for _, a := range w.Assets() {
		go func(asset wallet.Asset, c chan<- OutputError) {
			err = u.importDividendByAsset(asset)
			c <- OutputError{ticker: asset.Ticker(), error: err}
		}(a, resultChannel)
	}

	output := Output{status: Success, message: "Dividendos importados com sucesso!", errors: []OutputError{}}
	for range w.Assets() {
		result := <-resultChannel
		if result.error != nil {
			output.status = PartialSuccess
			output.message = "Não foi possível importar o dividendo de alguns ativos."
			output.errors = append(output.errors, result)
		}
	}

	if len(output.errors) == len(w.Assets()) {
		output.status = Error
		output.message = "Não foi possível importar os dividendos."
	}

	return output
}

func (u ImportDividends) importDividendByAsset(asset wallet.Asset) error {
	stockDomain, err := u.stockGateway.FindByTicker(asset.Ticker())
	if err != nil {
		return err
	}

	dividends, err := u.dividendGateway.ImportDividends(asset.Ticker())
	if err != nil {
		return err
	}

	stockDomain.AddDividends(dividends...)
	asset.AddDividends(dividends...)

	_, err = u.stockGateway.UpdateStock(stockDomain)
	if err != nil {
		return err
	}

	return nil
}
