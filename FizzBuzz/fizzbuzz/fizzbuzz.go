package fizzbuzz

import "strconv"

func NumberToString(n int) string {
	const Fizz = "Fizz"
	const Buzz = "Buzz"
	const FizzBuzz = Fizz + "," + Buzz

	switch {
	case n == 0:
		return "0"
	case n%3 == 0 && n%5 == 0:
		return FizzBuzz
	case n%3 == 0:
		return Fizz
	case n%5 == 0:
		return Buzz
	}

	return strconv.Itoa(n)
}
