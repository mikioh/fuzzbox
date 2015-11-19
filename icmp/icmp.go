// Copyright 2016 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package icmp

import "golang.org/x/net/icmp"

func Fuzz(data []byte) int {
	icmp.ParseIPv4Header(data)
	var lastErr error
	for _, n := range []int{1, 58} {
		if _, err := icmp.ParseMessage(n, data); err != nil {
			lastErr = err
		}
	}
	if lastErr != nil {
		return 0
	}
	return 1
}
