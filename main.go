/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/raohwork/fakerng/fakerng"
)

func dump(rng fakerng.FakeRNG) {
	fmt.Println(rng.Source())
	if test {
		try(rng)
	}
}

func try(rng fakerng.FakeRNG) {
	digit := rng.Digits()
	base := int(math.Pow10(digit))
	for x := 0; x < base; x++ {
		data := fmt.Sprintf("%0"+strconv.Itoa(digit)+"d", x)
		code, _ := rng.Encode(data)
		fmt.Printf("%s => %s\n", data, code)
	}
}

var (
	test   bool
	unsafe bool
	pass   uint
)

func main() {
	var i uint
	flag.BoolVar(&test, "test", false, "dump test data")
	flag.BoolVar(&unsafe, "unsafe", false, "do not call IsSafe() to validate safety")
	flag.UintVar(&pass, "pass", 1, "passes to use")
	flag.UintVar(&i, "digits", 6, "number of digits")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	if pass > 1 {
		dump(gennpass(pass, int(i)))
		return
	}

	dump(genrng(int(i)))
}
