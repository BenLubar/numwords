package numwords

// 0 <= n < 1000
func piece(buf []byte, n int) []byte {
	if n == 0 {
		return append(buf, " zero"...)
	}

	switch n / 100 {
	case 1:
		buf = append(buf, " one hundred"...)
	case 2:
		buf = append(buf, " two hundred"...)
	case 3:
		buf = append(buf, " three hundred"...)
	case 4:
		buf = append(buf, " four hundred"...)
	case 5:
		buf = append(buf, " five hundred"...)
	case 6:
		buf = append(buf, " six hundred"...)
	case 7:
		buf = append(buf, " seven hundred"...)
	case 8:
		buf = append(buf, " eight hundred"...)
	case 9:
		buf = append(buf, " nine hundred"...)
	}

	switch n / 10 % 10 {
	case 1:
		switch n % 10 {
		case 0:
			buf = append(buf, " ten"...)
		case 1:
			buf = append(buf, " eleven"...)
		case 2:
			buf = append(buf, " twelve"...)
		case 3:
			buf = append(buf, " thirteen"...)
		case 4:
			buf = append(buf, " fourteen"...)
		case 5:
			buf = append(buf, " fifteen"...)
		case 6:
			buf = append(buf, " sixteen"...)
		case 7:
			buf = append(buf, " seventeen"...)
		case 8:
			buf = append(buf, " eighteen"...)
		case 9:
			buf = append(buf, " nineteen"...)
		}
	case 2:
		buf = append(buf, " twenty"...)
	case 3:
		buf = append(buf, " thirty"...)
	case 4:
		buf = append(buf, " forty"...)
	case 5:
		buf = append(buf, " fifty"...)
	case 6:
		buf = append(buf, " sixty"...)
	case 7:
		buf = append(buf, " seventy"...)
	case 8:
		buf = append(buf, " eighty"...)
	case 9:
		buf = append(buf, " ninety"...)
	}

	if n/10%10 != 1 {
		switch n % 10 {
		case 1:
			buf = append(buf, " one"...)
		case 2:
			buf = append(buf, " two"...)
		case 3:
			buf = append(buf, " three"...)
		case 4:
			buf = append(buf, " four"...)
		case 5:
			buf = append(buf, " five"...)
		case 6:
			buf = append(buf, " six"...)
		case 7:
			buf = append(buf, " seven"...)
		case 8:
			buf = append(buf, " eight"...)
		case 9:
			buf = append(buf, " nine"...)
		}
	}

	return buf
}
