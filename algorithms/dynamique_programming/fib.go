package fib

func Fib(n int) int {
	if n == 1 {
		return 1
	}
	if n == 0 {
		return 0
	}
	return Fib(n-1) + Fib(n-2)
}
