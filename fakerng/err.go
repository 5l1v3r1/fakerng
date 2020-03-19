/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package fakerng

import (
	"fmt"
	"regexp"
	"strconv"
)

var reDigit *regexp.Regexp

func init() {
	reDigit = regexp.MustCompile("^[0-9]+$")
}

func isDigit(str string, digit int) (err error) {
	if len(str) == digit && reDigit.MatchString(str) {
		return
	}

	return &ErrFormat{str: str, digit: digit}
}

type ErrFormat struct {
	digit int
	str   string
}

func (e *ErrFormat) Error() string {
	return fmt.Sprintf(
		"fakerng: malform input for %d digit rng: %s",
		e.digit, e.str,
	)
}

type ErrLength int

func (e ErrLength) Error() string {
	return "fakerng: unsupported digit length: " + strconv.Itoa(int(e))
}

type ErrKey int

func (e ErrKey) Error() string {
	return "fakerng: key has no modinverse: " + strconv.Itoa(int(e))
}

type ErrNPassLength struct{}

func (e ErrNPassLength) Error() string {
	return "fakerng: NPass need all RNGs with same digits"
}

type ErrZero struct{}

func (e ErrZero) Error() string {
	return "fakerng: faied to find zero value"
}
