package fibonacci

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return fibonacci(n-2) + fibonacci(n-1)
}

func Fibonacci(n ...int) map[int]int {
	var res map[int]int = make(map[int]int, len(n))
	for _, i := range n {
		res[i] = fibonacci(i)
	}
	return res
}
