package iteration

func Repeat(character string, number int) string {
	var repeated string
	var repeatCount = number
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}