//line :1
package xaIY_5DR

import (
	"encoding/binary"
	"math/bits"
)

var Le5rqThl = [64]uint32{
	0x428a2f98, 0x71374491, 0xb5c0fbcf, 0xe9b5dba5, 0x3956c25b, 0x59f111f1, 0x923f82a4, 0xab1c5ed5,
	0xd807aa98, 0x12835b01, 0x243185be, 0x550c7dc3, 0x72be5d74, 0x80deb1fe, 0x9bdc06a7, 0xc19bf174,
	0xe49b69c1, 0xefbe4786, 0x0fc19dc6, 0x240ca1cc, 0x2de92c6f, 0x4a7484aa, 0x5cb0a9dc, 0x76f988da,
	0x983e5152, 0xa831c66d, 0xb00327c8, 0xbf597fc7, 0xc6e00bf3, 0xd5a79147, 0x06ca6351, 0x14292967,
	0x27b70a85, 0x2e1b2138, 0x4d2c6dfc, 0x53380d13, 0x650a7354, 0x766a0abb, 0x81c2c92e, 0x92722c85,
	0xa2bfe8a1, 0xa81a664b, 0xc24b8b70, 0xc76c51a3, 0xd192e819, 0xd6990624, 0xf40e3585, 0x106aa070,
	0x19a4c116, 0x1e376c08, 0x2748774c, 0x34b0bcb5, 0x391c0cb3, 0x4ed8aa4a, 0x5b9cca4f, 0x682e6ff3,
	0x748f82ee, 0x78a5636f, 0x84c87814, 0x8cc70208, 0x90befffa, 0xa4506ceb, 0xbef9a3f7, 0xc67178f2,
}

const chunkSize = 64

func ys7QvSCZ(kdNdo9nm []byte) []byte {
	EyAMyNhm :=  /*line _zuFr0SI.go:1*/ make([]byte, 8)
	 /*line QEfY8l0L.go:1*/ binary.BigEndian.PutUint64(EyAMyNhm,  /*line LvWzH3pl.go:1*/ uint64( /*line Foh2RtFe.go:1*/ len(kdNdo9nm)*8))
	kdNdo9nm =  /*line _adNFmvz.go:1*/ append(kdNdo9nm, 0x80)
	for ( /*line axtZoL5I.go:1*/ len(kdNdo9nm)+8)%64 != 0 {
		kdNdo9nm =  /*line BFra9NzI.go:1*/ append(kdNdo9nm, 0x00)
	}
	kdNdo9nm =  /*line rAkU2XO6.go:1*/ append(kdNdo9nm, EyAMyNhm...)

	return kdNdo9nm
}

