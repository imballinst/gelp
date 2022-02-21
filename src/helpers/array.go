package helpers

func GetUniqueIntegers(values []int) []int {
	var includedInts = map[int]bool{}
	var result = []int{}

	for _, val := range values {
		_, doesntExist := includedInts[val]
		if !doesntExist {
			includedInts[val] = true
			result = append(result, val)
		}
	}

	return result
}

func GetRangeArrayFromTwoIntegers(start, end int) []int {
	var rangeArray []int
	bigger, smaller := start, end
	if bigger < smaller {
		bigger, smaller = smaller, bigger
	}

	i := bigger

	for i >= smaller {
		rangeArray = append(rangeArray, i)
		i--
	}

	return rangeArray
}
