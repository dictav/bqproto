package bqproto

import (
	"math/big"
)

const (
	scaleDecimal = 1_000_000_000
)

var (
	MaxDecimal big.Rat
	MinDecimal big.Rat

	one = big.NewInt(1)
)

func init() {
	MaxDecimal = *big.NewRat(99999999999999999, 1)
	MaxDecimal.Mul(&MaxDecimal, big.NewRat(1000000000000, 1))
	MaxDecimal.Add(&MaxDecimal, big.NewRat(999999999999, 1))
	MaxDecimal.Add(&MaxDecimal, big.NewRat(999999999, scaleDecimal))

	MinDecimal = *big.NewRat(-1, 1)
	MinDecimal.Mul(&MinDecimal, &MaxDecimal)
}

func encodeDecimal(n *big.Rat) []byte {
	sign := n.Sign()

	if sign == 0 {
		return []byte{0}
	}

	bi := big.NewInt(scaleDecimal)

	bi.Mul(bi, n.Num())
	bi.Div(bi, n.Denom())

	if sign == 1 {
		return reverseBytes(bi.Bytes())
	}

	bilen := uint(bi.BitLen()/8+1) * 8

	bi.Add(bi, new(big.Int).Lsh(one, bilen))

	b := bi.Bytes()

	// special case
	//
	// ex.)
	// -8388608 [11111111 10000000 0 0] => [10000000 0 0]
	//   -32768 [11111111 10000000 0]   => [10000000 0]
	//     -128 [11111111 10000000]     => [10000000]
	if len(b) > 1 && b[0] == 0xff && b[1] == 0x80 {
		b = b[1:]
	}

	return reverseBytes(b)
}

func reverseBytes(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}
