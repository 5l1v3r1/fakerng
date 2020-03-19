/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package fakerng

type FakeRNG interface {
	Digits() int
	Encode(str string) (ret string, err error)
	Decode(str string) (ret string, err error)
	Source() string
}
