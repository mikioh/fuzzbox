// Copyright 2015 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ipaddr

import "github.com/mikioh/ipaddr"

func Fuzz(data []byte) int {
	c, err := ipaddr.Parse(string(data))
	if err != nil {
		return 0
	}
	if c.List() == nil {
		panic("no prefix found")
	}
	return 1
}
