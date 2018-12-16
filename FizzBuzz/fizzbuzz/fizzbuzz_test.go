package fizzbuzz

import "testing"

func TestNumberToString(t *testing.T) {
	type testCase struct {
		in   int
		want string
	}

	cases := [] testCase{
		{0, "0"},
		{1, "1"},
		{3, "Fizz"},
		{5, "Buzz"},
		{6, "Fizz"},
		{10, "Buzz"},
		{15, "Fizz,Buzz"},
		{30, "Fizz,Buzz"},
	}

	for _, c := range cases {
		got := NumberToString(c.in)
		if got != c.want {
			t.Errorf("NumberToString(%v) == %v, want: %v", c.in, got, c.want)
		}
	}
}
