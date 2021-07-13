package kzg

import (
	"github.com/69b33ebd29f/kzg-hyper/ff"
	gmcl "github.com/69b33ebd29f/mcl-wrapper"
)

// GenerateTestingSetup creates a setup of n values from the given secret. **for testing purposes only**
func GenerateTestingSetup(secret string, n uint64) ([]gmcl.G1, []gmcl.G2) {
	var s gmcl.Fr
	ff.SetFr(&s, secret)

	var sPow gmcl.Fr
	ff.CopyFr(&sPow, &ff.ONE)

	s1Out := make([]gmcl.G1, n, n)
	s2Out := make([]gmcl.G2, n, n)
	for i := uint64(0); i < n; i++ {
		gmcl.G1Mul(&s1Out[i], &ff.GenG1, &sPow)
		gmcl.G2Mul(&s2Out[i], &ff.GenG2, &sPow)
		var tmp gmcl.Fr
		ff.CopyFr(&tmp, &sPow)
		gmcl.FrMul(&sPow, &tmp, &s)
	}
	return s1Out, s2Out
}

// These are testing utils
func testPoly(polynomial ...uint64) []gmcl.Fr {
	n := len(polynomial)
	polynomialFr := make([]gmcl.Fr, n, n)
	for i := 0; i < n; i++ {
		ff.AsFr(&polynomialFr[i], polynomial[i])
	}
	return polynomialFr
}

func generateSetup(secret string, n uint64) ([]gmcl.G1, []gmcl.G2) {
	var s gmcl.Fr
	ff.SetFr(&s, secret)

	var sPow gmcl.Fr
	ff.CopyFr(&sPow, &ff.ONE)

	s1Out := make([]gmcl.G1, n, n)
	s2Out := make([]gmcl.G2, n, n)
	for i := uint64(0); i < n; i++ {
		gmcl.G1Mul(&s1Out[i], &ff.GenG1, &sPow)
		gmcl.G2Mul(&s2Out[i], &ff.GenG2, &sPow)
		var tmp gmcl.Fr
		ff.CopyFr(&tmp, &sPow)
		gmcl.FrMul(&sPow, &tmp, &s)
	}
	return s1Out, s2Out
}
