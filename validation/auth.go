package validation

import (
	"encoding/json"
	"errors"
	"github.com/bloodblue999/umhelp/presenter/req"
	"io"
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

	return r, nil
}
