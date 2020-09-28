package bigcomplex

import "math/big"

type Float struct {
	r *big.Float
	i *big.Float
}

func NewCmplFloat(r, i *big.Float) *Float {
	return &Float{r, i}
}

func (f0 *Float) Plus(f1 *Float) *Float {
	return &Float{new(big.Float).Add(f0.r, f1.r), new(big.Float).Add(f0.i, f1.i)}
}

func (f0 *Float) Minus(f1 *Float) *Float {
	return &Float{new(big.Float).Sub(f0.r, f1.r), new(big.Float).Sub(f0.i, f1.i)}
}

// (a + bi) * (c + di) = (ac - bd) + (ad + bc)i
func (f0 *Float) Times(f1 *Float) *Float {
	ac := new(big.Float).Mul(f0.r, f1.r)
	bd := new(big.Float).Mul(f0.i, f1.i)
	ad := new(big.Float).Mul(f0.r, f1.i)
	bc := new(big.Float).Mul(f0.i, f1.r)
	return &Float{new(big.Float).Sub(ac, bd), new(big.Float).Add(ad, bc)}
}

// (a + bi) / (c + di) = ((ac + bd) + (bc - ad)i) / (c^2 + d^2)
func (f0 *Float) Divides(f1 *Float) *Float {
	ac := new(big.Float).Mul(f0.r, f1.r)
	bd := new(big.Float).Mul(f0.i, f1.i)
	ad := new(big.Float).Mul(f0.r, f1.i)
	bc := new(big.Float).Mul(f0.i, f1.r)
	cc := new(big.Float).Mul(f1.r, f1.r)
	dd := new(big.Float).Mul(f1.i, f1.i)
	acPbd := new(big.Float).Add(ac, bd)
	bcMad := new(big.Float).Sub(bc, ad)
	ccPdd := new(big.Float).Add(cc, dd)
	return &Float{new(big.Float).Quo(acPbd, ccPdd), new(big.Float).Quo(bcMad, ccPdd)}
}

// SquaredAbs of a + bi = a^2 + b^2
func (f0 *Float) SquaredAbs() *big.Float {
	aa := new(big.Float).Mul(f0.r, f0.r)
	bb := new(big.Float).Mul(f0.i, f0.i)
	return new(big.Float).Add(aa, bb)
}
