package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/rjhoppe/go-cryptomasters/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	if len(currency) != 3 {
		return nil, fmt.Errorf("%v is not a valid crypto ticker", currency)
	}
	upCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, upCurrency))
	if err != nil {
		return nil, err
	}
	var response CEXResponse
	if res.StatusCode == http.StatusOK {
		// waits for entire response and reads it with the io reader
		reader, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		// reuse and reassign previous err because it is in same block
		// not going to directly unmarshal a JSON object to your struct
		// need an intermediary struct that represents the JSON
		err = json.Unmarshal(reader, &response)
		if err != nil {
			return nil, err
		}

	} else {
		// create own error message
		return nil, fmt.Errorf("invalid status code received: %v", res.StatusCode)
	}
	rate := datatypes.Rate{Currency: currency, Price: response.Bid}
	// use &rate to reference the rate pointer
	return &rate, nil
}
