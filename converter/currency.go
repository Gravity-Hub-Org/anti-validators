package converter

import "math/big"

// TODO: Ref to number
func StrConvert(amount string, fromDecimals int, toDecimals int) string {
	bigAmount := big.NewInt(0)
	bigAmount.SetString(amount, 10)

	base := *big.NewInt(10)
	fromBase := base.Exp(&base, big.NewInt(int64(fromDecimals)), nil)
	toBase := base.Exp(&base, big.NewInt(int64(fromDecimals)), nil)

	bigAmount.Div(bigAmount, fromBase)
	return bigAmount.Mul(bigAmount, toBase).String()
}
