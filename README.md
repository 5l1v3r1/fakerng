A simple tool to convert serial numbers into another number that looks like random number.

It uses modular multiplicative inverse to build a bijection function and its inverse function. The domain and codomain of both functions are n-digit integer (0 ~ 10^n-1, inclusive).

# Usage

``` golang
// integer
a := fakerng.MustRawConverter(4, 9529)
a.Enc(9991) // returns 4720
a.Dec(4720) // returns 9991
// string
str, err := a.Encode("9991") // returns "4720"
str, err := a.Decode("4720") // returns "9991"

// more digits (12 digits
b := fakerng.New(
        fakerng.MustRawConverter(6, 910110),
        fakerng.MustRawConverter(6, 977691),
)
str, err := b.Encode("000000002278") // returns "910110177871"
str, err := b.Decode("910110177871") // returns "000000002278"

// multiple pass to increase randomness (5 pass)
c := fakerng.NewNPass(
        fakerng.New(
                fakerng.MustRawConverter(4, 9566),
        ),
        fakerng.New(
                fakerng.MustRawConverter(4, 9456),
        ),
        fakerng.New(
                fakerng.MustRawConverter(4, 9247),
        ),
        fakerng.New(
                fakerng.MustRawConverter(4, 9953),
        ),
        fakerng.New(
                fakerng.MustRawConverter(4, 9683),
        ),
)
str, err := c.Encode("9991") // returns "5373"
str, err := c.Decode("5373") // returns "9991"
```

There's a tool to generate codes above for you.

```sh
go get -u github.com/raohwork/fakerng

fakerng -h
```

# License

This software is subject to the terms of the Mozilla Public License, v. 2.0. If a copy of the MPL was not distributed with this file, You can obtain one at [Mozilla MPL](https://mozilla.org/MPL/2.0/).

