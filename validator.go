package taxid

import (
	"strings"
)

func Validate(country TaxCountry, taxID string) error {
	taxID = strings.ReplaceAll(taxID, " ", "")
	if taxID == "" {
		return ErrInvalidTaxID
	}

	switch country {
	case Austria:
		if isAustrianTaxIDValid(taxID) {
			return nil
		}
		return ErrInvalidTaxID

	case Germany:
		if isGermanTaxIDValid(taxID) {
			return nil
		}
		return ErrInvalidTaxID

	default:
		return ErrUnsupportedCountry
	}
}
