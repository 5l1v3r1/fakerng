/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package fakerng

import (
	"fmt"
	"strings"
)

type source struct {
	lines []string
	curlv uint
}

func (s *source) block() (ret *source) {
	s.curlv++
	return s
}

func (s *source) endBlock() (ret *source) {
	if s.curlv > 0 {
		s.curlv--
	}

	return s
}

func (s *source) line(str string, data ...interface{}) (ret *source) {
	l := fmt.Sprintf(strings.Repeat("\t", int(s.curlv))+str, data...)
	s.lines = append(s.lines, l)
	return s
}

func (s *source) merge(block *source) (ret *source) {
	for _, l := range block.lines {
		s.line(l)
	}
	return s
}

func (s *source) String() (ret string) {
	return strings.Join(s.lines, "\n")
}
