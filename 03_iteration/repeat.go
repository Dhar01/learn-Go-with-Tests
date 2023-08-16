package iteration

func Repeat(character string, count int) (repeated string) {
	if count == 0 {
		count = 1
	}

	for i := 0; i < count; i++ {
		repeated += character
	}

	return
}
