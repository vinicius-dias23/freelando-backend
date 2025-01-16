package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func readJSON(r *http.Request, data any) error {
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must have only a single JSON value")
	}

	return nil
}

func writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func errorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()

	return writeJSON(w, statusCode, payload)
}

func validateStruct(data interface{}) error {
	val := reflect.ValueOf(data)

	if val.Kind() != reflect.Struct {
		return errors.New(
			"necessário passar uma struct como parâmetro da função para a validação dos campos",
		)
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := val.Type().Field(i).Name

		switch field.Kind() {
		case reflect.String:
			if field.String() == "" {
				return fmt.Errorf("campo '%s' é obrigatório", fieldName)
			}
		case reflect.Int, reflect.Int64:
			if field.Int() <= 0 {
				return fmt.Errorf("campo '%s' deve ser maior que 0", fieldName)
			}
		}
	}

	return nil
}
