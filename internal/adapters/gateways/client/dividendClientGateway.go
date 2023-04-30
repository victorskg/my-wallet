package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/victorskg/my-wallet/internal/domain/stock"
	"github.com/victorskg/my-wallet/internal/gateways"
)

type Response struct {
	Results []struct {
		DividendsData struct {
			CashDividends []struct {
				AssetIssued   string    `json:"assetIssued"`
				PaymentDate   time.Time `json:"paymentDate"`
				Rate          float64   `json:"rate"`
				RelatedTo     string    `json:"relatedTo"`
				ApprovedOn    time.Time `json:"approvedOn"`
				IsinCode      string    `json:"isinCode"`
				Label         string    `json:"label"`
				LastDatePrior time.Time `json:"lastDatePrior"`
				Remarks       string    `json:"remarks"`
			} `json:"cashDividends"`
		} `json:"dividendsData"`
	} `json:"results"`
	RequestedAt time.Time `json:"requestedAt"`
}

type DividendClientGateway struct {
}

func NewDividendClientGateway() gateways.DividendGateway {
	return DividendClientGateway{}
}

func (d DividendClientGateway) ImportDividends(ticker string) ([]stock.Dividend, error) {
	resp, err := http.Get(fmt.Sprintf("https://brapi.dev/api/quote/%s?fundamental=false&dividends=true", ticker))
	if err != nil {
		return nil, fmt.Errorf("Erro ao importar dividendos de API. Causado por: %s.", err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Erro ao ler bytes de resposta da API. Causado por: %s.", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println(fmt.Printf("Error importing dividends from API. Status %d. Body %s", resp.StatusCode, string(body)))
		return nil, fmt.Errorf("Erro %d ao importar dividendos de API.", resp.StatusCode)
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("Erro ao ler JSON de resposta da API. Causado por: %s.", err.Error())
	}

	var dividends []stock.Dividend
	for _, dividend := range response.Results[0].DividendsData.CashDividends {
		var dType stock.DType
		if dividend.Label == "RENDIMENTO" {
			dType = stock.Income
		} else if dividend.Label == "AMORTIZACAO RF" {
			dType = stock.Amortization
		}

		stockDividend := stock.NewDividend(dividend.Rate, dividend.ApprovedOn, dividend.PaymentDate, dType)
		dividends = append(dividends, *stockDividend)
	}

	return dividends, nil
}
