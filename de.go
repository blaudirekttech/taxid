package taxid

// https://de.wikipedia.org/wiki/Steuerliche_Identifikationsnummer#Aufbau_der_Identifikationsnummer
func isGermanTaxIDValid(id string) bool {
	taxIDLength := len(id)

	// TaxID has to have exactly 11 digits
	if taxIDLength != 11 {
		return false
	}
	// First digit cannot be 0
	if id[0] == '0' {
		return false
	}
	/*
	   Make sure that within the first ten digits:
	   1. one digit appears exactly twice or thrice
	   2. one or two digits appear zero times
	   3. and all other digits appear exactly once
	*/
	uniqueValues := uniqueCharsInAString(id[:10])
	if uniqueValues != 8 && uniqueValues != 9 {
		return false
	}

	intSlice, ok := convertToIntSlice(id)
	if !ok {
		return false
	}

	// Checksum
	var sum int

	product := 10
	for i := 0; i < taxIDLength-1; i++ {
		sum = (intSlice[i] + product) % 10
		if sum == 0 {
			sum = 10
		}

		product = (sum * 2) % 11
	}

	checksum := 11 - product
	if checksum == 10 {
		checksum = 0
	}

	lastDigit := intSlice[taxIDLength-1]

	return lastDigit == checksum
}
