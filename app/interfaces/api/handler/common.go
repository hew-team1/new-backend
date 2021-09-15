package handler

import (
	"encoding"
	"io/ioutil"
	"net/http"

	"github.com/go-openapi/strfmt"
)

type requestModel interface {
	encoding.BinaryUnmarshaler
	Validate(registry strfmt.Registry) error
}

func parseModelRequest(r *http.Request, i requestModel) error {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}

	i.UnmarshalBinary(b)
	if err := i.Validate(strfmt.Default); err != nil {
		return err
	}
	return nil
}
