package highOrder

func Somar(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}