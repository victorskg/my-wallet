package handlers

import (
	"encoding/csv"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/victorskg/my-wallet/pkg/http/response"

	"github.com/victorskg/my-wallet/internal/usecases/wallet"
)

const (
	inputFileName   = "investments"
	fileContentType = "multipart/form-data"
	jsonContentType = "application/json"
)

var expectedColumnsInOrder = []string{"Entrada/Saída", "Data", "Movimentação", "Produto", "Instituição",
	"Quantidade", "Preço unitário", "Valor da Operação"}

type MakeInvestmentHandler struct {
	makeInvestment wallet.MakeInvestment
}

func NewMakeInvestmentHandler(makeInvestment wallet.MakeInvestment) MakeInvestmentHandler {
	return MakeInvestmentHandler{
		makeInvestment: makeInvestment,
	}
}

func (h MakeInvestmentHandler) MakeInvestment(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.Header.Get("Content-Type"), fileContentType) {
		h.processFileContent(w, r)
		return
	}

	if strings.Contains(r.Header.Get("Content-Type"), jsonContentType) {
		h.processJsonContent(w, r)
		return
	}

	response.WriteResponseMessage(w, "Requisição fora do formato esperado. Envie um arquivo csv ou faça um lançamento manual",
		http.StatusUnsupportedMediaType)
	return
}

func (h MakeInvestmentHandler) processFileContent(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile(inputFileName)
	if err != nil {
		response.WriteResponseMessage(w, "O arquivo csv de investimentos é obrigatório.",
			http.StatusBadRequest)
		return
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		response.WriteResponseMessage(w, "Não foi possível ler o arquivo enviado.",
			http.StatusBadRequest)
		return
	}

	walletID, err := uuid.Parse(chi.URLParam(r, "walletID"))
	if err != nil {
		response.WriteResponseMessage(w, "O ID enviado não corresponde a um ID válido.",
			http.StatusBadRequest)
		return
	}

	investmentInput, err := h.createInvestmentInputFromFileData(walletID, data)
	if err != nil {
		response.WriteResponseMessage(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.makeInvestment.Execute(*investmentInput)
	if err != nil {
		response.WriteResponseMessage(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.WriteResponseMessage(w, "Investimentos lançados com sucesso.", http.StatusCreated)
}
func (h MakeInvestmentHandler) createInvestmentInputFromFileData(walletID uuid.UUID, data [][]string) (*wallet.MakeInvestmentInput, error) {
	actualHeader := strings.Join(data[0], ",")
	expectedHeader := strings.Join(expectedColumnsInOrder, ",")
	if actualHeader != expectedHeader {
		return nil, fmt.Errorf("As colunas enviadas estão incorretas. Por favor, envie as colunas na ordem: %s",
			expectedHeader)
	}

	var assets []wallet.InvestmentAsset
	for _, line := range data[1:] {
		if line[0] == "Credito" && line[2] == "Transferência - Liquidação" {
			asset, err := h.createInvestmentAssetFromFileLine(line)
			if err != nil {
				return nil, err
			}

			assets = append(assets, *asset)
		}
	}

	return &wallet.MakeInvestmentInput{
		WalletID: walletID,
		Assets:   assets,
	}, nil
}

func (h MakeInvestmentHandler) createInvestmentAssetFromFileLine(line []string) (*wallet.InvestmentAsset, error) {
	date, err := time.Parse("02/01/2006", line[1])
	if err != nil {
		return nil, errors.New("A data enviada não corresponde a uma data válida. Envia a data no formato DD/MM/YYYY.")
	}

	ticker := strings.TrimSpace(strings.Split(line[3], "-")[0])

	quotas, err := strconv.ParseInt(line[5], 10, 16)
	if err != nil {
		return nil, fmt.Errorf("A quantidade de cotas enviada para o ativo %s na data %v é inválida.", ticker, date)
	}

	price, err := strconv.ParseFloat(strings.TrimPrefix(strings.ReplaceAll(line[6], " ", ""), "R$"), 64)
	if err != nil {
		return nil, fmt.Errorf("O preço da cota enviado para o ativo %s na data %v é inválida.", ticker, date)
	}

	return &wallet.InvestmentAsset{
		Ticker:  ticker,
		BuyDate: date,
		Price:   price,
		Quotas:  uint32(quotas),
	}, nil
}

func (h MakeInvestmentHandler) processJsonContent(w http.ResponseWriter, r *http.Request) {
	//TODO Implement json handler
}
