package validation

import (
	"encoding/json"
	"errors"
	"github.com/bloodblue999/umhelp/presenter/req"
	"io"
	"regexp"
)

func GetAndValidateLoginRequest(rc io.ReadCloser) (r *req.LoginRequest, err error) {
	defer rc.Close()

	body, err := io.ReadAll(rc)
	if err != nil {
		return nil, errors.New("cannot read requisition body")
	}

	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, errors.New("invalid requisition body parse")
	}

	regexDocumentValidator := regexp.MustCompile(`^[0-9]{3}\.[0-9]{3}\.[0-9]{3}-[0-9]{2}$`)
	if !regexDocumentValidator.MatchString(r.Document) {
		return nil, errors.New("invalid document number format. Correct format 123.456.789-00")
	}

	if len(r.Password) < 9 || len(r.Password) > 100 {
		return nil, errors.New("invalid password length. Password must be between 9 and 100 length ")
	}

	return r, nil
}
