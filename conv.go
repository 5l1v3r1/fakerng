/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package main

import (
	"errors"
	"log"
	"math"
	"math/rand"

	"github.com/raohwork/fakerng/fakerng"
)

func genrawconv(i int) (key int, err error) {
	base := int(math.Pow10(i))
	for x := 0; x < 10; x++ {
		key = rand.Int() % (base / 10)
		if key == 0 {
			log.Print(key)
			x--
			continue
		}
		key = base - key

		rng, e := fakerng.RawConverter(i, key)
		if e != nil {
			log.Printf("failed try for %d: %s", key, e)
			x--
			continue
		}

		if unsafe {
			return
		}

		if !rng.IsSafe() {
			log.Printf("key %d is not safe", key)
			continue
		}

		return
	}

	err = errors.New("cannot generate good param after 10 tries")
	return
}

func split(i int) (parts []int) {
	const sz = 6
	cnt := i / sz
	parts = make([]int, cnt, cnt+1)
	for idx, _ := range parts {
		parts[idx] = sz
	}

	if x := i % sz; x > 0 {
		parts = append(parts, x)
	}

	for x := len(parts) - 1; x > 0; x-- {
		n := parts[x]
		if n >= 4 {
			break
		}

		parts[x-1] -= 4 - n
		parts[x] = 4
	}

	return
}

func genmulti(i int) (ret *fakerng.RNG) {
	sz := split(i)
	c := make([]*fakerng.Converter, 0, len(sz))
	for _, digit := range sz {
		key, err := genrawconv(digit)
		if err != nil {
			log.Fatal(err)
			return
		}

		c = append(c, fakerng.MustRawConverter(digit, key))
	}

	return fakerng.New(c...)
}
