/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package fakerng

// NPass uses multiple RNGs to make serial numbers "more random"
type NPass struct {
	rngs []*RNG
}

// NewNPass creates NPass instance
func NewNPass(rngs ...*RNG) (ret *NPass, err error) {
	digits := rngs[0].Digits()
	for _, r := range rngs {
		if digits != r.Digits() {
			err = ErrNPassLength{}
			return
		}
	}

	return &NPass{rngs: rngs}, nil
}

func (r *NPass) source() (ret *source) {
	buf := &source{}
	buf.line("fakerng.NewNPass(")
	buf.block()
	for _, c := range r.rngs {
		buf.merge(c.source())
		buf.lines[len(buf.lines)-1] += ","
	}
	buf.endBlock()
	buf.line(")")

	return buf
}

// Source dumps golang source of the NPass instance
func (r *NPass) Source() (ret string) {
	return r.source().String()
}

// Digits returns supported string length
func (r *NPass) Digits() (ret int) {
	return r.rngs[0].Digits()
}

// Encode converts a n-digit string (data) to another n-digit string (code)
//
// If str is not n-digit string, the err will be ErrFormat.
func (n *NPass) Encode(str string) (ret string, err error) {
	ret = str
	for _, r := range n.rngs {
		ret, err = r.Encode(ret)
		if err != nil {
			return "", err
		}
	}

	return
}

// Decode converts a n-digit string (code) to another n-digit string (data)
//
// If str is not n-digit string, the err will be ErrFormat.
func (n *NPass) Decode(str string) (ret string, err error) {
	ret = str
	for x := len(n.rngs) - 1; x >= 0; x-- {
		ret, err = n.rngs[x].Decode(ret)
		if err != nil {
			return "", err
		}
	}

	return
}
