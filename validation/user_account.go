package validation

import (
	"encoding/json"
	"errors"
	"github.com/bloodblue999/umhelp/presenter/req"
	"io"
	"regexp"
)

func GetAndValidateCreateUserAccount(rc io.ReadCloser) (r *req.CreateUserAccount, err error) {
	defer rc.Close()

	body, err := io.ReadAll(rc)
	if err != nil {
		return nil, errors.New("cannot read requisition body")
	}

	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, errors.New("invalid requisition body parse")
	}

	if len(r.FirstName) < 3 || len(r.FirstName) > 30 {
		return nil, errors.New("invalid firstName size. first name size must be between 3 and 30 characters")
	}

	if len(r.LastName) < 3 || len(r.LastName) > 255 {
		return nil, errors.New("invalid lastName size. last name size must be between 3 and 255 characters")
	}

	regexDocumentValidator := regexp.MustCompile(`^[0-9]{3}\.[0-9]{3}\.[0-9]{3}-[0-9]{2}$`)
	if !regexDocumentValidator.MatchString(r.Document) {
		return nil, errors.New("invalid document number format. Correct format 123.456.789-00")
	}

	return r, nil
}
