package stock

import "time"

type Dividend struct {
	value       float32
	baseDate    time.Time
	paymentDate time.Time
	yeld        float32
}
