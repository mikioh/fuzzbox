// Copyright 2017 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ipv6

import "golang.org/x/net/ipv6"

func Fuzz(data []byte) int {
	var cm ipv6.ControlMessage
	if err := cm.Parse(data); err != nil {
		return 0
	}
	return 1
}
