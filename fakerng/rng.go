/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package fakerng

// Converter maps a N-digit string to another n-digit string
//
// It uses several Converter to do the job.
type RNG struct {
	c      []*Converter
	digits int
}

// New Creates a new RNG instance
//
// The number of digit is determined by all Converters. If you pass a 6-digit
// Convert and a 5-digit Converter, the resulting RNG will be 11-digit.
func New(c ...*Converter) (ret *RNG) {
	ret = &RNG{
		c: c,
	}

	for _, x := range c {
		ret.digits += x.digits
	}

	return
}

func (r *RNG) source() (ret *source) {
	buf := &source{}
	buf.line("fakerng.New(")
	buf.block()
	for _, c := range r.c {
		buf.merge(c.source())
		buf.lines[len(buf.lines)-1] += ","
	}
	buf.endBlock()
	buf.line(")")

	return buf
}

// Source dumps golang source of the RNG instance
func (r *RNG) Source() (ret string) {
	return r.source().String()
}

// Digits returns supported string length
func (r *RNG) Digits() (ret int) {
	return r.digits
}

// Encode converts a n-digit string (data) to another n-digit string (code)
//
// If str is not n-digit string, the err will be ErrFormat.
func (r *RNG) Encode(str string) (ret string, err error) {
	if err = isDigit(str, r.digits); err != nil {
		return
	}

	begin := 0
	for _, c := range r.c {
		end := begin + c.digits
		x, err := c.Encode(str[begin:end])
		if err != nil {
			return "", err
		}
		ret += x
		begin = end
	}

	return
}

// Decode converts a n-digit string (code) to another n-digit string (data)
//
// If str is not n-digit string, the err will be ErrFormat.
func (r *RNG) Decode(str string) (ret string, err error) {
	if err = isDigit(str, r.digits); err != nil {
		return
	}

	begin := 0
	for _, c := range r.c {
		end := begin + c.digits
		x, err := c.Decode(str[begin:end])
		if err != nil {
			return "", err
		}
		ret += x
		begin = end
	}

	return
}
