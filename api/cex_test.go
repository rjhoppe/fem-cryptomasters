package api_test

import (
	"testing"

	"github.com/rjhoppe/go-cryptomasters/api"
)

// pass empty crypto ticker
func TestAPICall(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		// no assert here
		// raise an error
		t.Error("Crypto ticker was not found")
	}
}
