package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/victorskg/my-wallet/pkg/http/response"
	"io"
	"log"
	"net/http"
	"strings"
)

func Deserialize[T any](input *T, w http.ResponseWriter, r *http.Request) error {
	contentTypeHeader := r.Header.Get("Content-Type")
	if contentTypeHeader != "" {
		if !strings.Contains(contentTypeHeader, "application/json") {
			msg := "O conteúdo da requisição é inválido."
			response.WriteResponseMessage(w, msg, http.StatusUnsupportedMediaType)
			return errors.New(msg)
		}
	}

	// Set up the decoder and call the DisallowUnknownFields() method on it.
	// This will cause Decode() to return a "json: unknown field ..." error
	// if it encounters any extra unexpected fields in the JSON. Strictly
	// speaking, it returns an error for "keys which do not match any
	// non-ignored, exported fields in the destination".
	dec := json.NewDecoder(r.Body)

	err := dec.Decode(input)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var msg string

		switch {
		// Catch any syntax errors in the JSON and send an error message
		// which interpolates the location of the problem to make it
		// easier for the client to fix.
		case errors.As(err, &syntaxError):
			msg = fmt.Sprintf("O corpo da requisição contem caracteres inválidos (na posição %d).", syntaxError.Offset)
			response.WriteResponseMessage(w, msg, http.StatusBadRequest)

		// In some circumstances Decode() may also return an
		// io.ErrUnexpectedEOF error for syntax errors in the JSON. There
		// is an open issue regarding this at
		// https://github.com/golang/go/issues/25956.
		case errors.Is(err, io.ErrUnexpectedEOF):
			msg = fmt.Sprintf("O corpo da requisição é inválido.")
			response.WriteResponseMessage(w, msg, http.StatusBadRequest)

		// Catch any type errors, like trying to assign a string in the
		// JSON request body to a int field in our Person struct. We can
		// interpolate the relevant field name and position into the error
		// message to make it easier for the client to fix.
		case errors.As(err, &unmarshalTypeError):
			msg = fmt.Sprintf("O corpo da requisição contém valores inválidos para o campo %q.", unmarshalTypeError.Field)
			response.WriteResponseMessage(w, msg, http.StatusBadRequest)

		// An io.EOF error is returned by Decode() if the request body is
		// empty.
		case errors.Is(err, io.EOF):
			msg = "O corpo da requisição é obrigatório."
			response.WriteResponseMessage(w, msg, http.StatusBadRequest)

		// Otherwise default to logging the error and sending a 500 Internal
		// Server Error response.
		default:
			log.Print(err.Error())
			response.WriteResponseMessage(w, err.Error(), http.StatusInternalServerError)
		}

		return err
	}

	// Call decode again, using a pointer to an empty anonymous struct as
	// the destination. If the request body only contained a single JSON
	// object this will return an io.EOF error. So if we get anything else,
	// we know that there is additional data in the request body.
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		msg := "A requisição deve conter um único objeto no corpo."
		response.WriteResponseMessage(w, msg, http.StatusBadRequest)
		return errors.New(msg)
	}

	return nil
}
