package numwords

var digits = [...]string{
	" zero",
	" one",
	" two",
	" three",
	" four",
	" five",
	" six",
	" seven",
	" eight",
	" nine",
}

var teens = [...]string{
	" ten",
	" eleven",
	" twelve",
	" thirteen",
	" fourteen",
	" fifteen",
	" sixteen",
	" seventeen",
	" eighteen",
	" nineteen",
}

var tens = [...]string{
	"",
	" ten",
	" twenty",
	" thirty",
	" forty",
	" fifty",
	" sixty",
	" seventy",
	" eighty",
	" ninety",
}

// 0 <= n < 1000
func piece(buf []byte, n int) []byte {
	if n == 0 {
		return append(buf, digits[0]...)
	}

	if h := n / 100; h != 0 {
		buf = append(append(buf, digits[h]...), " hundred"...)
	}

	if t := n / 10 % 10; t == 1 {
		buf = append(buf, teens[n%10]...)
	} else if t != 0 {
		buf = append(buf, tens[t]...)
	}

	if o := n % 10; o != 0 && n/10%10 != 1 {
		buf = append(buf, digits[o]...)
	}

	return buf
}