func Vw2WoqqX(kdNdo9nm []byte) [32]byte {
	kdNdo9nm =  /*line m0EOGLTU.go:1*/ ys7QvSCZ(kdNdo9nm)

	goAW401K, fCqjYtt6, vzMizMAy, itIP2ByV, ri8wxafH, tjZVztDa, x3btOA2l, oaOEcnee :=  /*line _AkaK2iS.go:1*/ uint32(0x6a09e667),  /*line GIMB_0TR.go:1*/ uint32(0xbb67ae85),  /*line QPCjLSkd.go:1*/ uint32(0x3c6ef372),  /*line XOjPr79k.go:1*/ uint32(0xa54ff53a),
		 /*line KUZi3DCC.go:1*/ uint32(0x510e527f),  /*line wiUQxUVy.go:1*/ uint32(0x9b05688c),  /*line kDo1pdPS.go:1*/ uint32(0x1f83d9ab),  /*line uMEuq84M.go:1*/ uint32(0x5be0cd19)

	for iMByIZp0 := 0; iMByIZp0 <  /*line ZYPU5dBd.go:1*/ len(kdNdo9nm); iMByIZp0 += chunkSize {
		var iBmv8zrT [64]uint32
		for idzQKtZu := 0; idzQKtZu*4 < chunkSize; idzQKtZu++ {
			iBmv8zrT[idzQKtZu] =  /*line h8egaGOj.go:1*/ binary.BigEndian.Uint32(kdNdo9nm[iMByIZp0+idzQKtZu*4 : iMByIZp0+(idzQKtZu+1)*4])
		}

		for idzQKtZu := 16; idzQKtZu < 64; idzQKtZu++ {
			iAMkUhcH :=  /*line fhxicu9z.go:1*/ bits.RotateLeft32(iBmv8zrT[idzQKtZu-15], -7) ^  /*line Glen5Btc.go:1*/ bits.RotateLeft32(iBmv8zrT[idzQKtZu-15], -18) ^ (iBmv8zrT[idzQKtZu-15] >> 3)
			kw0mLN4V :=  /*line gnct4evx.go:1*/ bits.RotateLeft32(iBmv8zrT[idzQKtZu-2], -17) ^  /*line BFfzqrS9.go:1*/ bits.RotateLeft32(iBmv8zrT[idzQKtZu-2], -19) ^ (iBmv8zrT[idzQKtZu-2] >> 10)
			iBmv8zrT[idzQKtZu] = iBmv8zrT[idzQKtZu-16] + iAMkUhcH + iBmv8zrT[idzQKtZu-7] + kw0mLN4V
		}

		jkaWZWr3, ytpv7Fq6, cM51f7as, c3DFzJxK, iaqIONzz, quIJzx4t, yYB7n2Kl, vB7RJcOO := goAW401K, fCqjYtt6, vzMizMAy, itIP2ByV, ri8wxafH, tjZVztDa, x3btOA2l, oaOEcnee
		for idzQKtZu := 0; idzQKtZu < 64; idzQKtZu++ {
			WmfGM9K_ :=  /*line XJFHiT98.go:1*/ bits.RotateLeft32(iaqIONzz, -6) ^  /*line DwIDJowJ.go:1*/ bits.RotateLeft32(iaqIONzz, -11) ^  /*line bNfPfBNz.go:1*/ bits.RotateLeft32(iaqIONzz, -25)
			cujd6iWL := (iaqIONzz & quIJzx4t) ^ ((^iaqIONzz) & yYB7n2Kl)
			khYsNgsw := vB7RJcOO + WmfGM9K_ + cujd6iWL + Le5rqThl[idzQKtZu] + iBmv8zrT[idzQKtZu]
			EUnrRCmx :=  /*line qCIdlOWW.go:1*/ bits.RotateLeft32(jkaWZWr3, -2) ^  /*line Lfc8aG7N.go:1*/ bits.RotateLeft32(jkaWZWr3, -13) ^  /*line j5lxZczU.go:1*/ bits.RotateLeft32(jkaWZWr3, -22)
			fAkQYphG := (jkaWZWr3 & ytpv7Fq6) ^ (jkaWZWr3 & cM51f7as) ^ (ytpv7Fq6 & cM51f7as)
			euX3kGs9 := EUnrRCmx + fAkQYphG
			vB7RJcOO = yYB7n2Kl
			yYB7n2Kl = quIJzx4t
			quIJzx4t = iaqIONzz
			iaqIONzz = c3DFzJxK + khYsNgsw
			c3DFzJxK = cM51f7as
			cM51f7as = ytpv7Fq6
			ytpv7Fq6 = jkaWZWr3
			jkaWZWr3 = khYsNgsw + euX3kGs9
		}
		goAW401K += jkaWZWr3
		fCqjYtt6 += ytpv7Fq6
		vzMizMAy += cM51f7as
		itIP2ByV += c3DFzJxK
		ri8wxafH += iaqIONzz
		tjZVztDa += quIJzx4t
		x3btOA2l += yYB7n2Kl
		oaOEcnee += vB7RJcOO
	}

	c_ehY6dh := [32]byte{}
	 /*line dRdKybWc.go:1*/ binary.BigEndian.PutUint32(c_ehY6dh[:4], goAW401K)
	 /*line w2QaDzJZ.go:1*/ binary.BigEndian.PutUint32(c_ehY6dh[4:8], fCqjYtt6)
	 /*line mCkuiB7d.go:1*/ binary.BigEndian.PutUint32(c_ehY6dh[8:12], vzMizMAy)
	 /*line B66KSLjd.go:1*/ binary.BigEndian.PutUint32(c_ehY6dh[12:16], itIP2ByV)
	 /*line E_GhpZ3y.go:1*/ binary.BigEndian.PutUint32(c_ehY6dh[16:20], ri8wxafH)
	 /*line La1RqG_M.go:1*/ binary.BigEndian.PutUint32(c_ehY6dh[20:24], tjZVztDa)
	 /*line lTcP8Yzs.go:1*/ binary.BigEndian.PutUint32(c_ehY6dh[24:28], x3btOA2l)
	 /*line UK_J7VtK.go:1*/ binary.BigEndian.PutUint32(c_ehY6dh[28:], oaOEcnee)

	return c_ehY6dh
}

var garbleActionID = "C5XDWc-kaLO2DFMyydkj"
