/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package fakerng

import (
	"fmt"
	"testing"
)

func TestNPassEncDec(t *testing.T) {
	np, _ := NewNPass(
		New(MustConverter(5, 98411)),
		New(MustConverter(5, 93089)),
	)

	for i := 0; i <= 99999; i++ {
		data := fmt.Sprintf("%05d", i)
		t.Log(i)
		code, err := np.Encode(data)
		if err != nil {
			t.Fatalf("unexpected error when encoding: %s", err)
		}
		actual, err := np.Decode(code)
		if err != nil {
			t.Fatalf("unexpected error when decoding: %s", err)
		}

		if data != actual {
			t.Log("expect:", data)
			t.Log("actual:", actual)
			t.Fatal("unexpected result")
		}
	}

}
