package bigcomplex

import "math/big"

type Rat struct {
	r *big.Rat
	i *big.Rat
}

func NewCmplRat(r, i *big.Rat) *Rat {
	return &Rat{r, i}
}

func (r0 *Rat) Plus(r1 *Rat) *Rat {
	return &Rat{new(big.Rat).Add(r0.r, r1.r), new(big.Rat).Add(r0.i, r1.i)}
}

func (r0 *Rat) Minus(r1 *Rat) *Rat {
	return &Rat{new(big.Rat).Sub(r0.r, r1.r), new(big.Rat).Sub(r0.i, r1.i)}
}

// (a + bi) * (c + di) = (ac - bd) + (ad + bc)i
func (r0 *Rat) Times(r1 *Rat) *Rat {
	ac := new(big.Rat).Mul(r0.r, r1.r)
	bd := new(big.Rat).Mul(r0.i, r1.i)
	ad := new(big.Rat).Mul(r0.r, r1.i)
	bc := new(big.Rat).Mul(r0.i, r1.r)
	return &Rat{new(big.Rat).Sub(ac, bd), new(big.Rat).Add(ad, bc)}
}

// (a + bi) / (c + di) = ((ac + bd) + (bc - ad)i) / (c^2 + d^2)
func (r0 *Rat) Divides(r1 *Rat) *Rat {
	ac := new(big.Rat).Mul(r0.r, r1.r)
	bd := new(big.Rat).Mul(r0.i, r1.i)
	ad := new(big.Rat).Mul(r0.r, r1.i)
	bc := new(big.Rat).Mul(r0.i, r1.r)
	cc := new(big.Rat).Mul(r1.r, r1.r)
	dd := new(big.Rat).Mul(r1.i, r1.i)
	acPbd := new(big.Rat).Add(ac, bd)
	bcMad := new(big.Rat).Sub(bc, ad)
	ccPdd := new(big.Rat).Add(cc, dd)
	return &Rat{new(big.Rat).Quo(acPbd, ccPdd), new(big.Rat).Quo(bcMad, ccPdd)}
}

// SquaredAbs of a + bi = a^2 + b^2
func (r0 *Rat) SquaredAbs() *big.Rat {
	aa := new(big.Rat).Mul(r0.r, r0.r)
	bb := new(big.Rat).Mul(r0.i, r0.i)
	return new(big.Rat).Add(aa, bb)
}
