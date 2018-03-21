package numwords

import (
	"math/big"
	"math/bits"
)

func Uint8(n uint8) string {
	return string(piece(nil, int(n))[1:])
}

func Int8(n int8) string {
	if n < 0 {
		return "negative " + Uint8(uint8(-n))
	}
	return Uint8(uint8(n))
}

func Uint16(n uint16) string {
	return Uint64(uint64(n))
}

func Int16(n int16) string {
	return Int64(int64(n))
}

func Uint32(n uint32) string {
	return Uint64(uint64(n))
}

func Int32(n int32) string {
	return Int64(int64(n))
}

func Uint(n uint) string {
	return Uint64(uint64(n))
}

func Int(n int) string {
	return Int64(int64(n))
}

func Uint64(n uint64) string {
	if n == 0 {
		return "zero"
	}

	var buf []byte
	for i, j := int64(6), uint64(1e18); i >= 0; i, j = i-1, j/1000 {
		if part := int(n / j % 1000); part != 0 {
			buf = piece(buf, part)
			buf = thousands(buf, i)
		}
	}

	return string(buf[1:])
}

func Int64(n int64) string {
	if n < 0 {
		return "negative " + Uint64(uint64(-n))
	}
	return Uint64(uint64(n))
}

func Big(n *big.Int) string {
	switch n.Sign() {
	case -1:
		var abs big.Int
		abs.Abs(n)
		return "negative " + nat(&abs)
	case 0:
		return "zero"
	case 1:
		return nat(n)
	}

	panic("unreachable")
}

// minimum number of bits required to represent a number that is more than 1000^(2^63 - 1)
const big63Limit = 91918136096478347264

// if this next line is an error, the code at the start of nat needs to be enabled
var _ = [65 - bits.UintSize]struct{}{}

var bigThousand = big.NewInt(1000)
var bigOne = big.NewInt(1)

func nat(n *big.Int) string {
	// // not needed on currently-existing hardware
	// if n.BitLen() >= big63Limit {
	//	 return bigNat(n)
	// }

	var i int64
	var j big.Int
	j.Set(bigThousand)

	for n.Cmp(&j) >= 0 {
		i++
		j.Mul(&j, bigThousand)
	}

	j.Div(&j, bigThousand)

	var p big.Int
	var buf []byte

	for i >= 0 {
		p.Div(n, &j)
		p.Mod(&p, bigThousand)
		if part := int(p.Int64()); part != 0 {
			buf = piece(buf, part)
			buf = thousands(buf, i)
		}

		i--
		j.Div(&j, bigThousand)
	}

	return string(buf[1:])
}

func bigNat(n *big.Int) string {
	var i big.Int
	var j big.Int
	j.Set(bigThousand)

	for n.Cmp(&j) >= 0 {
		i.Add(&i, bigOne)
		j.Mul(&j, bigThousand)
	}

	j.Div(&j, bigThousand)

	var p big.Int
	var buf []byte

	for i.Sign() >= 0 {
		p.Div(n, &j)
		p.Mod(&p, bigThousand)
		if part := int(p.Int64()); part != 0 {
			buf = piece(buf, part)
			buf = bigThousands(buf, &i)
		}

		i.Sub(&i, bigOne)
		j.Div(&j, bigThousand)
	}

	return string(buf[1:])
}
