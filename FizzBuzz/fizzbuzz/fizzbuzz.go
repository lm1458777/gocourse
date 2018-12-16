package fizzbuzz

import "strconv"

func NumberToString(n int) string {
	const Fizz = "Fizz"
	const Buzz = "Buzz"
	const FizzBuzz = Fizz + "," + Buzz

	if n == 0 {
		return "0"
	}

	if n%3 == 0 && n%5 == 0 {
		return FizzBuzz
	}

	if n%3 == 0 {
		return Fizz
	}

	if n%5 == 0 {
		return Buzz
	}

	return strconv.Itoa(n)
}
