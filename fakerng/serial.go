/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package fakerng

import (
	"math"
	"math/big"
	"strconv"
	"strings"
)

// Converter maps a N-digit interger to another interger with same digit.
//
// It uses modular multiplicative inverse to build a bijection function and its
// inverse function. The domain and codomain of both functions are n-digit integer (
// 0 ~ 10^n-1, inclusive).
type Converter struct {
	digits int
	base   *big.Int
	encKey *big.Int
	decKey *big.Int
	key    int
	over   [2]int
}

func (r *Converter) source() (ret *source) {
	buf := &source{}
	buf.line("fakerng.MustRawConverter(%d, %d)", r.digits, r.key)
	return buf
}

// Source dumps golang source of the Converter instance
func (r *Converter) Source() (ret string) {
	return r.source().String()
}

func (r *Converter) s2i(str string) (ret int, err error) {
	if err = isDigit(str, r.digits); err != nil {
		return
	}

	i64, _ := strconv.ParseInt(str, 10, 64)
	ret = int(i64)
	return
}

func (r *Converter) i2s(i int) (ret string) {
	ret = strconv.Itoa(i)
	if l := len(ret); l < r.digits {
		ret = strings.Repeat("0", r.digits-l) + ret
	}

	return
}

// Enc converts a n-digit integer (data) to another n-digit integer (code)
func (r *Converter) Enc(serial int) (ret int) {
	switch serial {
	case 0:
		return r.key
	case 1:
		return r.over[1]
	case r.over[0]:
		return 0
	}

	x := big.NewInt(int64(serial))
	x = x.Mul(x, r.encKey)
	x = x.Mod(x, r.base)
	return int(x.Int64())
}

// Encode converts a n-digit string (data) to another n-digit string (code)
//
// If str is not n-digit string, the err will be ErrFormat.
func (r *Converter) Encode(str string) (ret string, err error) {
	i, err := r.s2i(str)
	if err != nil {
		return
	}

	ret = r.i2s(r.Enc(i))
	return
}

// Dec converts a n-digit integer (code) back to another n-digit integer (data)
func (r *Converter) Dec(code int) (ret int) {
	switch code {
	case 0:
		return r.over[0]
	case r.key:
		return 0
	case r.over[1]:
		return 1
	}

	x := big.NewInt(int64(code))
	x = x.Mul(x, r.decKey)
	x = x.Mod(x, r.base)
	return int(x.Int64())
}

// Decode converts a n-digit string (code) to another n-digit string (data)
//
// If str is not n-digit string, the err will be ErrFormat.
func (r *Converter) Decode(str string) (ret string, err error) {
	i, err := r.s2i(str)
	if err != nil {
		return
	}

	ret = r.i2s(r.Dec(i))
	return
}

// IsSafe tests if both function are safe
//
// It checks every element with following rules:
//
//   1. x == Dec(Enc(x))
//   2. at least 99% of all element fulfills x != Enc(x)
//   3. both functions are bijection
//
// This is *super* time/memory consuming task, do not use this in production.
func (r *Converter) IsSafe() (ok bool) {
	ok, _, _ = rngIs2Way(r)
	return ok && rngNotDupe(r)
}

func rngIs2Way(r *Converter) (ok bool, orig, code int) {
	base := int(r.base.Int64())
	cnt := 0
	for i := 0; i < base-1; i++ {
		code = r.Enc(i)
		actual := r.Dec(code)
		if actual != i {
			orig = i
			return
		}
		if actual == code {
			cnt++
		}
	}
	if cnt > 3 {
		return
	}

	ok = true
	return
}

func rngNotDupe(r *Converter) (ok bool) {
	m := map[int]bool{}
	base := int(r.base.Int64())
	for i := 0; i < base-1; i++ {
		m[r.Enc(i)] = true
	}

	ok = len(m) == base-1
	return
}

// RawConverter creates a Converter that converts n-digit string
//
// It does not check if number of digits is reasonable.
func RawConverter(digits, key int) (ret *Converter, err error) {
	base := int(math.Pow10(digits)) + 1
	ret = &Converter{
		base:   big.NewInt(int64(base)),
		encKey: big.NewInt(int64(key)),
		key:    key,
		digits: digits,
		over:   [2]int{0, 0},
	}
	x := big.NewInt(0).ModInverse(ret.encKey, ret.base)
	if x == nil {
		err = ErrKey(key)
		return
	}

	ret.decKey = x

	ret.over[1] = ret.Enc(base - 1)
	for x := 0; x < base-1; x++ {
		if ret.Enc(x) >= base-1 {
			ret.over[0] = x
			return
		}
	}

	err = ErrZero{}
	return
}

// MustRawConverter wraps RawConverter with panic
func MustRawConverter(digits, key int) (ret *Converter) {
	ret, err := RawConverter(digits, key)
	if err != nil {
		panic(err)
	}

	return
}

// NewConverter creates a Converter that converts n-digit string
//
// It check if number of digits is reasonable (3~7 digits):
//
//   * You should use plain mapping when digits < 3.
//   * It consumes too many resources to run IsSafe() when digits > 7.
func NewConverter(digits, key int) (ret *Converter, err error) {
	if digits < 3 || digits > 7 {
		err = ErrLength(digits)
		return
	}
	base := int(math.Pow10(digits)) + 1
	ret = &Converter{
		base:   big.NewInt(int64(base)),
		encKey: big.NewInt(int64(key)),
		key:    key,
		digits: digits,
		over:   [2]int{0, 0},
	}
	x := big.NewInt(0).ModInverse(ret.encKey, ret.base)
	if x == nil {
		err = ErrKey(key)
		return
	}

	ret.decKey = x

	ret.over[1] = ret.Enc(base - 1)
	for x := 0; x < base-1; x++ {
		if ret.Enc(x) >= base-1 {
			ret.over[0] = x
			return
		}
	}

	err = ErrZero{}
	return
}

// MustConverter wraps NewConverter with panic
func MustConverter(digits, key int) (ret *Converter) {
	ret, err := NewConverter(digits, key)
	if err != nil {
		panic(err)
	}

	return
}
