package converter

import "math/big"

func ConvertStrAmount(amount string, fromDecimals uint8, toDecimals uint8) string {
	bigAmount := big.NewInt(0)
	bigAmount.SetString(amount, 10)

	base := *big.NewInt(10)
	fromBase := base.Exp(&base, big.NewInt(int64(fromDecimals)), nil)
	toBase := base.Exp(&base, big.NewInt(int64(fromDecimals)), nil)

	bigAmount.Div(bigAmount, fromBase)
	return bigAmount.Mul(bigAmount, toBase).String()
}
