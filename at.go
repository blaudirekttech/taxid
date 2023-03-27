package taxid

func isAustrianTaxIDValid(id string) bool {
	taxIDLength := len(id)
	if taxIDLength != 9 {
		return false
	}
	c, ok := convertToIntSlice(id)
	if !ok {
		return false
	}

	sum := c[0] + c[2] + c[4] + c[6]
	sum += sumDigits(2*c[1]) + sumDigits(2*c[3]) + sumDigits(2*c[5]) + sumDigits(2*c[7])
	check := digitAt(100-sum, 1)
	return c[8] == check
}
