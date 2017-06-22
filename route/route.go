// Copyright 2016 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin dragonfly freebsd netbsd openbsd

package route

import "golang.org/x/net/route"

func Fuzz(data []byte) int {
	var lastErr error
	for i := 0; i < 64; i++ {
		if _, err := route.ParseRIB(route.RIBType(i), data); err != nil {
			lastErr = err
			continue
		}
	}
	if lastErr != nil {
		return 0
	}
	return 1
}
