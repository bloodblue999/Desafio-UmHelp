package validation

import (
	"encoding/json"
	"fmt"
	"github.com/bloodblue999/umhelp/presenter/req"
	"io"
)

func GetAndValidateMoneyTransaction(rc io.ReadCloser) (r *req.CreateMoneyTransaction, err error) {
	defer rc.Close()

	body, err := io.ReadAll(rc)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	if r.MoneyValue <= 0 {
		return nil, fmt.Errorf("money value must be greater than 0, actual is `%d`", r.MoneyValue)
	}

	if r.ReceiverID == r.SenderID {
		return nil, fmt.Errorf("senderID and receiverID wallets must be diferent")
	}

	return r, nil
}
