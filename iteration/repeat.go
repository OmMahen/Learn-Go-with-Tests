package iteration

func Repeat(character string, times int) (result string) {
	for i := 0; i < times; i++ {
		result += character
	}
	return
}
