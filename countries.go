package taxid

// TaxCountry is the country supported by taxid for validation checks
type TaxCountry int64

const (
	Germany TaxCountry = iota
	Austria
)
