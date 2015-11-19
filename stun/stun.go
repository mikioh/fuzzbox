// Copyright 2015 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stun

import "github.com/mikioh/stun"

func Fuzz(data []byte) int {
	_, m, err := stun.ParseMessage(data, nil)
	if err != nil {
		return 0
	}
	b := make([]byte, m.Len())
	if _, err := m.Marshal(b, nil); err != nil {
		panic(err)
	}
	return 1
}
