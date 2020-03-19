/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package fakerng

import "testing"

func TestRNG(t *testing.T) {
	cases := [][2]string{
		{"000000", "00000"},
		{"000001", "00000"},
		{"000010", "00000"},
		{"000100", "00000"},
		{"000000", "00001"},
		{"000000", "00010"},
		{"000000", "00100"},
		{"000001", "00001"},
	}

	c1 := MustConverter(6, 926017)
	c2 := MustConverter(5, 93457)
	rng := New(c1, c2)

	for _, c := range cases {
		t.Run(c[0]+"_"+c[1], func(t *testing.T) {

			expect, err := rng.Encode(c[0] + c[1])
			if err != nil {
				t.Fatalf("rng.Encode: unexpected error: %s", err)
			}
			dec, err := rng.Decode(expect)
			if err != nil {
				t.Fatalf("rng.Decode: unexpected error: %s", err)
			}

			if dec != c[0]+c[1] {
				t.Log("enc:", expect)
				t.Log("origin:", c[0]+c[1])
				t.Log("actual:", dec)
				t.Fatal("unexpected result")
			}

			p1, err := c1.EncStr(c[0])
			if err != nil {
				t.Fatalf("c1.EncStr: unexpected error: %s", err)
			}
			p2, err := c2.EncStr(c[1])
			if err != nil {
				t.Fatalf("c2.EncStr: unexpected error: %s", err)
			}

			if expect != p1+p2 {
				t.Log("expect:", expect)
				t.Log("p1:", p1)
				t.Log("p2:", p2)
				t.Fatal("unexpected output")
			}
		})
	}
}
