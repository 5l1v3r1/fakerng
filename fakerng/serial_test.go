/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package fakerng

import (
	"log"
	"testing"
)

func TestRNG2Way(t *testing.T) {
	rng, err := NewConverter(6, 999169)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if ok, orig, code := rngIs2Way(rng); !ok {
		log.Fatal("NOT 2 WAY:", orig, ", ", code)
	}
}

func TestRNGNotDupe(t *testing.T) {
	rng, err := NewConverter(6, 999169)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if !rngNotDupe(rng) {
		log.Fatal("DUPED")
	}
}
