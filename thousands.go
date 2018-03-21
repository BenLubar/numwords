package numwords

import "math/big"

var simpleThousandParts = [...]string{
	"n",
	"m",
	"b",
	"tr",
	"quadr",
	"quint",
	"sext",
	"sept",
	"oct",
	"non",
	"dec",
	"undec",
	"duodec",
	"tredec",
	"quattuordec",
	"quindec",
	"sexdec",
	"septendec",
	"octodec",
	"novemdec",
	"vigint",
}

// 0 <= n
func thousands(buf []byte, n int64) []byte {
	if n == 0 {
		return buf
	}
	if n == 1 {
		return append(buf, " thousand"...)
	}
	n--
	buf = append(buf, ' ')
	if n < 1000 {
		return append(thousandPart(buf, n), "illion"...)
	}

	for i, j := int64(6), int64(1e18); i >= 0; i, j = i-1, j/1000 {
		if part := n / j % 1000; part != 0 {
			buf = thousandPart(buf, part)
			buf = append(buf, "illi"...)
		}
	}

	return append(buf, "on"...)
}

// 0 <= n
func bigThousands(buf []byte, n *big.Int) []byte {
	if n.IsInt64() {
		return thousands(buf, n.Int64())
	}

	var i big.Int
	var j big.Int
	j.Set(bigThousand)

	for n.Cmp(&j) > 0 {
		i.Add(&i, bigOne)
		j.Mul(&j, bigThousand)
	}
	j.Div(&j, bigThousand)

	var nMinusOne big.Int
	nMinusOne.Sub(n, bigOne)

	buf = append(buf, ' ')
	var p big.Int

	for i.Sign() >= 0 {
		p.Div(&nMinusOne, &j)
		p.Mod(&p, bigThousand)
		if part := int(p.Int64()); part != 0 {
			buf = piece(buf, part)
			buf = bigThousands(buf, &i)
		}
		if part := p.Int64(); part != 0 {
			buf = thousandPart(buf, part)
			buf = append(buf, "illi"...)
		}

		i.Sub(&i, bigOne)
		j.Div(&j, bigThousand)
	}

	return append(buf, "on"...)
}

// 0 <= n < 1000
func thousandPart(buf []byte, n int64) []byte {
	if n < int64(len(simpleThousandParts)) {
		return append(buf, simpleThousandParts[n]...)
	}

	s, x, nm := false, false, false

	switch n % 10 {
	case 1:
		buf = append(buf, "un"...)
	case 2:
		buf = append(buf, "duo"...)
	case 3:
		buf = append(buf, "tre"...)
		s = true
	case 4:
		buf = append(buf, "quattuor"...)
	case 5:
		buf = append(buf, "quinqua"...)
	case 6:
		buf = append(buf, "se"...)
		x = true
	case 7:
		buf = append(buf, "septe"...)
		nm = true
	case 8:
		buf = append(buf, "octo"...)
	case 9:
		buf = append(buf, "nove"...)
		nm = true
	}

	switch n / 10 % 10 {
	case 1, 6, 7:
		if nm {
			buf = append(buf, 'n')
		}
	case 2:
		if nm {
			buf = append(buf, 'm')
		} else if s {
			buf = append(buf, 's')
		}
	case 3, 4, 5:
		if nm {
			buf = append(buf, 'n')
		} else if s {
			buf = append(buf, 's')
		}
	case 8:
		if nm {
			buf = append(buf, 'm')
		} else if x {
			buf = append(buf, 'x')
		}
	}

	switch n / 10 % 10 {
	case 1:
		buf = append(buf, "deci"...)
	case 2:
		buf = append(buf, "viginti"...)
	case 3:
		buf = append(buf, "triginta"...)
	case 4:
		buf = append(buf, "quadraginta"...)
	case 5:
		buf = append(buf, "quinquaginta"...)
	case 6:
		buf = append(buf, "sexaginta"...)
	case 7:
		buf = append(buf, "septuaginta"...)
	case 8:
		buf = append(buf, "octoginta"...)
	case 9:
		buf = append(buf, "nonaginta"...)
	}

	switch n / 10 % 100 {
	case 10:
		if nm {
			buf = append(buf, 'n')
		} else if x {
			buf = append(buf, 'x')
		}
	case 20, 60, 70:
		if nm {
			buf = append(buf, 'n')
		}
	case 30, 40, 50:
		if nm {
			buf = append(buf, 'n')
		} else if s {
			buf = append(buf, 's')
		}
	case 80:
		if nm {
			buf = append(buf, 'm')
		} else if x {
			buf = append(buf, 'x')
		}
	}

	switch n / 100 % 10 {
	case 1:
		buf = append(buf, "centi"...)
	case 2:
		buf = append(buf, "ducenti"...)
	case 3:
		buf = append(buf, "trecenti"...)
	case 4:
		buf = append(buf, "quadringenti"...)
	case 5:
		buf = append(buf, "quingenti"...)
	case 6:
		buf = append(buf, "sescenti"...)
	case 7:
		buf = append(buf, "septingenti"...)
	case 8:
		buf = append(buf, "octingenti"...)
	case 9:
		buf = append(buf, "nongenti"...)
	}

	return buf
}
