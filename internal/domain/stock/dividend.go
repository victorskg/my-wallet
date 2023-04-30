package stock

import "time"

type DType string

const (
	Income       DType = "income"
	Amortization DType = "amortization"
)

type Dividend struct {
	value       float64
	baseDate    time.Time
	paymentDate time.Time
	dType       DType
}

func (d Dividend) BaseDate() time.Time {
	return d.baseDate
}

func (d Dividend) Value() float64 {
	return d.value
}
