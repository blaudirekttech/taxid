package taxid

func isAustrianTaxIDValid(id string) bool {
	const taxIDLength = 9

	// convert id to int64 slice for further processing of the id
	c, ok := convertToInt64Slice(id)
	if !ok || len(c) != taxIDLength {
		return false
	}

	sum := c[0] + c[2] + c[4] + c[6]
	sum += sumDigits(2*c[1]) + sumDigits(2*c[3]) + sumDigits(2*c[5]) + sumDigits(2*c[7])

	if sum < 0 || sum > 100 {
		return false
	}

	check := (100 - sum) % 10 // Get the last digit of the sum

	return c[taxIDLength-1] == check
}
