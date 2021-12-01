package main

import (
	"fmt"
)

func Gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return Gcd(y, x%y)
}

func ExtendedGcd(a, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}
	d, x, y := ExtendedGcd(b, a%b)
	return d, y, (x - (a/b)*y)
}

func IsPrime(a int) bool {
	for i := 2; i*i <= a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

func TestExtendedGcd(a, b int) (ok bool) {
	d, x, y := ExtendedGcd(a, b)
	ok = a*x+b*y == d
	fmt.Printf("%d * %d + %d * %d = %d (ok = %t)\n", a, x, b, y, d, ok)
	return
}

func main2() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println(IsPrime(100))
	fmt.Println(IsPrime(21))
	fmt.Println(IsPrime(2))
	fmt.Println(IsPrime(3))
	fmt.Println(IsPrime(7))
	fmt.Println(IsPrime(1000000021))
}
