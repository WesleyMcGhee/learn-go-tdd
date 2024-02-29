package iteration

const repeatCount = 5

func Repeat(character string, repeatAmount int) string {
	if repeatAmount == 0 {
		repeatAmount = repeatCount
	}
	returnString := ""
	for i := 0; i < repeatAmount; i++ {
		returnString += character
	}

	return returnString
}