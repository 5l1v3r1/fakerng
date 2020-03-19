/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package main

import (
	"log"

	"github.com/raohwork/fakerng/fakerng"
)

func genrng(i int) (ret *fakerng.RNG) {
	if i <= 7 {
		key, err := genrawconv(i)
		if err != nil {
			log.Fatal(err)
		}
		return fakerng.New(fakerng.MustRawConverter(i, key))
	}

	return genmulti(i)
}
