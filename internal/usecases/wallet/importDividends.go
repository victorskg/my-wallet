package wallet

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

type ImportDividendsOutputError struct {
	Ticker string
	Error  error
}

type ImportDividendsOutput struct {
	Message string
	Status  OutputStatus
	Errors  []ImportDividendsOutputError
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

func (u ImportDividends) Execute(walletID uuid.UUID) ImportDividendsOutput {
	w, err := u.walletGateway.GetWallet(walletID)
	if err != nil {
		return ImportDividendsOutput{Status: Error, Message: err.Error()}
	}

	resultChannel := make(chan ImportDividendsOutputError)
	for _, a := range w.Assets() {
		go func(asset wallet.Asset, c chan<- ImportDividendsOutputError) {
			err = u.importDividendByAsset(asset)
			c <- ImportDividendsOutputError{Ticker: asset.Ticker(), Error: err}
		}(a, resultChannel)
	}

	output := ImportDividendsOutput{Status: Success, Message: "Dividendos importados com sucesso!", Errors: []ImportDividendsOutputError{}}
	for range w.Assets() {
		result := <-resultChannel
		if result.Error != nil {
			output.Status = PartialSuccess
			output.Message = "Não foi possível importar o dividendo de alguns ativos."
			output.Errors = append(output.Errors, result)
		}
	}

	if len(output.Errors) == len(w.Assets()) {
		output.Status = Error
		output.Message = "Não foi possível importar os dividendos."
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
