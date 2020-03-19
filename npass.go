/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package main

import "github.com/raohwork/fakerng/fakerng"

func gennpass(n uint, i int) (ret *fakerng.NPass) {
	rngs := make([]*fakerng.RNG, n)
	for x := 0; x < int(n); x++ {
		rngs[x] = genrng(i)
	}

	np, _ := fakerng.NewNPass(rngs...)
	return np
}
