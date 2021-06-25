package even_odd

func IsEven(i int) bool {
	if i%2 == 1 || i%2 == -1 {
		return false
	}
	return true
}
