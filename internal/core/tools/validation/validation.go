package core_validation

func IsIntInBounds(number, minValue, maxValue int) bool {
	return minValue <= number && number <= maxValue
}
