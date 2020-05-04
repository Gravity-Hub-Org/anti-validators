package converter

import "math/big"

func ConvertStrAmount(amount string, fromDecimals uint8, toDecimals uint8) string {
	bigAmount := big.NewFloat(0)
	bigAmount, _ = bigAmount.SetString(amount)

	fromBase := big.NewInt(int64(fromDecimals))
	toBase := big.NewInt(int64(toDecimals))

	fromBase.Exp(big.NewInt(10), fromBase, nil)
	toBase.Exp(big.NewInt(10), toBase, nil)

	bigAmount = bigAmount.Quo(bigAmount, new(big.Float).SetInt(fromBase))
	bigIntAmount := new(big.Int)
	bigAmount.Mul(bigAmount, new(big.Float).SetInt(toBase)).Int(bigIntAmount)
	return bigIntAmount.String()
}
