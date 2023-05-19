package error

import "fmt"

type NotFound struct {
	message string
}

func NewNotFoundError(identifier any) NotFound {
	return NotFound{message: fmt.Sprintf("Recurso de identificador '%v' não encontrato.", identifier)}
}

func (e NotFound) Error() string {
	return e.message
}
