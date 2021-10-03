package string

func Add(x, y int) int {
	return x + y
}

const repeatCount = 5

func Repeat(char string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += char
	}
	return repeated
}
