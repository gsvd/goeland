package validation

import (
	"strings"

	"github.com/gsvd/goeland/pkg/errorsx"
)

func ValidateXMPPAddress(address string) error {
	var errs errorsx.ValidationErrors

	if address == "" {
		errs.Add("address", errorsx.ErrCodeEmptyAddress)
	} else {
		parts := strings.SplitN(address, "@", 2)
		if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
			errs.Add("address", errorsx.ErrCodeInvalidAddressFormat)
		}
	}

	if errs.HasErrors() {
		return errs.AsAPIError()
	}

	return nil
}
