package handlers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/victorskg/my-wallet/pkg/http/response"

	domain "github.com/victorskg/my-wallet/internal/domain/stock"
	usecases "github.com/victorskg/my-wallet/internal/usecases/stock"
)

const (
	inputFileName   = "stocks"
	fileContentType = "multipart/form-data"
)

var expectedColumnsInOrder = []string{"Ticker", "Type", "Category:subcategory",
	"Administrator", "Book Value", "Patrimony", "P/VP"}

type SaveStocksHandler struct {
	saveStocks usecases.SaveStock
}

func NewSaveStocksHandler(saveStocks usecases.SaveStock) SaveStocksHandler {
	return SaveStocksHandler{
		saveStocks: saveStocks,
	}
}

func (h SaveStocksHandler) SaveStocks(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.Header.Get("Content-Type"), fileContentType) {
		response.WriteResponseMessage(w, "Formato do arquivo não suportado. Formato aceito: csv.",
			http.StatusUnsupportedMediaType)
		return
	}

	file, _, err := r.FormFile(inputFileName)
	if err != nil {
		response.WriteResponseMessage(w, "O arquivo csv de ações é obrigatório.",
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

	stocks, err := h.createStocksFromFileData(data)
	if err != nil {
		response.WriteResponseMessage(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.saveStocks.Execute(stocks...)
	if err != nil {
		response.WriteResponseMessage(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.WriteResponseMessage(w, "Ações criadas com sucesso.", http.StatusCreated)
}

func (h SaveStocksHandler) createStocksFromFileData(data [][]string) ([]domain.Stock, error) {
	actualHeader := strings.Join(data[0], ",")
	expectedHeader := strings.Join(expectedColumnsInOrder, ",")
	if actualHeader != expectedHeader {
		return nil, fmt.Errorf("As colunas enviadas estão incorretas. Por favor, envie as colunas na ordem: %s",
			expectedHeader)
	}

	var stocks []domain.Stock
	for _, line := range data[1:] {
		stocks = append(stocks, h.createStockFromFileLine(line))
	}
	return stocks, nil
}

func (h SaveStocksHandler) createStockFromFileLine(line []string) domain.Stock {
	ticker := line[0]
	sType := line[1]
	category, subCategory := h.extractCategoryAndSubcategory(line[2])
	administrator := line[3]
	bookValue, _ := h.parseNumber(line[4])
	patrimony, _ := h.parseNumber(line[5])
	pvp, _ := h.parseNumber(line[6])
	return *domain.NewStock(ticker, domain.StockType(sType), category, subCategory,
		administrator, float32(bookValue), patrimony, float32(pvp))
}

func (h SaveStocksHandler) parseNumber(value string) (float64, error) {
	hydratedValue := strings.ReplaceAll(strings.ReplaceAll(value, ".", ""), ",", ".")
	return strconv.ParseFloat(hydratedValue, 64)
}

func (h SaveStocksHandler) extractCategoryAndSubcategory(categoryAndSubcategory string) (string, string) {
	values := strings.Split(categoryAndSubcategory, ":")
	category := values[0]
	subCategory := values[0]

	if len(values) == 2 {
		subCategory = values[1]
	}

	return category, subCategory
}
